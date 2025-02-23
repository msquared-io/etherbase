package streaming

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	"github.com/msquared-io/etherbase/backend/internal/schema"
	"github.com/msquared-io/etherbase/backend/internal/sources"
	"github.com/msquared-io/etherbase/backend/internal/state"
	"github.com/msquared-io/etherbase/backend/internal/subscription"
	"github.com/msquared-io/etherbase/backend/internal/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Add these types at the top of the file after imports
type ViewCallback struct {
	EventUpdates  []SubscriberEventUpdate
	StateUpdates  []SubscriberStateUpdate
	BlockNumber   uint64
}

type SubscriberEventUpdate struct {
	Event            DecodedEvent
	Subscriptions    map[types.ClientID][]types.SubscriptionID
	TransactionIndex uint
}

type SubscriberStateUpdate struct {
	SubscriptionID    types.SubscriptionID
	ContractAddress   common.Address
	State            map[string]interface{}
	TransactionIndex uint
}

type DecodedEvent struct {
	Sequence        int
	ContractAddress common.Address
	TransactionHash common.Hash
	Name            string
	ID              string
	NonIndexedArgs  map[string]interface{}
	IndexedArgs     map[string]interface{}
}

// Add after the other type definitions
type DecodedEventsResult struct {
	DecodedEvents []DecodedEvent
	Errors        map[string]string
}

type UpdatesResult struct {
	EventUpdates  []SubscriberEventUpdate
	StateUpdates  []SubscriberStateUpdate
	BlockNumber   uint64
}

// SubscriptionResult holds the event and state subscriptions for clients
type SubscriptionResult struct {
	EventSubscriptions map[types.ClientID][]types.SubscriptionID
	StateSubscriptions map[types.ClientID][]types.SubscriptionID
}

// DataType represents the type of data stored in the EthDB
// type DataType uint8

// const (
// 	DataTypeNone DataType = iota
// 	DataTypeString
// 	DataTypeBool
// 	DataTypeUint256
// 	DataTypeInt256
// 	DataTypeBytes
// )

// StreamingView handles real-time blockchain event streaming and state updates
type StreamingView struct {
	onFlush    func(updates *ViewCallback)
	shouldStop atomic.Bool
	updatesMu sync.RWMutex
	// currentEventFilterState struct {
	// 	Addresses []common.Address
	// 	Events    []schema.EventSchema
	// }
	// filterStateMu     sync.RWMutex // Only needed for currentEventFilterState
	transactionHashes sync.Map      // Replace map with sync.Map
	client           *ethclient.Client
	lastProcessedBlock uint64
	lastProcessedMu    sync.RWMutex
}

// NewStreamingView creates a new StreamingView instance
func NewStreamingView(client *ethclient.Client, onFlush func(updates *ViewCallback)) *StreamingView {
	sv := &StreamingView{
		onFlush:           onFlush,
		client:            client,
	}
	return sv
}

