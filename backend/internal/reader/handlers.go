package reader

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/websocket"
	"github.com/oklog/ulid/v2"

	"github.com/msquared-io/etherbase/backend/internal/schema"
	"github.com/msquared-io/etherbase/backend/internal/state"
	"github.com/msquared-io/etherbase/backend/internal/subscription"
	"github.com/msquared-io/etherbase/backend/internal/types"
)

var abiArgsCache sync.Map

func expandPaths(paths []interface{}) []types.StatePath {
	pathArrays := make([][]string, len(paths))
	for i, p := range paths {
		if arr, ok := p.([]interface{}); ok {
			pathArrays[i] = make([]string, len(arr))
			for j, a := range arr {
				pathArrays[i][j] = a.(string)
			}
		} else {
			pathArrays[i] = []string{p.(string)}
		}
	}
	return cartesianProduct(pathArrays)
}

func cartesianProduct(arrays [][]string) []types.StatePath {
	result := []types.StatePath{{}}
	for _, array := range arrays {
		newResult := []types.StatePath{}
		for _, r := range result {
			for _, a := range array {
				newResult = append(newResult, append(r, a))
			}
		}
		result = newResult
	}
	return result
}

func handleSubscribe(conn *websocket.Conn, client *subscription.Client, data json.RawMessage) {
	var subMsg SubscribeMessage
	if err := json.Unmarshal(data, &subMsg); err != nil {
		sendError(conn, "Invalid subscribe message format")
		return
	}

	if subMsg.PendingID == "" {
		sendError(conn, "Missing pendingId in subscribe message")
		return
	}

	subscriptionID := ulid.Make().String()

	if subMsg.EventSubscription != nil {
		if err := handleEventSubscription(client, subscriptionID, subMsg.EventSubscription); err != nil {
			sendMessage(conn, "subscription_failed", map[string]interface{}{
				"pendingId": subMsg.PendingID,
				"error":     err.Error(),
			})
			return
		}
	}
	if subMsg.StateSubscription != nil {
		if err := handleStateSubscription(conn, client, subscriptionID, subMsg.StateSubscription); err != nil {
			sendMessage(conn, "subscription_failed", map[string]interface{}{
				"pendingId": subMsg.PendingID,
				"error":     err.Error(),
			})
			return
		}
	}
	
	if subMsg.EventSubscription == nil && subMsg.StateSubscription == nil {
		sendError(conn, "No eventSubscription or stateSubscription present")
		return
	}

	sendMessage(conn, "subscription_success", map[string]interface{}{
		"pendingId":      subMsg.PendingID,
		"subscriptionId": subscriptionID,
	})
}

func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
    for _, v := range slice {
        if predicate(v) {
            return v, true
        }
    }
    var zero T
    return zero, false
}

type AbiArgument struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Indexed bool `json:"indexed"`
}