// Start begins watching for blockchain events
func (sv *StreamingView) Start(ctx context.Context) error {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	startBlock, err := sv.client.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("failed to get start block: %w", err)
	}
	log.Printf("Starting streaming view at block %d", startBlock)

	sv.shouldStop.Store(false)


	headers := make(chan *ethTypes.Header)
	sub, err := sv.client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return fmt.Errorf("failed to subscribe to new headers: %w", err)
	}

	go func() {
		defer sub.Unsubscribe()

		for {
			select {
			case err := <-sub.Err():
				log.Printf("Error in block subscription: %v", err)
				// return // dont return as we want to keep the subscription alive and retry
			case header := <-headers:
				if sv.shouldStop.Load() {
					return
				}

				log.Printf("Processing block %d", header.Number.Uint64())
				go func() {
					if err := sv.handleNewBlock(ctx, header); err != nil {
						log.Printf("Error handling block: %v", err)
					}
				}()
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

// Stop halts the event watching
func (sv *StreamingView) Stop() {
	sv.shouldStop.Store(true)
}

type Transaction struct {
	Hash common.Hash `json:"hash"`
}

type Transactions []Transaction

type Block struct {
	Transactions Transactions `json:"transactions"`
}

func (sv *StreamingView) handleNewBlock(ctx context.Context, header *ethTypes.Header) error {
	if header.TxHash == ethTypes.EmptyTxsHash {
		log.Printf("Skipping empty block %d", header.Number.Uint64())
		return nil
	}
	
	blockNumber := header.Number.Uint64()
	startTime := time.Now()
	log.Printf("Processing block %d, start time: %v", blockNumber, startTime)

	// Skip if we've already processed this block
	sv.lastProcessedMu.RLock()
	if blockNumber <= sv.lastProcessedBlock {
		sv.lastProcessedMu.RUnlock()
		return nil
	}
	sv.lastProcessedMu.RUnlock()

	// consider changing to pendng so we get it faster?

	var raw json.RawMessage
	err := sv.client.Client().CallContext(ctx, &raw, "eth_getBlockByHash", header.Hash(), true)
	if err != nil {
		return fmt.Errorf("failed to get block: %w", err)
	}
 
	var head *ethTypes.Header
	if err := json.Unmarshal(raw, &head); err != nil {
		return fmt.Errorf("failed to unmarshal header: %w", err)
	}
	if head == nil {
		return fmt.Errorf("block not found")
	}

	var body Block
	if err := json.Unmarshal(raw, &body); err != nil {
		return fmt.Errorf("failed to unmarshal block: %w", err)
	}

	// async wait for all receipts
	receipts := make([]ethTypes.Receipt, len(body.Transactions))
	var wg sync.WaitGroup
	wg.Add(len(body.Transactions))

	for i, tx := range body.Transactions {
		go func(i int, txHash common.Hash) {
			defer wg.Done()
			receipt, err := sv.client.TransactionReceipt(ctx, txHash)
			if err != nil {
				log.Printf("Failed to get receipt for tx %s: %v", txHash.String(), err)
				return
			}
			receipts[i] = *receipt
		}(i, tx.Hash)
	}

	wg.Wait()

	log.Printf("Got receipts in %v", time.Since(startTime))

	transactions := body.Transactions
	if len(transactions) == 0 {
		return fmt.Errorf("no transactions found in block %d but header says there are", blockNumber)
	}

	logs := make([]ethTypes.Log, 0)
	for _, receipt := range receipts {
		for _, log := range receipt.Logs {
			logs = append(logs, *log)
		}
	}

	if len(logs) == 0 {
		log.Printf("No logs found in block %d", blockNumber)
		return nil
	}

	log.Printf("Found %d relevant logs in block %d", len(logs), blockNumber)

	// Process all logs in the block
	decodedResult, err := sv.decodeLogsIntoEvents(logs)
	if err != nil {
		return fmt.Errorf("failed to decode logs: %w", err)
	}

	if len(decodedResult.DecodedEvents) == 0 {
		log.Printf("No events found in block %d", blockNumber)
		return nil
	}

	log.Printf("Decoded %d events from %d logs in block %d", len(decodedResult.DecodedEvents), len(logs), blockNumber)

	// we lock the function here as the below are not thread safe
	sv.updatesMu.Lock()
	defer sv.updatesMu.Unlock()

	// Get updates from all events
	updates, err := sv.getUpdatesFromEvents(ctx, decodedResult.DecodedEvents)
	if err != nil {
		return fmt.Errorf("failed to get updates: %w", err)
	}

	if len(updates.EventUpdates) == 0 && len(updates.StateUpdates) == 0 {
		log.Printf("No updates found in block %d", blockNumber)
		return nil
	}

	log.Printf("Got %d event updates and %d state updates from %d events in block %d", len(updates.EventUpdates), len(updates.StateUpdates), len(decodedResult.DecodedEvents), blockNumber)

	// Call onFlush with all updates for this block
	if sv.onFlush != nil {
		sv.onFlush(&ViewCallback{
			EventUpdates:  updates.EventUpdates,
			StateUpdates: updates.StateUpdates,
			BlockNumber:  blockNumber,
		})
	}

	// Update last processed block
	sv.lastProcessedMu.Lock()
	if blockNumber > sv.lastProcessedBlock {
		sv.lastProcessedBlock = blockNumber
	}
	sv.lastProcessedMu.Unlock()

	log.Printf("Processed block %d in %v", blockNumber, time.Since(startTime))
	return nil
}

func (sv *StreamingView) decodeLogsIntoEvents(logs []ethTypes.Log) (*DecodedEventsResult, error) {
	result := &DecodedEventsResult{
		DecodedEvents: make([]DecodedEvent, 0),
		Errors:        make(map[string]string),
	}

	for i, _log := range logs {
		// Skip logs with no topics
		if len(_log.Topics) == 0 {
			continue
		}

		eventSchema := sv.getEventSchemaFromLog(_log)
		if eventSchema == nil {
			result.Errors[sv.getEventIDFromLog(_log)] = "Event has no registered event schema"
			continue
		}

		decodedEvent, err := sv.decodeEvent(eventSchema, _log, i)
		if err != nil {
			log.Printf("Failed to decode event: %v", err)
			continue
		}

		result.DecodedEvents = append(result.DecodedEvents, *decodedEvent)
	}

	return result, nil
}

func (sv *StreamingView) getEventSchemaFromLog(log ethTypes.Log) *schema.EventSchema {
	eventID := sv.getEventIDFromLog(log)
	return schema.GetSchemaProvider().GetEventSchemaFromID(log.Address, eventID)
}

func (sv *StreamingView) getEventIDFromLog(log ethTypes.Log) string {
	return fmt.Sprintf("%s:%d", log.Topics[0].Hex(), len(log.Topics)-1)
}

func (sv *StreamingView) getUpdatesFromEvents(ctx context.Context, decodedEvents []DecodedEvent) (*UpdatesResult, error) {
	result := &UpdatesResult{
		EventUpdates:  make([]SubscriberEventUpdate, 0),
		StateUpdates: make([]SubscriberStateUpdate, 0),
	}

	fullContractUpdates := make(map[common.Address]bool)
	
	for _, event := range decodedEvents {
		// Track transaction indices
		if _, exists := sv.transactionHashes.Load(event.TransactionHash); !exists {
			sv.transactionHashes.Store(event.TransactionHash, uint(0))
		}

		eventSchema := schema.GetSchemaProvider().GetEventSchemaFromID(event.ContractAddress, event.ID)
		if eventSchema == nil {
			continue
		}

		subs := sv.getSubscriptionsToEvent(eventSchema, event)
		
		// Handle state updates
		if len(subs.StateSubscriptions) > 0 {
			if eventSchema.Name == "EthDBPathUpdate" {
				// Handle EthDBPathUpdate events
				stateUpdate, err := sv.handleEthDBPathUpdate(event, subs.StateSubscriptions)
				if err != nil {
					log.Printf("Error handling EthDBPathUpdate: %v", err)
					continue
				}
				result.StateUpdates = append(result.StateUpdates, stateUpdate...)
			} else {
				// Handle other state-changing events
				if !fullContractUpdates[event.ContractAddress] {
					stateUpdates, err := sv.handleContractStateUpdate(ctx, event, subs.StateSubscriptions)
					if err != nil {
							log.Printf("Error handling contract state update: %v", err)
						continue
					}
					result.StateUpdates = append(result.StateUpdates, stateUpdates...)
					fullContractUpdates[event.ContractAddress] = true
				}
			}
		}

		// Handle event updates
		if len(subs.EventSubscriptions) > 0 {
			result.EventUpdates = append(result.EventUpdates, SubscriberEventUpdate{
				Event:            event,
				Subscriptions:    subs.EventSubscriptions,
				TransactionIndex: func() uint {
					val, _ := sv.transactionHashes.Load(event.TransactionHash)
					return val.(uint)
				}(),
			})
		}
	}

	return result, nil
}

func buildTopics(events []schema.EventSchema) [][]common.Hash {
	var topics [][]common.Hash
	for _, event := range events {
		topics = append(topics, []common.Hash{event.EventTopic})
	}
	return topics
}

type AbiArgument struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Indexed bool `json:"indexed"`
}

func decodeAbiParameters(args []string, data []byte) ([]interface{}, error) {
	argsList := make([]AbiArgument, len(args))
	for i, arg := range args {
		argsList[i] = AbiArgument{
			Name:    "",
			Type:    arg,
			Indexed: false,
		}
	} 
	argListJson, err := json.Marshal(argsList)
	if err != nil {
		return nil, err
	}

	var realArgs abi.Arguments 
	err = json.Unmarshal(argListJson, &realArgs)
	if err != nil {
		return nil, err
	}

	decoded, err := realArgs.Unpack(data)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

func (sv *StreamingView) decodeEvent(eventSchema *schema.EventSchema, log ethTypes.Log, index int) (*DecodedEvent, error) {
	// customContractABI := schema.GetSchemaProvider().GetContractABI(log.Address)
	// if customContractABI == nil {
	// 	return nil, fmt.Errorf("no contract ABI found for address: %s", log.Address.Hex())
	// }

	topics := log.Topics[1:]
	
	values, err := eventSchema.Abi.Inputs.NonIndexed().Unpack(log.Data)
	if err != nil {
		return nil, err
	}

	nonIndexedArgs := make(map[string]interface{})
	indexedArgs := make(map[string]interface{})

	topicIndex := 0
	valueIndex := 0

	if len(values) > 0 {
		for _, arg := range eventSchema.Args {
			if arg.IsIndexed {
				// if the arg is a fixed-size type, abi decode it to the correct type
				if arg.ArgType == "uint256" || arg.ArgType == "int256" || arg.ArgType == "address" || arg.ArgType == "bool" || arg.ArgType == "bytes32" {
					decoded, err := decodeAbiParameters([]string{arg.ArgType}, topics[topicIndex].Bytes())
					if err != nil {
						return nil, err
					}
					// if address, parse as address
					var decodedTopic interface{}
					if arg.ArgType == "address" {
						decodedTopic = decoded[0].(common.Address).Hex()
					} else {
						decodedTopic = decoded[0]
					}

					indexedArgs[arg.Name] = decodedTopic
					topicIndex++
				} else {
					indexedArgs[arg.Name] = topics[topicIndex]
					topicIndex++
				}
			} else {
				nonIndexedArgs[arg.Name] = values[valueIndex]
				valueIndex++
			}
		}
	}

	return &DecodedEvent{
		Sequence:        index,
		ContractAddress: log.Address,
		TransactionHash: log.TxHash,
		Name:           eventSchema.Name,
		ID:             sv.getEventIDFromLog(log),
		NonIndexedArgs: nonIndexedArgs,
		IndexedArgs:    indexedArgs,
	}, nil
}

func (sv *StreamingView) getSubscriptionsToEvent(eventSchema *schema.EventSchema, decodedEvent DecodedEvent) SubscriptionResult {
	eventSubs := make(map[types.ClientID][]types.SubscriptionID)
	stateSubs := make(map[types.ClientID][]types.SubscriptionID)

	// Get all clients from subscription manager
	subscriptionManager := subscription.GetManager()
	clients := subscriptionManager.AllClients()

	for _, client := range clients {
		// Check event subscriptions
		for subID, eventSub := range client.GetEventSubscriptions() {
			if len(eventSub.ContractAddresses) > 0 {
				isAllowedContract := false
				for _, addr := range eventSub.ContractAddresses {
					if addr == decodedEvent.ContractAddress {
						isAllowedContract = true
						break
					}
				}
				if !isAllowedContract {
					continue
				}
			}

			// Check if event name matches any in the subscription
			for _, event := range eventSub.Events {
				if event.Name == eventSchema.Name || event.ID == eventSchema.EventTopic.String() {
					// if the event subscription has args, check if the decoded event matches the args
					skip := false
					for argName, argValues := range event.NonIndexedArgs {
						decodedArgValue, ok := decodedEvent.NonIndexedArgs[argName]
						if !ok {
							// if the arg is not in the decoded event, continue
							continue
						}

						// if decodedArgValue is bytes, convert to base64 string as that is what JSON unmarshals to
						if decodedArgValueBytes, ok := decodedArgValue.([]byte); ok {
							decodedArgValue = base64.StdEncoding.EncodeToString(decodedArgValueBytes)
						} 

						// compare the arg value as correct type, deep equal
						for _, argValue := range argValues {
							if !reflect.DeepEqual(decodedArgValue, argValue) {
								skip = true
								break
							}
						}
					}
					for argName, argValues := range event.IndexedArgs {
						decodedArgValue, ok := decodedEvent.IndexedArgs[argName]
						if !ok {
							// if the arg is not in the decoded event, continue
							continue
						}

						// compare the values directly, as they are both keccak256 hashes
						// if any of the values are equal, then the event matches
						exists := false
						for _, argValue := range argValues {
							if decodedArgValue == argValue {
								exists = true
								break
							}
						}
						if !exists {
							skip = true
							break
						}
					}
					if skip {
						continue
					}
					if _, exists := eventSubs[client.ID]; !exists {
						eventSubs[client.ID] = make([]types.SubscriptionID, 0)
					}
					eventSubs[client.ID] = append(eventSubs[client.ID], subID)
					break
				}
			}
		}

		// Check state subscriptions
		for subID, stateSub := range client.GetStateSubscriptions() {
			if stateSub.ContractAddress != decodedEvent.ContractAddress {
				continue
			}

			// we dont care about any event updates inside source contracts as we control them entirely
			isContractSource := sources.GetSourceRegistry().IsSource(stateSub.ContractAddress)
			if !isContractSource {
				// todo improve some of this logic
				if stateSub.Options.Repoll.OnAnyContractEvent {
					if _, exists := stateSubs[client.ID]; !exists {
						stateSubs[client.ID] = make([]types.SubscriptionID, 0)
					}
					stateSubs[client.ID] = append(stateSubs[client.ID], subID)
					continue
				}

				// Check listenEvents
				if len(stateSub.Options.Repoll.ListenEvents) > 0 {
					for _, listenEvent := range stateSub.Options.Repoll.ListenEvents {
						if listenEvent.Name == eventSchema.Name {
							// if the listen event subscription has args, check if the decoded event matches the args
							skip := false
							for argName, argValues := range listenEvent.NonIndexedArgs {
								decodedArgValue, ok := decodedEvent.NonIndexedArgs[argName]
								if !ok {
									// if the arg is not in the decoded event, continue
									continue
								}

								// if decodedArgValue is bytes, convert to base64 string as that is what JSON unmarshals to
								if decodedArgValueBytes, ok := decodedArgValue.([]byte); ok {
									decodedArgValue = base64.StdEncoding.EncodeToString(decodedArgValueBytes)
								} 

								// compare the arg value as correct type, deep equal
								for _, argValue := range argValues {
									if !reflect.DeepEqual(decodedArgValue, argValue) {
										skip = true
										break
									}
								}
							}
							for argName, argValues := range listenEvent.IndexedArgs {
								decodedArgValue, ok := decodedEvent.IndexedArgs[argName]
								if !ok {
									continue
								}

								exists := false
								for _, argValue := range argValues {
									if decodedArgValue == argValue {
										exists = true
										break
									}
								}
								if !exists {
									skip = true
									break
								}
							}
							if skip {
								continue
							}

							if _, exists := stateSubs[client.ID]; !exists {
								stateSubs[client.ID] = make([]types.SubscriptionID, 0)
							}
							stateSubs[client.ID] = append(stateSubs[client.ID], subID)
							break
						}
					}
				}
			}

			// Handle special events
			if eventSchema.Name == "EthDBUpdate" {
				if _, exists := stateSubs[client.ID]; !exists {
					stateSubs[client.ID] = make([]types.SubscriptionID, 0)
				}
				stateSubs[client.ID] = append(stateSubs[client.ID], subID)
			} else if eventSchema.Name == "EthDBPathUpdate" {
				// Check if event path matches or is parent of subscription path
				if eventPath, ok := decodedEvent.NonIndexedArgs["path"].([]string); ok {
					for _, statePath := range stateSub.StatePaths {
						// Check if statePath is a prefix of eventPath
						if isPathPrefix(statePath, eventPath) {
							if _, exists := stateSubs[client.ID]; !exists {
								stateSubs[client.ID] = make([]types.SubscriptionID, 0)
							}
							stateSubs[client.ID] = append(stateSubs[client.ID], subID)
							break
						}
					}
				}
			}
		}
	}

	return SubscriptionResult{
		EventSubscriptions: eventSubs,
		StateSubscriptions: stateSubs,
	}
}

// Helper function to check if one path is a prefix of another
func isPathPrefix(prefix []string, path []string) bool {
	if len(prefix) > len(path) {
		return false
	}
	for i, segment := range prefix {
		if segment != path[i] {
			return false
		}
	}
	return true
}

func (sv *StreamingView) handleEthDBPathUpdate(event DecodedEvent, stateSubscriptions map[types.ClientID][]types.SubscriptionID) ([]SubscriberStateUpdate, error) {
	updates := make([]SubscriberStateUpdate, 0)

	// Extract path and data from event args
	path, ok := event.NonIndexedArgs["path"].([]string)
	if !ok {
		return nil, fmt.Errorf("invalid path in EthDBPathUpdate event")
	}

	data, ok := event.NonIndexedArgs["data"].([]byte)
	if !ok {
		return nil, fmt.Errorf("invalid data in EthDBPathUpdate event")
	}

	dataType, ok := event.NonIndexedArgs["dataType"].(uint8)
	if !ok {
		return nil, fmt.Errorf("invalid dataType in EthDBPathUpdate event")
	}

	log.Printf("Processing EthDBPathUpdate event with path %v and data type %d and data %v", path, dataType, data)

	// todo verbose logging
	// stdlog.Printf("Processing EthDBPathUpdate event with path %v and data type %d and data %v", path, dataType, data)

	// Decode the value based on data type
	// var value interface{}
	// switch types.DataType(dataType) {
	// case types.DataTypeNone:
	// 	value = nil
	// case types.DataTypeString:
	// 	value = string(data)
	// case types.DataTypeBool:
	// 	value = len(data) > 0 && data[0] != 0
	// case types.DataTypeUint256:
	// 	value = new(big.Int).SetBytes(data)
	// case types.DataTypeInt256:
	// 	bigInt := new(big.Int).SetBytes(data)
	// 	if bigInt.Bit(255) == 1 { // Check if negative
	// 		bigInt.Sub(bigInt, new(big.Int).Lsh(big.NewInt(1), 256))
	// 	}
	// 	value = bigInt
	// case types.DataTypeBytes:
	// 	value = data
	// default:
	// 	return nil, fmt.Errorf("unsupported data type: %d", dataType)
	// }

	var value interface{}

	if len(data) == 0 {
		switch types.DataType(dataType) {
		case types.DataTypeNone:
			value = nil
		case types.DataTypeBool:
			value = false
		case types.DataTypeUint256:
			value = big.NewInt(0)
		case types.DataTypeInt256:
			value = big.NewInt(0)
		case types.DataTypeBytes:
			value = []byte{}
		default:
			return nil, fmt.Errorf("unsupported data type for zero length data: %d", dataType)
		}
	} else {
		var err error
		value, err = decodeAbiParameters([]string{types.DataTypeToString(types.DataType(dataType))}, data)
		if err != nil {
			return nil, fmt.Errorf("failed to decode data: %w", err)
		}
		value = value.([]interface{})[0]
	}

	// Create nested state object
	state := make(map[string]interface{})
	current := state
	for i, key := range path {
		if i == len(path)-1 {
			current[key] = value
		} else {
			current[key] = make(map[string]interface{})
			current = current[key].(map[string]interface{})
		}
	}

	// Create updates for all subscriptions
	for _, subscriptionIDs := range stateSubscriptions {
		for _, subID := range subscriptionIDs {
			updates = append(updates, SubscriberStateUpdate{
				SubscriptionID:    subID,
				ContractAddress:   event.ContractAddress,
				State:            state,
				TransactionIndex: func() uint {
					val, _ := sv.transactionHashes.Load(event.TransactionHash)
					return val.(uint)
				}(),
			})
		}
	}

	return updates, nil
}

func (sv *StreamingView) handleContractStateUpdate(ctx context.Context, event DecodedEvent, stateSubscriptions map[types.ClientID][]types.SubscriptionID) ([]SubscriberStateUpdate, error) {
	updates := make([]SubscriberStateUpdate, 0)

	// Get state manager instance
	stateManager := state.GetManager()

	// Fetch state for all subscriptions
	stateResult, _, err := stateManager.FetchStateForSubscriptions(ctx, stateSubscriptions)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch state: %w", err)
	}

	// Create updates for each subscription
	for subscriptionID, subState := range stateResult {
		updates = append(updates, SubscriberStateUpdate{
			SubscriptionID:    subscriptionID,
			ContractAddress:   event.ContractAddress,
			State:            subState,
			TransactionIndex: func() uint {
				val, _ := sv.transactionHashes.Load(event.TransactionHash)
				return val.(uint)
			}(),
		})
	}

	return updates, nil
}

// Add this helper function
func equalAddresses(a, b []common.Address) bool {
	if len(a) != len(b) {
		return false
	}
	seen := make(map[common.Address]int)
	for _, v := range a {
		seen[v]++
	}
	for _, v := range b {
		seen[v]--
		if seen[v] < 0 {
			return false
		}
	}
	return true
}

func equalEvents(a, b []schema.EventSchema) bool {
	if len(a) != len(b) {
		return false
	}
	seen := make(map[string]int)
	for _, v := range a {
		seen[v.EventTopic.String()]++
	}
	for _, v := range b {
		seen[v.EventTopic.String()]--
		if seen[v.EventTopic.String()] < 0 {
			return false
		}
	}
	return true
} 