// todo coalesce with other copies
func encodeAbiParameters(args []string, values []interface{}) ([]byte, error) {
	argsList := make([]AbiArgument, len(args))
	for i, arg := range args {
		argsList[i] = AbiArgument{
			Name: "",
			Type: arg,
			Indexed: true,
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

	return realArgs.Pack(values...)
}

func handleEventSubscription(client *subscription.Client, subID string, sub *types.MessageEventSubscription) error {
	valid, err := client.ValidateEventSubscription(sub)
	if !valid {
		return err
	}

	var addresses []common.Address
	if sub.ContractAddresses != nil {
		addresses = make([]common.Address, len(sub.ContractAddresses))
		for i, addr := range sub.ContractAddresses {
			addresses[i] = common.HexToAddress(addr.Hex())
		}
	} else if sub.ContractAddress != nil {
		addresses = []common.Address{common.HexToAddress(sub.ContractAddress.Hex())}
	}

	// we don't support yet, todo allow adding a global schema via an ID?
	if len(addresses) == 0 {
		return fmt.Errorf("no contract addresses provided")
	}

	if len(sub.Events) == 0 {
		return fmt.Errorf("no events provided")
	}

	var events []types.Event
	events = make([]types.Event, len(sub.Events))
	for i, event := range sub.Events {
		targetEventSchema := schema.GetSchemaProvider().GetEventSchemaByName(addresses[0], event.Name)
		if targetEventSchema == nil {
			return fmt.Errorf("no schema found for event %s", event.Name)
		}

		// sort indexed and non-indexed args - if non-indexed arg, we save the hash of the arg value now as an optimisation as that is what we will compare
		indexedArgs := make(map[string][]interface{})
		nonIndexedArgs := make(map[string][]interface{})
		for argName, argValues := range event.Args {
			schemaArg, ok := Find(targetEventSchema.Args, func(a schema.Argument) bool {
				return a.Name == argName
			})
			if !ok {
				return fmt.Errorf("no schema arg found for event %s", event.Name)
			}

			for _, argValue := range argValues {
				if schemaArg.IsIndexed {
					// if the indexed arg is a string, bytes, struct, or array, we hash it. otherwise, abi encode it with padding to 32 bytes
				if schemaArg.ArgType == "string" || schemaArg.ArgType == "bytes" || schemaArg.ArgType == "struct" || schemaArg.ArgType == "array" {
						// enforce argValue to be a string. If it isn't, we return an error
						if argValueString, ok := argValue.(string); ok {
							indexedArgs[argName] = append(indexedArgs[argName], crypto.Keccak256Hash([]byte(argValueString)))
						} else {
							return fmt.Errorf("indexed arg %s with value %v is not a string", argName, argValue)
						}
				} else {
					value := argValue
					if schemaArg.ArgType == "address" {
						value = common.HexToAddress(value.(string))
					}
					
					indexedArgs[argName] = append(indexedArgs[argName], value)
				}
			} else {
					nonIndexedArgs[argName] = append(nonIndexedArgs[argName], argValue)
				}
			}
		}

		newEvent := types.Event{
			Name: event.Name,
			IndexedArgs: indexedArgs,
			NonIndexedArgs: nonIndexedArgs,
		}
		events[i] = newEvent
	}


	eventSub := &types.EventSubscription{
		ContractAddresses: addresses,
		Events:            events,
	}

	client.SetEventSubscription(types.SubscriptionID(subID), eventSub)
	return nil
}

var (
	trueVal = true
)

func handleStateSubscription(conn *websocket.Conn, client *subscription.Client, subID string, sub *types.MessageStateSubscription) error {
	valid, err := client.ValidateStateSubscriptionEvents(sub)
	if !valid {
		return err
	}
	
	address := common.HexToAddress(sub.ContractAddress.Hex())

	options := types.StateSubscriptionOptions{
	}

	if sub.Options != nil && sub.Options.CombineStateUpdatesFromSameContract != nil {
		options.CombineStateUpdatesFromSameContract = *sub.Options.CombineStateUpdatesFromSameContract
	} else {
		options.CombineStateUpdatesFromSameContract = false
	}
	
	if sub.Options != nil && sub.Options.CombineStateUpdatesFromSameTransaction != nil {
		options.CombineStateUpdatesFromSameTransaction = *sub.Options.CombineStateUpdatesFromSameTransaction
	} else {
		options.CombineStateUpdatesFromSameTransaction = true
	}

	if sub.Options != nil && sub.Options.FullStateUpdateOnChange != nil {
		options.FullStateUpdateOnChange = *sub.Options.FullStateUpdateOnChange
	} else {
		options.FullStateUpdateOnChange = false
	}

	repoll := types.RepollOptions{
	}

	if sub.Options != nil && sub.Options.Repoll != nil {
		if sub.Options.Repoll.OnAnyContractEvent != nil {
			repoll.OnAnyContractEvent = *sub.Options.Repoll.OnAnyContractEvent
		}

		if sub.Options.Repoll.ListenEvents != nil {
			listenEvents := make([]types.ListenEvent, len(*sub.Options.Repoll.ListenEvents))
			for i, event := range *sub.Options.Repoll.ListenEvents {
				targetEventSchema := schema.GetSchemaProvider().GetEventSchemaByName(address, event.Name)
				if targetEventSchema == nil {
					return fmt.Errorf("no schema found for event %s", event.Name)
				}

				// sort indexed and non-indexed args - if non-indexed arg, we save the hash of the arg value now as an optimisation as that is what we will compare
				indexedArgs := make(map[string][]interface{})
				nonIndexedArgs := make(map[string][]interface{})
				for argName, argValues := range event.Args {
					schemaArg, ok := Find(targetEventSchema.Args, func(a schema.Argument) bool {
						return a.Name == argName
					})
					if !ok {
						return fmt.Errorf("no schema arg found for event %s", event.Name)
					}
					for _, argValue := range argValues {
					if schemaArg.IsIndexed {
						// if the indexed arg is a string, bytes, struct, or array, we hash it. otherwise, abi encode it with padding to 32 bytes
						if schemaArg.ArgType == "string" || schemaArg.ArgType == "bytes" || schemaArg.ArgType == "struct" || schemaArg.ArgType == "array" {
								// enforce argValue to be a string. If it isn't, we return an error
								if argValueString, ok := argValue.(string); ok {
									indexedArgs[argName] = append(indexedArgs[argName], crypto.Keccak256Hash([]byte(argValueString)))
								} else {
									return fmt.Errorf("indexed arg %s with value %v is not a string", argName, argValue)
								}
						} else {
							// special case for address
							value := argValue
							if schemaArg.ArgType == "address" {
								value = common.HexToAddress(value.(string)).Hex()
							}
							indexedArgs[argName] = append(indexedArgs[argName], value)
						}
					} else {
							nonIndexedArgs[argName] = append(nonIndexedArgs[argName], argValue)
						}
					}
				}

				listenEvents[i] = types.ListenEvent{
					Name: event.Name,
					IndexedArgs: indexedArgs,
					NonIndexedArgs: nonIndexedArgs,
				}
			}
			repoll.ListenEvents = listenEvents
		}
	}

	// default to repoll on any contract event if there are no listen events
	if repoll.OnAnyContractEvent == false && len(repoll.ListenEvents) == 0 {
		repoll.OnAnyContractEvent = true
	}

	options.Repoll = repoll

	stateSub := &types.StateSubscription{
		ContractAddress: address,
		StatePaths:      expandPaths(sub.StatePath),
		Options:         options,
	}	

	client.SetStateSubscription(types.SubscriptionID(subID), stateSub)

	// Fetch initial state in background
	go func() {
		updates, blockNumber, err := state.GetManager().FetchStateForSubscriptions(
			context.Background(),
			map[types.ClientID][]types.SubscriptionID{
				client.ID: {types.SubscriptionID(subID)},
			},
		)
		if err != nil {
			log.Printf("Error fetching initial state: %v", err)
			return
		}

			if subState, ok := updates[types.SubscriptionID(subID)]; ok {
			sendMessage(conn, "initial_state", map[string]interface{}{
				"updates": []map[string]interface{}{
					{
						"type":           "state",
						"subscriptionId": subID,
						"data": map[string]interface{}{
							"contractAddress": sub.ContractAddress.Hex(),
							"state":           subState,
						},
					},
				},
				"block": blockNumber,
			})
		}
	}()

	return nil
}

func handleUnsubscribe(conn *websocket.Conn, client *subscription.Client, data json.RawMessage) {
	var msg types.UnsubscribeMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		sendError(conn, "Invalid unsubscribe message format")
		return
	}

	client.RemoveSubscription(types.SubscriptionID(msg.SubscriptionID))
	log.Printf("[websocket] unsubscribed from %s", msg.SubscriptionID)
}