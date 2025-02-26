package integration

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/msquared-io/etherbase/backend/internal/abiparser"
	"github.com/msquared-io/etherbase/backend/internal/config"
	"github.com/msquared-io/etherbase/backend/internal/integration/go/erc20"
	"github.com/msquared-io/etherbase/backend/internal/integration/go/etherbase"
	"github.com/msquared-io/etherbase/backend/internal/integration/go/etherbasesource"
	"github.com/msquared-io/etherbase/backend/internal/integration/go/multicall3"
	"github.com/msquared-io/etherbase/backend/internal/integration/go/privatevariablecustomstateevent"
	"github.com/msquared-io/etherbase/backend/internal/integration/go/publicethdbupdateevent"
	"github.com/msquared-io/etherbase/backend/internal/reader"
	"github.com/msquared-io/etherbase/backend/internal/types"
	"github.com/msquared-io/etherbase/backend/internal/writer"
)

type DeployedContract struct {
	ContractAddress common.Address
	ABI            string
	Contract       interface{}
}

type EventDefinition struct {
	Name string
	Args []Argument
}

type Argument struct {
	Name      string
	ArgType   string
	IsIndexed bool
}

// Message types
type MessageType string

const (
	TypeSubscribe    MessageType = "subscribe"
	TypeUnsubscribe MessageType = "unsubscribe"
	TypeSetValue    MessageType = "set_value"
	TypeEmitEvent   MessageType = "emit_event"
	TypeUpdates     MessageType = "updates"
	TypeError       MessageType = "error"
	TypeInitialState MessageType = "initial_state"
	TypeSubscriptionFailed MessageType = "subscription_failed"
	TypeSubscriptionSucceess MessageType = "subscription_success"
)

type ClientMessage struct {
	Type MessageType     `json:"type"`
	Data json.RawMessage `json:"data"`
}

type SubscribeMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type SubscriptionSuccess struct {
		PendingID      string `json:"pendingId"`
		SubscriptionID string `json:"subscriptionId"`
}

type StateUpdateMessage struct {
	Block   uint64  `json:"block"`
	Updates []StateUpdate `json:"updates"`
}

type EventUpdateMessage struct {
	Block   uint64  `json:"block"`
	Updates []EventUpdate `json:"updates"`
}

type EventUpdate struct {
	Type           string `json:"type"`
	SubscriptionID string `json:"subscriptionId"`
	Data           EtherbaseEvent `json:"data"`
}

type EtherbaseEvent struct {
	ContractAddress string `json:"contractAddress"`
	Block           uint64 `json:"block"`
	Name            string `json:"name"`
	EventID         string `json:"eventId"`
	Args            map[string]interface{} `json:"args"`
}

type StateUpdate struct {
	Type           string `json:"type"`
	SubscriptionID string `json:"subscriptionId"`
	Data struct {
		ContractAddress string                 `json:"contractAddress"`
		State          map[string]interface{} `json:"state"`
	} `json:"data"`
}

type UpdateMessage struct {
	Block   uint64  `json:"block"`
	Updates []Update `json:"updates"`
}

type Update struct {
	Type           string `json:"type"`
	SubscriptionID string `json:"subscriptionId"`
	Data           interface{} `json:"data"`
}

type StateData struct {
	ContractAddress string                 `json:"contractAddress"`
	State          map[string]interface{} `json:"state"`
}

type EventData struct {
	ContractAddress string                 `json:"contractAddress"`
	Block          uint64                 `json:"block"`
	Name           string                 `json:"name"`
	EventID        string                 `json:"eventId"`
	Args           map[string]interface{} `json:"args"`
}

// Helper functions for checking updates
func hasStateUpdate(t *testing.T, msg ClientMessage, index int, expectedState map[string]interface{}) {
	var updateMessage UpdateMessage
	if err := json.Unmarshal(msg.Data, &updateMessage); err != nil {
		t.Fatalf("failed to unmarshal update message: %v", err)
	}
	assert.Greater(t, updateMessage.Block, uint64(0))
	require.True(t, index < len(updateMessage.Updates), "update index out of range")
	assert.Equal(t, "state", updateMessage.Updates[index].Type)

	stateData, ok := updateMessage.Updates[index].Data.(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, expectedState, stateData["state"])
}

func compareMapValues(t *testing.T, expected, actual interface{}, path string) {
	switch expectedValue := expected.(type) {
	case map[string]interface{}:
		actualMap, isMap := actual.(map[string]interface{})
		require.True(t, isMap, "value at path %s should be a map", path)
		
		for key, expectedSubValue := range expectedValue {
			actualSubValue, exists := actualMap[key]
			require.True(t, exists, "missing key %s at path %s", key, path)
			compareMapValues(t, expectedSubValue, actualSubValue, path+"."+key)
		}
	case []interface{}:
		actualArray, isArray := actual.([]interface{})
		require.True(t, isArray, "value at path %s should be an array", path)
		require.Equal(t, len(expectedValue), len(actualArray), "array length mismatch at path %s", path)
		
		for i, expectedItem := range expectedValue {
			compareMapValues(t, expectedItem, actualArray[i], fmt.Sprintf("%s[%d]", path, i))
		}
	default:
		assert.Equal(t, expected, actual, "value mismatch at path %s", path)
	}
}

func hasEventUpdate(t *testing.T, msg ClientMessage, index int, expectedEvent map[string]interface{}) {
	var updateMessage UpdateMessage
	if err := json.Unmarshal(msg.Data, &updateMessage); err != nil {
		t.Fatalf("failed to unmarshal update message: %v", err)
	}
	assert.Greater(t, updateMessage.Block, uint64(0))
	require.True(t, index < len(updateMessage.Updates), "update index out of range")
	assert.Equal(t, "event", updateMessage.Updates[index].Type)

	eventData, ok := updateMessage.Updates[index].Data.(map[string]interface{})
	require.True(t, ok)
	for key, expectedValue := range expectedEvent {
		actualValue, exists := eventData[key]
		require.True(t, exists, "missing key %s in event data", key)
		compareMapValues(t, expectedValue, actualValue, key)
	}
}

func hasUpdates(t *testing.T, msg ClientMessage, expectedCount int) {
	var updateMessage UpdateMessage
	if err := json.Unmarshal(msg.Data, &updateMessage); err != nil {
		t.Fatalf("failed to unmarshal update message: %v", err)
	}
	assert.Greater(t, updateMessage.Block, uint64(0))
	assert.Len(t, updateMessage.Updates, expectedCount)
}

func hasBlockNumber(t *testing.T, msg ClientMessage, blockNumber uint64) {
	var updateMessage UpdateMessage
	if err := json.Unmarshal(msg.Data, &updateMessage); err != nil {
		t.Fatalf("failed to unmarshal update message: %v", err)
	}
	assert.Equal(t, blockNumber, updateMessage.Block)
}

func getEventABI(event EventDefinition) string {
	abi := fmt.Sprintf("event %s(", event.Name)
	for i, param := range event.Args {
		abi += param.ArgType
		if param.IsIndexed {
			abi += " indexed"
		}
		abi += " " + param.Name
		if i < len(event.Args)-1 {
			abi += ", "
		}
	}
	abi += ")"
	return abi
}

func getEventTopic(event EventDefinition) common.Hash {
	signature := fmt.Sprintf("%s(%s)", event.Name,
		func() string {
			types := make([]string, len(event.Args))
			for i, arg := range event.Args {
				types[i] = arg.ArgType
			}
			return strings.Join(types, ",")
		}())
	return crypto.Keccak256Hash([]byte(signature))
}

func parseEvent(event EventDefinition) (string, common.Hash, []Argument) {
	eventABI := getEventABI(event)
	parsedABI, err := abi.JSON(strings.NewReader(fmt.Sprintf("[{%s}]", eventABI)))
	if err != nil {
		panic(err)
	}

	eventTopic := getEventTopic(event)
	return parsedABI.Events[event.Name].Name, eventTopic, event.Args
}

func deployTestContracts() (map[string]*DeployedContract, *ethclient.Client, error) {
	// Connect to local Ethereum node
	client, err := ethclient.Dial("ws://localhost:8545")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	// Use a pre-funded account from the local node
	key, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80") // Hardhat's first default private key
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(31337)) // Hardhat's chain ID
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	deployedContracts := make(map[string]*DeployedContract)

	// Deploy PrivateVariableCustomStateEvent
	privAddress, _, priv, err := privatevariablecustomstateevent.DeployPrivateVariableCustomStateEvent(auth, client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to deploy PrivateVariableCustomStateEvent: %v", err)
	}
	deployedContracts["PrivateVariableCustomStateEvent"] = &DeployedContract{
		ContractAddress: privAddress,
		Contract:       priv,
	}

	// Deploy ERC20
	erc20Address, _, erc20, err := erc20.DeployERC20(auth, client, big.NewInt(0))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to deploy ERC20: %v", err)
	}
	deployedContracts["ERC20"] = &DeployedContract{
		ContractAddress: erc20Address,
		Contract:       erc20,
	}

	validatorAddress, _, _, err := etherbase.DeployDataValidator(auth, client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to deploy DataValidator: %v", err)
	}
	
	etherbaseAddress, _, etherbase, err := etherbase.DeployEtherbase(auth, client, validatorAddress)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to deploy Etherbase: %v", err)
	}
	fmt.Printf("etherbaseAddress: %s\n", etherbaseAddress.Hex())
	deployedContracts["Etherbase"] = &DeployedContract{
		ContractAddress: etherbaseAddress,
		Contract:       etherbase,
	}

	multicallAddress, _, multicall, err := multicall3.DeployMulticall3(auth, client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to deploy Multicall: %v", err)
	}
	deployedContracts["Multicall3"] = &DeployedContract{
		ContractAddress: multicallAddress,
		Contract:       multicall,
	}
	
	sourceTx, err := etherbase.CreateSource(auth)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create source: %v", err)
	}
	sourceReceipt, err := bind.WaitMined(context.Background(), client, sourceTx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed waiting for transaction: %v", err)
	}
	if sourceReceipt.Status != uint64(1) {
		return nil, nil, fmt.Errorf("transaction failed")
	}
	sourceAddress := sourceReceipt.Logs[0].Address
	source, err := etherbasesource.NewEtherbaseSource(sourceAddress, client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create EtherbaseSource: %v", err)
	}
	deployedContracts["EtherbaseSource"] = &DeployedContract{
		ContractAddress: sourceAddress,
		Contract:       source,
	}

	// Deploy PublicVariableEthDBUpdateMappingEvent
	mappingEventAddr, _, mappingEvent, err := publicethdbupdateevent.DeployPublicVariableEthDBUpdateMappingEvent(auth, client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to deploy PublicVariableEthDBUpdateMappingEvent: %v", err)
	}
	deployedContracts["PublicVariableEthDBUpdateMappingEvent"] = &DeployedContract{
		ContractAddress: mappingEventAddr,
		Contract:       mappingEvent,
	}

	// Deploy PublicVariableEthDBUpdateMultipleEvent
	multipleEventAddr, _, multipleEvent, err := publicethdbupdateevent.DeployPublicVariableEthDBUpdateMultipleEvent(auth, client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to deploy PublicVariableEthDBUpdateMultipleEvent: %v", err)
	}
	deployedContracts["PublicVariableEthDBUpdateMultipleEvent"] = &DeployedContract{
		ContractAddress: multipleEventAddr,
		Contract:       multipleEvent,
	}

	return deployedContracts, client, nil
}

func getAbiFromEventDefinition(eventDef EventDefinition) (*abi.Event, error) {
	eventAbi := fmt.Sprintf("event %s(", eventDef.Name)
	for _, arg := range eventDef.Args {
		eventAbi += fmt.Sprintf("%s ", arg.ArgType)
		if arg.IsIndexed {
			eventAbi += "indexed "
		}
		eventAbi += fmt.Sprintf("%s, ", arg.Name)
	}
	eventAbi = eventAbi[:len(eventAbi)-2] + ")"
	abiElement, err := abiparser.ParseEventSignature(eventAbi)
	if err != nil {
		return nil, err
	}
	abiString, err := json.Marshal(abiElement)
	if err != nil {
		return nil, err
	}

	var outAbiEvent abi.Event
	err = json.Unmarshal(abiString, &outAbiEvent)
	if err != nil {
		return nil, err
	}

	return &outAbiEvent, nil
}

func registerEvent(t *testing.T, ethClient *ethclient.Client, source *etherbasesource.EtherbaseSource, eventDef EventDefinition) *abi.Event {
	eventTopic := getEventTopic(eventDef)
	eventArgs := make([]etherbasesource.Argument, len(eventDef.Args))
	for i, arg := range eventDef.Args {
		eventArgs[i] = etherbasesource.Argument{
			Name: arg.Name,
			ArgType: arg.ArgType,
			IsIndexed: arg.IsIndexed,
		}
	}
	numIndexedArgs := 0
	for _, arg := range eventDef.Args {
		if arg.IsIndexed {
			numIndexedArgs++
		}
	}
	eventId := fmt.Sprintf("%s:%d", eventTopic.Hex(), numIndexedArgs)

	key, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		t.Fatalf("failed to create transactor: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(31337))
	if err != nil {
		t.Fatalf("failed to create transactor: %v", err)
	}
	tx, err := source.RegisterEventSchema(auth, eventDef.Name, eventId, eventTopic, eventArgs)
	if err != nil {
		t.Fatalf("failed to register event schema: %v", err)
	}
	receipt, err := bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		t.Fatalf("failed waiting for transaction: %v", err)
	}
	if receipt.Status != uint64(1) {
		t.Fatalf("transaction failed")
	}

	abi, err := getAbiFromEventDefinition(eventDef)
	if err != nil {
		t.Fatalf("failed to get ABI: %v", err)
	}
	
	return abi
}

// Helper struct to manage test infrastructure
type testEnv struct {
	t            *testing.T
	contracts    map[string]*DeployedContract
	ethClient   *ethclient.Client
	ws          *websocket.Conn
	wsWriter    *websocket.Conn
	server      *httptest.Server
	writerServer *httptest.Server
	auth        *bind.TransactOpts
	done        chan bool
	errChan     chan error
}

// Create a new test environment
func setupTestEnv(t *testing.T) (*testEnv, error) {
	env := &testEnv{
		t:       t,
		done:    make(chan bool),
		errChan: make(chan error),
	}

	// Deploy test contracts
	contracts, ethClient, err := deployTestContracts()
	if err != nil {
		return nil, fmt.Errorf("failed to deploy contracts: %v", err)
	}
	env.contracts = contracts
	env.ethClient = ethClient

	// Setup auth
	privateKey := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(31337))
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}
	env.auth = auth

	// Setup server
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}
	cfg.EtherbaseAddress = contracts["Etherbase"].ContractAddress
	cfg.MulticallAddress = contracts["Multicall3"].ContractAddress
	cfg.RpcURL = "ws://localhost:8545"
	cfg.ChainID = 31337
	
	// Setup reader server
	readerSrv, err := reader.NewServer(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create reader server: %v", err)
	}

	env.server = httptest.NewServer(readerSrv.Router())
	go readerSrv.Start()

	// Setup writer server
	writerSrv, err := writer.NewServer(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create writer server: %v", err)
	}

	env.writerServer = httptest.NewServer(writerSrv.Router())
	go writerSrv.Start()

	time.Sleep(5 * time.Second)

	// Setup WebSocket connections
	wsURL := "ws" + strings.TrimPrefix(env.server.URL, "http") + "/read"
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to reader websocket: %v", err)
	}
	env.ws = ws

	// use a separate private key to avoid nonce conflicts
	privateKey2 := "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	wsWriterURL := "ws" + strings.TrimPrefix(env.writerServer.URL, "http") + "/write?privateKey=0x" + privateKey2
	wsWriter, _, err := websocket.DefaultDialer.Dial(wsWriterURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to writer websocket: %v", err)
	}
	env.wsWriter = wsWriter

	return env, nil
}

// Cleanup test environment
func (env *testEnv) cleanup() {
	if env.ws != nil {
		env.ws.Close()
	}
	if env.wsWriter != nil {
		env.wsWriter.Close()
	}
	if env.server != nil {
		env.server.Close()
	}
	if env.writerServer != nil {
		env.writerServer.Close()
	}
}

// Helper to handle websocket messages
func (env *testEnv) handleMessages(expectedMsgs []func(ClientMessage) error) {
	messageCount := 0
	go func() {
		for {
			var msg ClientMessage
			if err := env.ws.ReadJSON(&msg); err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					env.errChan <- fmt.Errorf("websocket error: %v", err)
				}
				return
			}

			if msg.Type == TypeError || msg.Type == TypeSubscriptionFailed {
				env.errChan <- fmt.Errorf("error: %v", string(msg.Data))
				return
			}

			fmt.Printf("msg: %v\n", msg)

			if messageCount >= len(expectedMsgs) {
				env.errChan <- fmt.Errorf("received unexpected message: %v", msg)
				return
			}

			if err := expectedMsgs[messageCount](msg); err != nil {
				env.errChan <- err
				return
			}

			messageCount++
			if messageCount == len(expectedMsgs) {
				env.done <- true
				return
			}
		}
	}()
}


type AbiArgument struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Indexed bool `json:"indexed"`
}

func encodeAbiParameters(args []string, values []interface{}) ([]byte, error) {
	argsList := make([]AbiArgument, len(args))
	for i, arg := range args {
		argsList[i] = AbiArgument{
			Name: "",
			Type: arg,
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

	return realArgs.Pack(values...)
}

func isSubscriptionSuccess(t *testing.T, msg ClientMessage, pendingID string) {
	assert.Equal(t, TypeSubscriptionSucceess, msg.Type)
	var subSuccess SubscriptionSuccess
	err := json.Unmarshal(msg.Data, &subSuccess)
	require.NoError(t, err)
	assert.Equal(t, pendingID, subSuccess.PendingID)
	assert.NotEmpty(t, subSuccess.SubscriptionID)
}

func TestEventSubscription(t *testing.T) {
	env, err := setupTestEnv(t)
	require.NoError(t, err)
	defer env.cleanup()

	source := env.contracts["EtherbaseSource"].Contract.(*etherbasesource.EtherbaseSource)
	sourceAddress := env.contracts["EtherbaseSource"].ContractAddress
	
	testEventName := "TestEvent"
	pendingID := "test-pending-id-1"

	// Register test event
	eventDef := EventDefinition{
		Name: testEventName,
		Args: []Argument{
			{Name: "value", ArgType: "bytes", IsIndexed: false},
		},
	}
	abi := registerEvent(t, env.ethClient, source, eventDef)
	time.Sleep(5 * time.Second)

	// Setup message handlers
	env.handleMessages([]func(ClientMessage) error{
		// Handle subscription confirmation
		func(msg ClientMessage) error {
			isSubscriptionSuccess(t, msg, pendingID)
			
			time.Sleep(2 * time.Second)

			// Emit test event
			data, err := abi.Inputs.Pack([]byte{0x12, 0x34})
			require.NoError(t, err)
			tx, err := source.EmitEvent(env.auth, testEventName, [][32]byte{}, data)
			require.NoError(t, err)
			receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status)
			return nil
		},
		// Handle event update
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasEventUpdate(t, msg, 0, map[string]interface{}{
				"contractAddress": sourceAddress.Hex(),
				"name": testEventName,
				"args": map[string]interface{}{
					"value": base64.StdEncoding.EncodeToString([]byte{0x12, 0x34}),
				},
			})
			return nil
		},
	})

	// Send subscription
	subMsg := ClientMessage{
		Type: TypeSubscribe,
		Data: json.RawMessage(fmt.Sprintf(`{
			"pendingId": "%s",
			"eventSubscription": {
				"contractAddress": "%s",
				"events": [{
					"name": "%s"
				}]
			}
		}`, pendingID, sourceAddress.Hex(), testEventName)),
	}
	require.NoError(t, env.ws.WriteJSON(subMsg))

	// Wait for completion or timeout
	select {
	case <-env.done:
	case err := <-env.errChan:
		t.Fatal(err)
	case <-time.After(30 * time.Second):
		t.Fatal("test timeout")
	}
}

func TestEventSubscriptionWithSpecificNonIndexedArg(t *testing.T) {
	env, err := setupTestEnv(t)
	require.NoError(t, err)
	defer env.cleanup()

	source := env.contracts["EtherbaseSource"].Contract.(*etherbasesource.EtherbaseSource)
	sourceAddress := env.contracts["EtherbaseSource"].ContractAddress
	
	testEventName := "TestEvent"
	pendingID := "test-pending-id-1"

	// Register test event
	eventDef := EventDefinition{
		Name: testEventName,
		Args: []Argument{
			{Name: "value", ArgType: "bytes", IsIndexed: false},
		},
	}
	abi := registerEvent(t, env.ethClient, source, eventDef)
	time.Sleep(5 * time.Second)

	// Setup message handlers
	env.handleMessages([]func(ClientMessage) error{
		// Handle subscription confirmation
		func(msg ClientMessage) error {
			isSubscriptionSuccess(t, msg, pendingID)
			
			time.Sleep(2 * time.Second)

			// emit a value we are not listening for
			data, err := abi.Inputs.Pack([]byte{0x12, 0x35})
			require.NoError(t, err)
			tx, err := source.EmitEvent(env.auth, testEventName, [][32]byte{}, data)
			require.NoError(t, err)
			receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status)

			time.Sleep(2 * time.Second)	

			// Emit test event
			data, err = abi.Inputs.Pack([]byte{0x12, 0x34})
			require.NoError(t, err)
			tx, err = source.EmitEvent(env.auth, testEventName, [][32]byte{}, data)
			require.NoError(t, err)
			receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status)
			return nil
		},
		// Handle event update
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasEventUpdate(t, msg, 0, map[string]interface{}{
				"contractAddress": sourceAddress.Hex(),
				"name": testEventName,
				"args": map[string]interface{}{
					"value": base64.StdEncoding.EncodeToString([]byte{0x12, 0x34}),
				},
			})

			return nil
		},
	})

	// Send subscription
	subMsg := ClientMessage{
		Type: TypeSubscribe,
		Data: json.RawMessage(fmt.Sprintf(`{
			"pendingId": "%s",
			"eventSubscription": {
				"events": [{
					"name": "%s",
					"args": {
						"value": ["%s"]
					}
				}],
				"contractAddress": "%s"
			}
		}`, pendingID, testEventName, base64.StdEncoding.EncodeToString([]byte{0x12, 0x34}), sourceAddress.Hex())),
	}
	require.NoError(t, env.ws.WriteJSON(subMsg))

	// Wait for completion or timeout
	select {
	case <-env.done:
	case err := <-env.errChan:
		t.Fatal(err)
	case <-time.After(50 * time.Second):
		t.Fatal("test timeout")
	}
}

func TestStateSubscriptionWithEtherbaseStateChange(t *testing.T) {
	env, err := setupTestEnv(t)
	require.NoError(t, err)
	defer env.cleanup()

	source := env.contracts["EtherbaseSource"].Contract.(*etherbasesource.EtherbaseSource)
	sourceAddress := env.contracts["EtherbaseSource"].ContractAddress
	pendingID := "test-pending-id-2"
	segments := []string{"testValue"}

	// Set initial state
	data, err := encodeAbiParameters([]string{"uint256"}, []interface{}{big.NewInt(100)})
	require.NoError(t, err)
	tx, err := source.SetValue(env.auth, segments, uint8(types.DataTypeUint256), data)
	require.NoError(t, err)
	receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)

	// Dummy transaction to bump block number
	tx, err = source.CreateWalletIdentity(env.auth, common.Address{}, []uint8{})
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)

	time.Sleep(5 * time.Second)

	// Setup message handlers
	env.handleMessages([]func(ClientMessage) error{
		// Handle subscription confirmation
		func(msg ClientMessage) error {
			isSubscriptionSuccess(t, msg, pendingID)
			return nil
		},
		// Handle initial state
		func(msg ClientMessage) error {
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"testValue": float64(100),
			})

			time.Sleep(2 * time.Second)

			// Update the value to 101
			data, err := encodeAbiParameters([]string{"uint256"}, []interface{}{big.NewInt(101)})
	require.NoError(t, err)
			tx, err := source.SetValue(env.auth, segments, uint8(types.DataTypeUint256), data)
			require.NoError(t, err)
			receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status)
			fmt.Printf("updated value to 101\n")
			return nil
		},
		// Handle state update
		func(msg ClientMessage) error {
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"testValue": float64(101),
			})
			return nil
		},
	})

	// Send subscription message
	subMsg := ClientMessage{
		Type: TypeSubscribe,
		Data: json.RawMessage(fmt.Sprintf(`{
			"pendingId": "%s",
			"stateSubscription": {
				"contractAddress": "%s",
				"statePath": %s
			}
		}`, pendingID, sourceAddress.Hex(), `["testValue"]`)),
	}
	require.NoError(t, env.ws.WriteJSON(subMsg))

	// Wait for completion or timeout
	select {
	case <-env.done:
	case err := <-env.errChan:
		t.Fatal(err)
	case <-time.After(30 * time.Second):
		t.Fatal("test timeout")
	}
}

func TestStateSubscriptionCustomContractAndEventAndOnlyThisContractAndOnlyMintingIndexedArg(t *testing.T) {
	env, err := setupTestEnv(t)
	require.NoError(t, err)
	defer env.cleanup()

	// Register custom contract with etherbase
	etherbase := env.contracts["Etherbase"].Contract.(*etherbase.Etherbase)
	tx, err := etherbase.AddCustomContract(
		env.auth,
		env.contracts["ERC20"].ContractAddress,
		erc20.ERC20ABI,
	)
	receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")


	// make a second ERC20 contract
	contract2Address, _, contract2, err := erc20.DeployERC20(env.auth, env.ethClient, big.NewInt(0))
	require.NoError(t, err)

	tx, err = etherbase.AddCustomContract(
		env.auth,
		contract2Address,
		erc20.ERC20ABI,
	)
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	time.Sleep(5 * time.Second)

	pendingID := "test-pending-id-3"
	contract := env.contracts["ERC20"].Contract.(*erc20.ERC20)

	address1 := env.auth.From
	
	pk2, err := crypto.GenerateKey()
	require.NoError(t, err)
	address2 := crypto.PubkeyToAddress(pk2.PublicKey)

	// mint 100 tokens to the auth
	tx, err = contract.Mint(env.auth, address1, big.NewInt(100))
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	transferBlock := uint64(0)
	mint1Block := uint64(0)
	mint2Block := uint64(0)

	// Setup message handlers
	env.handleMessages([]func(ClientMessage) error{
		// Handle subscription confirmation
		func(msg ClientMessage) error {
			isSubscriptionSuccess(t, msg, pendingID)
			return nil
		},
		// Handle initial state
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
 			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"balances": map[string]interface{}{address1.Hex(): float64(100)},
			})

			time.Sleep(2 * time.Second)

			// transfer 50 tokens from address2, which should not trigger an update as we only listen for minting
			tx, err = contract.Transfer(env.auth, address2, big.NewInt(50))
			require.NoError(t, err)
			receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status, "transaction failed")
			transferBlock = receipt.BlockNumber.Uint64()

			// mint 10000 on contract2, which should not trigger an update
			tx, err = contract2.Mint(env.auth, address1, big.NewInt(10000))
			require.NoError(t, err)
			receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status, "transaction failed")
			mint2Block = receipt.BlockNumber.Uint64()
			time.Sleep(2 * time.Second)

			// Call mint after receiving initial state
			tx, err := contract.Mint(env.auth, address1, big.NewInt(1))
			require.NoError(t, err)
	
			receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status, "transaction failed")
			mint1Block = receipt.BlockNumber.Uint64()
			return nil
		},
		// Handle state update after mint
		func(msg ClientMessage) error {
			require.True(t, mint1Block > transferBlock)
			require.True(t, mint1Block > mint2Block)
			hasBlockNumber(t, msg, mint1Block)

			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"balances": map[string]interface{}{address1.Hex(): float64(51)},
			})
			// Verify final contract state
			balance, err := contract.Balances(nil, address1)
			require.NoError(t, err)
			assert.Equal(t, int64(51), balance.Int64())

			return nil
		},
	})

	// Send subscription message
	subMsg := ClientMessage{
		Type: TypeSubscribe,
		Data: json.RawMessage(fmt.Sprintf(`{
			"pendingId": "%s",
			"stateSubscription": {
				"contractAddress": "%s",
				"statePath": ["balances", "%s"],
				"options": {
					"repoll": {
						"listenEvents": [{
							"name": "Transfer",
							"args": {
								"from": ["0x0000000000000000000000000000000000000000"]
							}
						}]
					}
				}
			}
		}`, pendingID, env.contracts["ERC20"].ContractAddress.Hex(), address1.Hex())),
	}
	require.NoError(t, env.ws.WriteJSON(subMsg))

	// Wait for completion or timeout
	select {
	case <-env.done:
	case err := <-env.errChan:
		t.Fatal(err)
	case <-time.After(30 * time.Second):
		t.Fatal("test timeout")
	}
}

// todo reader/handlers.go L150
// func TestEventSubscriptionCustomContractAndAnyContractIndexedArg(t *testing.T) {
// 	env, err := setupTestEnv(t)
// 	require.NoError(t, err)
// 	defer env.cleanup()

// 	// Register custom contract with etherbase
// 	etherbase := env.contracts["Etherbase"].Contract.(*etherbase.Etherbase)
// 	tx, err := etherbase.AddCustomContract(
// 		env.auth,
// 		env.contracts["ERC20"].ContractAddress,
// 		"ERC20",
// 		erc20.ERC20ABI,
// 	)
// 	receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
// 	require.NoError(t, err)
// 	require.Equal(t, uint64(1), receipt.Status, "transaction failed")


// 	// make a second ERC20 contract
// 	contract2Address, _, contract2, err := erc20.DeployERC20(env.auth, env.ethClient, big.NewInt(0))
// 	require.NoError(t, err)

// 	tx, err = etherbase.AddCustomContract(
// 		env.auth,
// 		contract2Address,
// 		"ERC20",
// 		erc20.ERC20ABI,
// 	)
// 	require.NoError(t, err)
// 	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
// 	require.NoError(t, err)
// 	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

// 	time.Sleep(5 * time.Second)

// 	pendingID := "test-pending-id-3"
// 	contract := env.contracts["ERC20"].Contract.(*erc20.ERC20)

// 	address1 := env.auth.From

// 	// mint 100 tokens to the auth
// 	tx, err = contract.Mint(env.auth, address1, big.NewInt(100))
// 	require.NoError(t, err)
// 	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
// 	require.NoError(t, err)
// 	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

// 	// Setup message handlers
// 	env.handleMessages([]func(ClientMessage) error{
// 		// Handle subscription confirmation
// 		func(msg ClientMessage) error {
// 			isSubscriptionSuccess(t, msg, pendingID)

// 			// mint 10000 on contract2, which should trigger an update
// 			tx, err = contract2.Mint(env.auth, address1, big.NewInt(10000))
// 			require.NoError(t, err)
// 			receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
// 			require.NoError(t, err)
// 			require.Equal(t, uint64(1), receipt.Status, "transaction failed")
			
// 			time.Sleep(2 * time.Second)

// 			// mint 1 on contract1, which should trigger an update
// 			tx, err := contract.Mint(env.auth, address1, big.NewInt(1))
// 			require.NoError(t, err)
			
// 			receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
// 			require.NoError(t, err)
// 			require.Equal(t, uint64(1), receipt.Status, "transaction failed")

// 			return nil
// 		},
// 		// Handle update after mint on contract2
// 		func(msg ClientMessage) error {
// 			hasUpdates(t, msg, 1)
// 			hasStateUpdate(t, msg, 0, map[string]interface{}{
// 				"balances": map[string]interface{}{address1.Hex(): float64(101)},
// 			})

// 			// Verify final contract state
// 			balance, err := contract.Balances(nil, address1)
// 			if err != nil {
// 				return err
// 			}
// 			assert.Equal(t, int64(101), balance.Int64())

// 			return nil
// 		},
// 	})

// 	// Send subscription message
// 	subMsg := ClientMessage{
// 		Type: TypeSubscribe,
// 		Data: json.RawMessage(fmt.Sprintf(`{
// 			"pendingId": "%s",
// 			"eventSubscription": {
// 				"contractAddresses": [],
// 				"events": [{
// 					"name": "Transfer",
// 					"args": {
// 						"from": "0x0000000000000000000000000000000000000000"
// 					}
// 				}]
// 			}
// 		}`, pendingID)),
// 	}
// 	require.NoError(t, env.ws.WriteJSON(subMsg))

// 	// Wait for completion or timeout
// 	select {
// 	case <-env.done:
// 	case err := <-env.errChan:
// 		t.Fatal(err)
// 	case <-time.After(30 * time.Second):
// 		t.Fatal("test timeout")
// 	}
// }

func TestStateAndEventSubscriptionCustomContractAndEvent(t *testing.T) {
	env, err := setupTestEnv(t)
	require.NoError(t, err)
	defer env.cleanup()

	// Register custom contract with etherbase
	etherbase := env.contracts["Etherbase"].Contract.(*etherbase.Etherbase)
	tx, err := etherbase.AddCustomContract(
		env.auth,
		env.contracts["ERC20"].ContractAddress,
		erc20.ERC20ABI,
	)
	require.NoError(t, err)
	
	receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	time.Sleep(5 * time.Second)

	pendingID := "test-pending-id-3"
	contract := env.contracts["ERC20"].Contract.(*erc20.ERC20)

	// get the address of the auth
	address1 := env.auth.From

	pk1, err := crypto.GenerateKey()
	require.NoError(t, err)
	address2 := crypto.PubkeyToAddress(pk1.PublicKey)

	// Mint 100 tokens to address1
	tx, err = contract.Mint(env.auth, address1, big.NewInt(100))
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	// Setup message handlers
	env.handleMessages([]func(ClientMessage) error{
		// Handle subscription confirmation
		func(msg ClientMessage) error {
			isSubscriptionSuccess(t, msg, pendingID)
			return nil
		},
		// Handle initial state
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"balances": map[string]interface{}{address1.Hex(): float64(100), address2.Hex(): float64(0)},
			})

			time.Sleep(2 * time.Second)

			tx, err := contract.Transfer(env.auth, address2, big.NewInt(10))
			require.NoError(t, err)
			
			receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status, "transaction failed")
			return nil
		},
		// Handle state update after transfer
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 2)
			hasEventUpdate(t, msg, 0, map[string]interface{}{
				"contractAddress": env.contracts["ERC20"].ContractAddress.Hex(),
				"name": "Transfer",
				"args": map[string]interface{}{
					"from": address1.Hex(),
					"to": address2.Hex(),
					"value": float64(10),
				},
			})
			hasStateUpdate(t, msg, 1, map[string]interface{}{
				"balances": map[string]interface{}{address1.Hex(): float64(90), address2.Hex(): float64(10)},
			})

			// Verify final contract state
			balance1, err := contract.Balances(nil, address1)
			require.NoError(t, err)
			assert.Equal(t, int64(90), balance1.Int64())	
			balance2, err := contract.Balances(nil, address2)
			require.NoError(t, err)
			assert.Equal(t, int64(10), balance2.Int64())

			// mint 100 tokens to address2
			tx, err = contract.Mint(env.auth, address2, big.NewInt(400))
			require.NoError(t, err)
			receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status, "transaction failed")
			return nil
		},
		// Handle state update after mint
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 2)
			hasEventUpdate(t, msg, 0, map[string]interface{}{
				"contractAddress": env.contracts["ERC20"].ContractAddress.Hex(),
				"name": "Transfer",
				"args": map[string]interface{}{
					"from": "0x0000000000000000000000000000000000000000",
					"to": address2.Hex(),
					"value": float64(400),
				},
			})
			// we don't get address1 as it was not updated
			hasStateUpdate(t, msg, 1, map[string]interface{}{
				"balances": map[string]interface{}{address2.Hex(): float64(410)},
			})
			return nil
		},
	})

	subMsg := ClientMessage{
		Type: TypeSubscribe,
		Data: json.RawMessage(fmt.Sprintf(`{
			"pendingId": "%s",
			"stateSubscription": {
				"contractAddress": "%s",
				"statePath": ["balances", ["%s", "%s"]]
			},
			"eventSubscription": {
				"contractAddresses": ["%s"],
				"events": [{
					"name": "Transfer"
				}]
			}
		}`, pendingID, env.contracts["ERC20"].ContractAddress.Hex(), address1.Hex(), address2.Hex(), env.contracts["ERC20"].ContractAddress.Hex())),
	}
	require.NoError(t, env.ws.WriteJSON(subMsg))

	// Wait for completion or timeout
	select {
	case <-env.done:
	case err := <-env.errChan:
		t.Fatal(err)
	case <-time.After(30 * time.Second):
		t.Fatal("test timeout")
	}
}

func TestStateSubscriptionCustomContractAndEventAndGranularArgsNoDeltaUpdate(t *testing.T) {
	env, err := setupTestEnv(t)
	require.NoError(t, err)
	defer env.cleanup()

	// Register custom contract with etherbase
	etherbase := env.contracts["Etherbase"].Contract.(*etherbase.Etherbase)
	tx, err := etherbase.AddCustomContract(
		env.auth,
		env.contracts["ERC20"].ContractAddress,
		erc20.ERC20ABI,
	)
	require.NoError(t, err)
	
	receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	time.Sleep(5 * time.Second)

	pendingID := "test-pending-id-3"
	contract := env.contracts["ERC20"].Contract.(*erc20.ERC20)


	// get the address of the auth
	address1 := env.auth.From

	pk1, err := crypto.GenerateKey()
	require.NoError(t, err)
	address2 := crypto.PubkeyToAddress(pk1.PublicKey)

	pk2, err := crypto.GenerateKey()
	require.NoError(t, err)
	address3 := crypto.PubkeyToAddress(pk2.PublicKey)


	// Mint 100 tokens to address1
	tx, err = contract.Mint(env.auth, address1, big.NewInt(100))
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	// Mint 50 tokens to address3
	tx, err = contract.Mint(env.auth, address3, big.NewInt(50))
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	// Setup message handlers
	env.handleMessages([]func(ClientMessage) error{
		// Handle subscription confirmation
		func(msg ClientMessage) error {
			isSubscriptionSuccess(t, msg, pendingID)
			return nil
		},
		// Handle initial state
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"balances": map[string]interface{}{address1.Hex(): float64(100), address2.Hex(): float64(0), address3.Hex(): float64(50)},
			})

			time.Sleep(2 * time.Second)

			tx, err := contract.Transfer(env.auth, address2, big.NewInt(10))
			require.NoError(t, err)
			
			receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status, "transaction failed")
			return nil
		},
		// Handle state update after transfer
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"balances": map[string]interface{}{address1.Hex(): float64(90), address2.Hex(): float64(10), address3.Hex(): float64(50)},
			})

			// Verify final contract state
			balance1, err := contract.Balances(nil, address1)
			require.NoError(t, err)
			assert.Equal(t, int64(90), balance1.Int64())	
			balance2, err := contract.Balances(nil, address2)
			require.NoError(t, err)
			assert.Equal(t, int64(10), balance2.Int64())
			return nil
		},
	})

	subMsg := ClientMessage{
		Type: TypeSubscribe,
		Data: json.RawMessage(fmt.Sprintf(`{
			"pendingId": "%s",
			"stateSubscription": {
				"contractAddress": "%s",
				"statePath": ["balances", ["%s", "%s", "%s"]],
				"options": {
					"repoll": {
						"listenEvents": [{
							"name": "Transfer",
							"args": {
								"from": ["%s", "%s", "%s"],
								"to": ["%s", "%s", "%s"]
							}
						}]
					},
					"fullStateUpdateOnChange": true
				}
			}
		}`, pendingID, env.contracts["ERC20"].ContractAddress.Hex(), address1.Hex(), address2.Hex(), address3.Hex(), address1.Hex(), address2.Hex(), address3.Hex(), address1.Hex(), address2.Hex(), address3.Hex())),
	}
	require.NoError(t, env.ws.WriteJSON(subMsg))

	// Wait for completion or timeout
	select {
	case <-env.done:
	case err := <-env.errChan:
		t.Fatal(err)
	case <-time.After(30 * time.Second):
		t.Fatal("test timeout")
	}
}

func TestStateSubscriptionCustomContractAndEventAndGranularArgsNoDeltaPreciseUpdate(t *testing.T) {
	env, err := setupTestEnv(t)
	require.NoError(t, err)
	defer env.cleanup()

	// Register custom contract with etherbase
	etherbase := env.contracts["Etherbase"].Contract.(*etherbase.Etherbase)
	tx, err := etherbase.AddCustomContract(
		env.auth,
		env.contracts["ERC20"].ContractAddress,
		erc20.ERC20ABI,
	)
	require.NoError(t, err)
	
	receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	time.Sleep(5 * time.Second)

	pendingID := "test-pending-id-3"
	contract := env.contracts["ERC20"].Contract.(*erc20.ERC20)


	// get the address of the auth
	address1 := env.auth.From

	pk1, err := crypto.GenerateKey()
	require.NoError(t, err)
	address2 := crypto.PubkeyToAddress(pk1.PublicKey)

	pk2, err := crypto.GenerateKey()
	require.NoError(t, err)
	address3 := crypto.PubkeyToAddress(pk2.PublicKey)


	// Mint 100 tokens to address1
	tx, err = contract.Mint(env.auth, address1, big.NewInt(100))
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	// Mint 50 tokens to address3
	tx, err = contract.Mint(env.auth, address3, big.NewInt(50))
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	var mintBlock uint64
	var transferBlock uint64

	// Setup message handlers
	env.handleMessages([]func(ClientMessage) error{
				// Handle subscription confirmation
		func(msg ClientMessage) error {
				isSubscriptionSuccess(t, msg, pendingID)
			return nil
		},
		// Handle initial state
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"balances": map[string]interface{}{address1.Hex(): float64(100), address2.Hex(): float64(0), address3.Hex(): float64(50)},
			})

			tx, err := contract.Mint(env.auth, address3, big.NewInt(22))
			require.NoError(t, err)
			receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status, "transaction failed")
			mintBlock = receipt.BlockNumber.Uint64()
			time.Sleep(2 * time.Second)

			tx, err = contract.Transfer(env.auth, address2, big.NewInt(10))
			require.NoError(t, err)
			
			receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status, "transaction failed")
			transferBlock = receipt.BlockNumber.Uint64()
			return nil
		},
		// Handle state update after transfer
		func(msg ClientMessage) error {
			// we expect to get the update from the transfer
			assert.Greater(t, transferBlock, mintBlock)
			hasBlockNumber(t, msg, transferBlock)

			hasUpdates(t, msg, 1)
			// address3 is present as we read all the state we subscribe to when the event triggers
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"balances": map[string]interface{}{address1.Hex(): float64(90), address2.Hex(): float64(10), address3.Hex(): float64(72)},
			})

			// Verify final contract state
			balance1, err := contract.Balances(nil, address1)
			require.NoError(t, err)
			assert.Equal(t, int64(90), balance1.Int64())	
			balance2, err := contract.Balances(nil, address2)
			require.NoError(t, err)
			assert.Equal(t, int64(10), balance2.Int64())
			return nil
		},
	})

	// Send subscription message - listen to address1, address2, and address3, but only watch for updates to address1 and address2
	subMsg := ClientMessage{
		Type: TypeSubscribe,
		Data: json.RawMessage(fmt.Sprintf(`{
			"pendingId": "%s",
			"stateSubscription": {
				"contractAddress": "%s",
				"statePath": ["balances", ["%s", "%s", "%s"]],
				"options": {
					"repoll": {
						"listenEvents": [{
							"name": "Transfer",
							"args": {
								"from": ["%s", "%s"],
								"to": ["%s", "%s"]
							}
						}]
					},
					"fullStateUpdateOnChange": true
				}
			}
		}`, pendingID, env.contracts["ERC20"].ContractAddress.Hex(), address1.Hex(), address2.Hex(), address3.Hex(), address1.Hex(), address2.Hex(), address1.Hex(), address2.Hex())),
	}
	require.NoError(t, env.ws.WriteJSON(subMsg))

	// Wait for completion or timeout
	select {
	case <-env.done:
	case err := <-env.errChan:
		t.Fatal(err)
	case <-time.After(30 * time.Second):
		t.Fatal("test timeout")
	}
}

// the two players need to be given different delta updates based on the data they know about. 
func TestStateSubscriptionCustomContractAndEventAndGranularArgsDeltaUpdateMultiplePlayers(t *testing.T) {
	env, err := setupTestEnv(t)
	require.NoError(t, err)
	defer env.cleanup()

	// Register custom contract with etherbase
	etherbase := env.contracts["Etherbase"].Contract.(*etherbase.Etherbase)
	tx, err := etherbase.AddCustomContract(
		env.auth,
		env.contracts["ERC20"].ContractAddress,
		erc20.ERC20ABI,
	)
	require.NoError(t, err)
	
	receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	time.Sleep(5 * time.Second)

	pendingID := "test-pending-id-3"
	contract := env.contracts["ERC20"].Contract.(*erc20.ERC20)


	// get the address of the auth
	address1 := env.auth.From

	pk1, err := crypto.GenerateKey()
	require.NoError(t, err)
	address2 := crypto.PubkeyToAddress(pk1.PublicKey)

	pk2, err := crypto.GenerateKey()
	require.NoError(t, err)
	address3 := crypto.PubkeyToAddress(pk2.PublicKey)


	// Mint 100 tokens to address1
	tx, err = contract.Mint(env.auth, address1, big.NewInt(100))
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	// Mint 50 tokens to address3
	tx, err = contract.Mint(env.auth, address3, big.NewInt(50))
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	// Setup message handlers
	env.handleMessages([]func(ClientMessage) error{
		// Handle subscription confirmation
		func(msg ClientMessage) error {
			isSubscriptionSuccess(t, msg, pendingID)
			return nil
		},
		// Handle initial state
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"balances": map[string]interface{}{address1.Hex(): float64(100), address2.Hex(): float64(0), address3.Hex(): float64(50)},
			})

			time.Sleep(2 * time.Second)

			tx, err := contract.Transfer(env.auth, address2, big.NewInt(10))
			require.NoError(t, err)
			
			receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status, "transaction failed")
			return nil
		},
		// Handle state update after transfer
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			// address3 is not present as it was not updated, ie delta update
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"balances": map[string]interface{}{address1.Hex(): float64(90), address2.Hex(): float64(10)},
			})

			// Verify final contract state
			balance1, err := contract.Balances(nil, address1)
			require.NoError(t, err)
			assert.Equal(t, int64(90), balance1.Int64())	
			balance2, err := contract.Balances(nil, address2)
			require.NoError(t, err)
			assert.Equal(t, int64(10), balance2.Int64())
			return nil
		},
	})

	subMsg := ClientMessage{
		Type: TypeSubscribe,
		Data: json.RawMessage(fmt.Sprintf(`{
			"pendingId": "%s",
			"stateSubscription": {
				"contractAddress": "%s",
				"statePath": ["balances", ["%s", "%s", "%s"]],
				"options": {
					"repoll": {
						"listenEvents": [{
							"name": "Transfer",
							"args": {
								"from": ["%s", "%s", "%s"],
								"to": ["%s", "%s", "%s"]
							}
						}]
					},
					"fullStateUpdateOnChange": false
				}
			}
		}`, pendingID, env.contracts["ERC20"].ContractAddress.Hex(), address1.Hex(), address2.Hex(), address3.Hex(), address1.Hex(), address2.Hex(), address3.Hex(), address1.Hex(), address2.Hex(), address3.Hex())),
	}
	require.NoError(t, env.ws.WriteJSON(subMsg))

	// Wait for completion or timeout
	select {
	case <-env.done:
	case err := <-env.errChan:
		t.Fatal(err)
	case <-time.After(30 * time.Second):
		t.Fatal("test timeout")
	}
}

func TestStateSubscriptionCustomContractAndEventAndGranularArgsDeltaUpdate(t *testing.T) {
	env, err := setupTestEnv(t)
	require.NoError(t, err)
	defer env.cleanup()

	// Register custom contract with etherbase
	etherbase := env.contracts["Etherbase"].Contract.(*etherbase.Etherbase)
	tx, err := etherbase.AddCustomContract(
		env.auth,
		env.contracts["ERC20"].ContractAddress,
		erc20.ERC20ABI,
	)
	require.NoError(t, err)
	
	receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	time.Sleep(5 * time.Second)

	pendingID := "test-pending-id-3"
	contract := env.contracts["ERC20"].Contract.(*erc20.ERC20)


	// get the address of the auth
	address1 := env.auth.From

	pk1, err := crypto.GenerateKey()
	require.NoError(t, err)
	address2 := crypto.PubkeyToAddress(pk1.PublicKey)

	pk2, err := crypto.GenerateKey()
	require.NoError(t, err)
	address3 := crypto.PubkeyToAddress(pk2.PublicKey)


	// Mint 100 tokens to address1
	tx, err = contract.Mint(env.auth, address1, big.NewInt(100))
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	// Mint 50 tokens to address3
	tx, err = contract.Mint(env.auth, address3, big.NewInt(50))
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	// Setup message handlers
	env.handleMessages([]func(ClientMessage) error{
		// Handle subscription confirmation
		func(msg ClientMessage) error {
			isSubscriptionSuccess(t, msg, pendingID)
			return nil
		},
		// Handle initial state
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"balances": map[string]interface{}{address1.Hex(): float64(100), address2.Hex(): float64(0), address3.Hex(): float64(50)},
			})

			time.Sleep(2 * time.Second)

			tx, err := contract.Transfer(env.auth, address2, big.NewInt(10))
			require.NoError(t, err)
			
			receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status, "transaction failed")
			return nil
		},
		// Handle state update after transfer
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			// address3 is not present as it was not updated, ie delta update
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"balances": map[string]interface{}{address1.Hex(): float64(90), address2.Hex(): float64(10)},
			})

			// Verify final contract state
			balance1, err := contract.Balances(nil, address1)
			require.NoError(t, err)
			assert.Equal(t, int64(90), balance1.Int64())	
			balance2, err := contract.Balances(nil, address2)
			require.NoError(t, err)
			assert.Equal(t, int64(10), balance2.Int64())
			return nil
		},
	})

	subMsg := ClientMessage{
		Type: TypeSubscribe,
		Data: json.RawMessage(fmt.Sprintf(`{
			"pendingId": "%s",
			"stateSubscription": {
				"contractAddress": "%s",
				"statePath": ["balances", ["%s", "%s", "%s"]],
				"options": {
					"repoll": {
						"listenEvents": [{
							"name": "Transfer",
							"args": {
								"from": ["%s", "%s", "%s"],
								"to": ["%s", "%s", "%s"]
							}
						}]
					},
					"fullStateUpdateOnChange": false
				}
			}
		}`, pendingID, env.contracts["ERC20"].ContractAddress.Hex(), address1.Hex(), address2.Hex(), address3.Hex(), address1.Hex(), address2.Hex(), address3.Hex(), address1.Hex(), address2.Hex(), address3.Hex())),
	}
	require.NoError(t, env.ws.WriteJSON(subMsg))

	// Wait for completion or timeout
	select {
	case <-env.done:
	case err := <-env.errChan:
		t.Fatal(err)
	case <-time.After(30 * time.Second):
		t.Fatal("test timeout")
	}
}

func TestStateSubscriptionWithEtherbaseStateChangeMultipleValues(t *testing.T) {
	env, err := setupTestEnv(t)
	require.NoError(t, err)
	defer env.cleanup()

	testAddress, _, test, err := publicethdbupdateevent.DeployPublicVariableEthDBUpdateUintEvent(env.auth, env.ethClient)
	if err != nil {
		t.Fatalf("failed to deploy PublicVariableCustomStateEvent: %v", err)
	}
	env.contracts["PublicVariableEthDBUpdateUintEvent"] = &DeployedContract{
		ContractAddress: testAddress,
		Contract:       test,
	}

	// Register custom contract with etherbase
	etherbase := env.contracts["Etherbase"].Contract.(*etherbase.Etherbase)
	tx, err := etherbase.AddCustomContract(
		env.auth,
		env.contracts["PublicVariableEthDBUpdateUintEvent"].ContractAddress,
		publicethdbupdateevent.PublicVariableEthDBUpdateUintEventABI,
	)
	require.NoError(t, err)
	
	receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "transaction failed")

	time.Sleep(5 * time.Second)

	pendingID1 := "test-pending-id-3"
	contract := env.contracts["PublicVariableEthDBUpdateUintEvent"].Contract.(*publicethdbupdateevent.PublicVariableEthDBUpdateUintEvent)
	contractAddr := env.contracts["PublicVariableEthDBUpdateUintEvent"].ContractAddress

	// Setup message handlers
	env.handleMessages([]func(ClientMessage) error{
		// Handle subscription confirmation
		func(msg ClientMessage) error {
			isSubscriptionSuccess(t, msg, pendingID1)
			return nil
		},
		// Handle initial state
		func(msg ClientMessage) error {
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"valueA": float64(1),
				"valueB": float64(10),
			})

			// Verify initial contract state
			valueA, err := contract.ValueA(nil)
			require.NoError(t, err)
			assert.Equal(t, int64(1), valueA.Int64())
			
			valueB, err := contract.ValueB(nil)
			require.NoError(t, err)
			assert.Equal(t, int64(10), valueB.Int64())

			time.Sleep(2 * time.Second)

			// Call increment
			tx, err := contract.Increment(env.auth)
			require.NoError(t, err)
			
			receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status, "transaction failed")
			return nil
		},
		// Handle state update after increment
		func(msg ClientMessage) error {
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"valueA": float64(2),
				"valueB": float64(12),
			})

			// Verify final contract state
			valueA, err := contract.ValueA(nil)
			require.NoError(t, err)
			assert.Equal(t, int64(2), valueA.Int64())
			
			valueB, err := contract.ValueB(nil)
			require.NoError(t, err)
			assert.Equal(t, int64(12), valueB.Int64())
			return nil
		},
	})

	// Send subscription message
	subMsg1 := ClientMessage{
		Type: TypeSubscribe,
		Data: json.RawMessage(fmt.Sprintf(`{
			"pendingId": "%s",
			"stateSubscription": {
				"contractAddress": "%s",
				"statePath": [["valueA", "valueB"]]
			}
		}`, pendingID1, contractAddr.Hex())),
	}
	require.NoError(t, env.ws.WriteJSON(subMsg1))

	// Wait for completion or timeout
	select {
	case <-env.done:
	case err := <-env.errChan:
		t.Fatal(err)
	case <-time.After(30 * time.Second):
		t.Fatal("test timeout")
	}
}

func TestStateSubscriptionWithMappings(t *testing.T) {
	env, err := setupTestEnv(t)
	require.NoError(t, err)
	defer env.cleanup()

	testAddress, _, test, err := publicethdbupdateevent.DeployPublicVariableEthDBUpdateMappingEvent(env.auth, env.ethClient)
	if err != nil {
		t.Fatalf("failed to deploy PublicVariableEthDBUpdateMappingEvent: %v", err)
	}
	env.contracts["PublicVariableEthDBUpdateMappingEvent"] = &DeployedContract{
		ContractAddress: testAddress,
		Contract:       test,
	}

	// Register custom contract with etherbase
	etherbase := env.contracts["Etherbase"].Contract.(*etherbase.Etherbase)
	tx, err := etherbase.AddCustomContract(
		env.auth,
		env.contracts["PublicVariableEthDBUpdateMappingEvent"].ContractAddress,
		publicethdbupdateevent.PublicVariableEthDBUpdateMappingEventABI,
	)
	require.NoError(t, err)
	
	receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)

	time.Sleep(5 * time.Second)

	pendingID1 := "test-pending-id-3"
	contract := env.contracts["PublicVariableEthDBUpdateMappingEvent"].Contract.(*publicethdbupdateevent.PublicVariableEthDBUpdateMappingEvent)
	contractAddr := env.contracts["PublicVariableEthDBUpdateMappingEvent"].ContractAddress

	// Setup message handlers
	env.handleMessages([]func(ClientMessage) error{
		// Handle subscription confirmation
		func(msg ClientMessage) error {
			isSubscriptionSuccess(t, msg, pendingID1)
			return nil
		},
		// Handle initial state
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"mappingA": map[string]interface{}{
					"1": float64(1),
					"3": float64(0),
				},
			})

			// Verify initial contract state
			initialValueA1, err := contract.MappingA(nil, big.NewInt(1))
			require.NoError(t, err)
			assert.Equal(t, int64(1), initialValueA1.Int64())
			
			initialValueA3, err := contract.MappingA(nil, big.NewInt(3))
			require.NoError(t, err)
			assert.Equal(t, int64(0), initialValueA3.Int64())

			time.Sleep(2 * time.Second)

			// Call increment
			tx, err := contract.Increment(env.auth)
			if err != nil {
				return fmt.Errorf("failed to call increment: %v", err)
			}
			
			receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
			if err != nil {
				return fmt.Errorf("failed waiting for transaction: %v", err)
			}
			if receipt.Status != uint64(1) {
				return fmt.Errorf("increment transaction failed")
			}
			return nil
		},
		// Handle state update after increment
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"mappingA": map[string]interface{}{
					"1": float64(2),
					// we don't get 3 as it was not updated
				},
			})

			// Verify final state
			finalValueA1, err := contract.MappingA(nil, big.NewInt(1))
			if err != nil {
				return err
			}
			assert.Equal(t, int64(2), finalValueA1.Int64())
			
			finalValueA3, err := contract.MappingA(nil, big.NewInt(3))
			if err != nil {
				return err
			}
			assert.Equal(t, int64(0), finalValueA3.Int64())
			return nil
		},
	})

	// Send subscription message
	subMsg := ClientMessage{
		Type: TypeSubscribe,
		Data: json.RawMessage(fmt.Sprintf(`{
			"pendingId": "%s",
			"stateSubscription": {
				"contractAddress": "%s",
				"statePath": ["mappingA", ["1", "3"]]
			}
		}`, pendingID1, contractAddr.Hex())),
	}
	require.NoError(t, env.ws.WriteJSON(subMsg))

	// Wait for completion or timeout
	select {
	case <-env.done:
	case err := <-env.errChan:
		t.Fatal(err)
	case <-time.After(30 * time.Second):
		t.Fatal("test timeout")
	}
}

func TestStateSubscriptionWithNestedStructs(t *testing.T) {
	env, err := setupTestEnv(t)
	require.NoError(t, err)
	defer env.cleanup()

	source := env.contracts["EtherbaseSource"].Contract.(*etherbasesource.EtherbaseSource)
	sourceAddr := env.contracts["EtherbaseSource"].ContractAddress
	pendingID1 := "test-pending-id-5"

	// Initialize state
	aliceSegments := []string{"users", "alice", "counter"}
	aliceData, err := encodeAbiParameters([]string{"uint256"}, []interface{}{big.NewInt(1)})
	require.NoError(t, err)
	tx, err := source.SetValue(env.auth, aliceSegments, uint8(types.DataTypeUint256), aliceData)
	require.NoError(t, err)
	receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)

	bobSegments := []string{"users", "bob", "counter"}
	bobData, err := encodeAbiParameters([]string{"uint256"}, []interface{}{big.NewInt(5)})
	require.NoError(t, err)
	tx, err = source.SetValue(env.auth, bobSegments, uint8(types.DataTypeUint256), bobData)
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)

	charlieSegments := []string{"users", "charlie", "counter"}
	charlieData, err := encodeAbiParameters([]string{"uint256"}, []interface{}{big.NewInt(10)})
	require.NoError(t, err)
	tx, err = source.SetValue(env.auth, charlieSegments, uint8(types.DataTypeUint256), charlieData)
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)

	// Dummy transaction to bump block number
	tx, err = source.CreateWalletIdentity(env.auth, common.Address{}, []uint8{})
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)

	time.Sleep(5 * time.Second)

	// Setup message handlers
	env.handleMessages([]func(ClientMessage) error{
		// Handle subscription confirmation
		func(msg ClientMessage) error {
			isSubscriptionSuccess(t, msg, pendingID1)
			return nil
		},
		// Handle initial state
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"users": map[string]interface{}{
					"alice": map[string]interface{}{
						"counter": float64(1),
					},
					"bob": map[string]interface{}{
						"counter": float64(5),
					},
				},
			})

			time.Sleep(2 * time.Second)

			// Update alice's counter
			aliceData, err := encodeAbiParameters([]string{"uint256"}, []interface{}{big.NewInt(2)})
			require.NoError(t, err)
			tx, err := source.SetValue(env.auth, aliceSegments, uint8(types.DataTypeUint256), aliceData)
			require.NoError(t, err)
			receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status)
			return nil
		},
		// Handle state update
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"users": map[string]interface{}{
					"alice": map[string]interface{}{
						"counter": float64(2),
					},
				},
			})
			return nil
		},
	})

	// Send subscription message
	subMsg := ClientMessage{
		Type: TypeSubscribe,
		Data: json.RawMessage(fmt.Sprintf(`{
			"pendingId": "%s",
			"stateSubscription": {
				"contractAddress": "%s",
				"statePath": ["users", ["alice", "bob"]]
			},
			"options": {
				"repoll": {
					"onAnyContractEvent": false
				}
			}
		}`, pendingID1, sourceAddr.Hex())),
	}
	require.NoError(t, env.ws.WriteJSON(subMsg))

	// Wait for completion or timeout
	select {
	case <-env.done:
	case err := <-env.errChan: 
		t.Fatal(err)
	case <-time.After(30 * time.Second):
		t.Fatal("test timeout")
	}
}

// func TestStateSubscriptionCustomContractWithNoInitialData(t *testing.T) {
// 	env, err := setupTestEnv(t)
// 	require.NoError(t, err)
// 	defer env.cleanup()

// 	testAddress, _, test, err := publicethdbupdateevent.DeployPublicVariableEthDBUpdateMappingInitiallyEmptyEvent(env.auth, env.ethClient)
// 	if err != nil {
// 		t.Fatalf("failed to deploy PublicVariableEthDBUpdateMappingInitiallyEmptyEvent: %v", err)
// 	}
// 	env.contracts["PublicVariableEthDBUpdateMappingInitiallyEmptyEvent"] = &DeployedContract{
// 		ContractAddress: testAddress,
// 		Contract:       test,
// 	}

// 	// Register custom contract with etherbase
// 	etherbase := env.contracts["Etherbase"].Contract.(*etherbase.Etherbase)
// 	tx, err := etherbase.AddCustomContract(
// 		env.auth,
// 		env.contracts["PublicVariableEthDBUpdateMappingInitiallyEmptyEvent"].ContractAddress,
// 		publicethdbupdateevent.PublicVariableEthDBUpdateMappingInitiallyEmptyEventABI,
// 	)
// 	require.NoError(t, err)
	
// 	receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
// 	require.NoError(t, err)
// 	require.Equal(t, uint64(1), receipt.Status)

// 	time.Sleep(5 * time.Second)

// 	pendingID1 := "test-pending-id-3"
// 	contract := env.contracts["PublicVariableEthDBUpdateMappingInitiallyEmptyEvent"].Contract.(*publicethdbupdateevent.PublicVariableEthDBUpdateMappingInitiallyEmptyEvent)
// 	contractAddr := env.contracts["PublicVariableEthDBUpdateMappingInitiallyEmptyEvent"].ContractAddress

// 	// Setup message handlers
// 	env.handleMessages([]func(ClientMessage) error{
// 		// Handle subscription confirmation
// 		func(msg ClientMessage) error {
// 			isSubscriptionSuccess(t, msg, pendingID1)
// 			return nil
// 		},
// 		// Handle initial state
// 		func(msg ClientMessage) error {
// 			hasUpdates(t, msg, 1)
// 			hasStateUpdate(t, msg, 0, map[string]interface{}{
// 				"mappingA": map[string]interface{}{
// 					"1": float64(0),
// 					"3": float64(0),
// 				},
// 			})

// 			// Verify initial contract state
// 			initialValueA1, err := contract.MappingA(nil, big.NewInt(1))
// 			require.NoError(t, err)
// 			assert.Equal(t, int64(1), initialValueA1.Int64())
			
// 			initialValueA3, err := contract.MappingA(nil, big.NewInt(3))
// 			require.NoError(t, err)
// 			assert.Equal(t, int64(0), initialValueA3.Int64())

// 			time.Sleep(2 * time.Second)

// 			// Call increment
// 			tx, err := contract.Increment(env.auth)
// 			if err != nil {
// 				return fmt.Errorf("failed to call increment: %v", err)
// 			}
			
// 			receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
// 			if err != nil {
// 				return fmt.Errorf("failed waiting for transaction: %v", err)
// 			}
// 			if receipt.Status != uint64(1) {
// 				return fmt.Errorf("increment transaction failed")
// 			}
// 			return nil
// 		},
// 		// Handle state update after increment
// 		func(msg ClientMessage) error {
// 			hasUpdates(t, msg, 1)
// 			hasStateUpdate(t, msg, 0, map[string]interface{}{
// 				"mappingA": map[string]interface{}{
// 					"1": float64(2),
// 				},
// 			})

// 			// Verify final state
// 			finalValueA1, err := contract.MappingA(nil, big.NewInt(1))
// 			if err != nil {
// 				return err
// 			}
// 			assert.Equal(t, int64(2), finalValueA1.Int64())
			
// 			finalValueA3, err := contract.MappingA(nil, big.NewInt(3))
// 			if err != nil {
// 				return err
// 			}
// 			assert.Equal(t, int64(0), finalValueA3.Int64())
// 			return nil
// 		},
// 	})

// 	// Send subscription message
// 	subMsg := ClientMessage{
// 		Type: TypeSubscribe,
// 		Data: json.RawMessage(fmt.Sprintf(`{
// 			"pendingId": "%s",
// 			"stateSubscription": {
// 				"contractAddress": "%s",
// 				"statePath": ["mappingA", ["1", "3"]]
// 			}
// 		}`, pendingID1, contractAddr.Hex())),
// 	}
// 	require.NoError(t, env.ws.WriteJSON(subMsg))

// 	// Wait for completion or timeout
// 	select {
// 	case <-env.done:
// 	case err := <-env.errChan:
// 		t.Fatal(err)
// 	case <-time.After(30 * time.Second):
// 		t.Fatal("test timeout")
// 	}
// }

func TestStateSubscriptionWithNestedStructsWithGlobalSubscription(t *testing.T) {
	env, err := setupTestEnv(t)
	require.NoError(t, err)
	defer env.cleanup()

	source := env.contracts["EtherbaseSource"].Contract.(*etherbasesource.EtherbaseSource)
	sourceAddr := env.contracts["EtherbaseSource"].ContractAddress
	pendingID1 := "test-pending-id-5"

	// Initialize state
	aliceSegments := []string{"users", "alice", "counter"}
	aliceData, err := encodeAbiParameters([]string{"uint256"}, []interface{}{big.NewInt(1)})
	require.NoError(t, err)
	tx, err := source.SetValue(env.auth, aliceSegments, uint8(types.DataTypeUint256), aliceData)
	require.NoError(t, err)
	receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)

	bobSegments := []string{"users", "bob", "counter"}
	bobData, err := encodeAbiParameters([]string{"uint256"}, []interface{}{big.NewInt(5)})
	require.NoError(t, err)
	tx, err = source.SetValue(env.auth, bobSegments, uint8(types.DataTypeUint256), bobData)
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)

	numberSegments := []string{"numbers", "one"}
	numberData, err := encodeAbiParameters([]string{"uint256"}, []interface{}{big.NewInt(1)})
	require.NoError(t, err)
	tx, err = source.SetValue(env.auth, numberSegments, uint8(types.DataTypeUint256), numberData)
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)

	// Dummy transaction to bump block number
	tx, err = source.CreateWalletIdentity(env.auth, common.Address{}, []uint8{})
	require.NoError(t, err)
	receipt, err = bind.WaitMined(context.Background(), env.ethClient, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status)

	time.Sleep(5 * time.Second)

	// Setup message handlers
	env.handleMessages([]func(ClientMessage) error{
		// Handle subscription confirmation
		func(msg ClientMessage) error {
			isSubscriptionSuccess(t, msg, pendingID1)
			return nil
		},
		// Handle initial state
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"users": map[string]interface{}{
					"alice": map[string]interface{}{
						"counter": float64(1),
					},
					"bob": map[string]interface{}{
						"counter": float64(5),
					},
				},
				"numbers": map[string]interface{}{
					"one": float64(1),
				},
			})

			time.Sleep(2 * time.Second)

			// Update alice's counter
			aliceData, err := encodeAbiParameters([]string{"uint256"}, []interface{}{big.NewInt(2)})
			require.NoError(t, err)
			tx, err := source.SetValue(env.auth, aliceSegments, uint8(types.DataTypeUint256), aliceData)
			require.NoError(t, err)
			receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
			require.NoError(t, err)
			require.Equal(t, uint64(1), receipt.Status)
			return nil
		},
		// Handle state update
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"users": map[string]interface{}{
					"alice": map[string]interface{}{
						"counter": float64(2),
					},
				},
			})
			return nil
		},
	})

	// Send subscription message
	subMsg := ClientMessage{
		Type: TypeSubscribe,
		Data: json.RawMessage(fmt.Sprintf(`{
			"pendingId": "%s",
			"stateSubscription": {
				"contractAddress": "%s",
				"statePath": []
			}
		}`, pendingID1, sourceAddr.Hex())),
	}
	require.NoError(t, env.ws.WriteJSON(subMsg))

	// Wait for completion or timeout
	select {
	case <-env.done:
	case err := <-env.errChan:
		t.Fatal(err)
	case <-time.After(30 * time.Second):
		t.Fatal("test timeout")
	}
}

// func TestSourceCreatedAndAndSchemaRegisteredBeforeAndAfterServerStartWorks(t *testing.T) {
// 	env := &testEnv{
// 		t: t,
// 		done: make(chan bool),
// 		errChan: make(chan error),
// 	}

// 	// Deploy test contracts
// 	contracts, ethClient, err := deployTestContracts()
// 	require.NoError(t, err)

// 	// Get auth
// 	key, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
// 	require.NoError(t, err)
// 	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(31337))
// 	require.NoError(t, err)

// 	// Create a new source before server start
// 	etherbase := contracts["Etherbase"].Contract.(*etherbase.Etherbase)
// 	sourceTx, err := etherbase.CreateSource(auth)
// 	require.NoError(t, err)
// 	sourceReceipt, err := bind.WaitMined(context.Background(), ethClient, sourceTx)
// 	require.NoError(t, err)
// 	require.Equal(t, uint64(1), sourceReceipt.Status)
// 	sourceAddress := sourceReceipt.Logs[0].Address
// 	source, err := etherbasesource.NewEtherbaseSource(sourceAddress, ethClient)
// 	require.NoError(t, err)

// 	// Register first event schema before server start
// 	event1Name := "TestEvent1"
// 	event1Def := EventDefinition{
// 		Name: event1Name,
// 		Args: []Argument{
// 			{Name: "value", ArgType: "uint256", IsIndexed: false},
// 		},
// 	}
// 	event1ABI := registerEvent(t, ethClient, source, event1Def)

// 	// Setup server
// 	// Setup auth
// 	key, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
// 	require.NoError(t, err)
// 	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(31337))
// 	require.NoError(t, err)
// 	env.auth = auth

// 	// Setup server
// 	cfg, err := config.LoadConfig()
// 	require.NoError(t, err)
// 	cfg.EtherbaseAddress = contracts["Etherbase"].ContractAddress
// 	cfg.MulticallAddress = contracts["Multicall3"].ContractAddress
// 	cfg.RpcURL = "ws://localhost:8545"

// 	// Setup reader server
// 	readerSrv, err := reader.NewServer(cfg)
// 	require.NoError(t, err)

// 	env.server = httptest.NewServer(readerSrv.Router())
// 	go readerSrv.Start()

// 	// Setup writer server
// 	writerSrv, err := writer.NewServer(cfg)
// 	require.NoError(t, err)

// 	env.writerServer = httptest.NewServer(writerSrv.Router())
// 	go writerSrv.Start()

// 	time.Sleep(5 * time.Second)

// 	// Setup WebSocket connections
// 	wsURL := "ws" + strings.TrimPrefix(env.server.URL, "http") + "/read"
// 	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
// 	require.NoError(t, err)
// 	env.ws = ws

// 	wsWriterURL := "ws" + strings.TrimPrefix(env.writerServer.URL, "http") + "/write"
// 	wsWriter, _, err := websocket.DefaultDialer.Dial(wsWriterURL, nil)
// 	require.NoError(t, err)
// 	env.wsWriter = wsWriter
// 	time.Sleep(5 * time.Second)

// 	// Register second event schema after server start
// 	event2Name := "TestEvent2"
// 	event2Def := EventDefinition{
// 		Name: event2Name,
// 		Args: []Argument{
// 			{Name: "value", ArgType: "uint256", IsIndexed: false},
// 		},
// 	}
// 	event2ABI := registerEvent(t, ethClient, source, event2Def)

// 	time.Sleep(5 * time.Second)

// 	// Setup WebSocket connection
// 	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/read"
// 	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
// 	require.NoError(t, err)
// 	defer ws.Close()

// 	pendingID := "test-pending-id"
// 	done := make(chan bool)
// 	errChan := make(chan error)

// 	// Setup message handlers
// 	messageCount := 0
// 	go func() {
// 		for {
// 			var msg ClientMessage
// 			if err := ws.ReadJSON(&msg); err != nil {
// 				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
// 					errChan <- fmt.Errorf("websocket error: %v", err)
// 				}
// 				return
// 			}

// 			if msg.Type == TypeError || msg.Type == TypeSubscriptionFailed {
// 				errChan <- fmt.Errorf("error: %v", string(msg.Data))
// 				return
// 			}

// 			switch messageCount {
// 			case 0:
// 				// Handle subscription confirmation
// 				isSubscriptionSuccess(t, msg, pendingID)

// 				// Emit first event
// 				data1, err := event1ABI.Inputs.Pack(big.NewInt(100))
// 				require.NoError(t, err)
// 				tx, err := source.EmitEvent(auth, event1Name, [][32]byte{}, data1)
// 				require.NoError(t, err)
// 				receipt, err := bind.WaitMined(context.Background(), ethClient, tx)
// 				require.NoError(t, err)
// 				require.Equal(t, uint64(1), receipt.Status)

// 				// Emit second event
// 				data2, err := event2ABI.Inputs.Pack(big.NewInt(200))
// 				require.NoError(t, err)
// 				tx, err = source.EmitEvent(auth, event2Name, [][32]byte{}, data2)
// 				require.NoError(t, err)
// 				receipt, err = bind.WaitMined(context.Background(), ethClient, tx)
// 				require.NoError(t, err)
// 				require.Equal(t, uint64(1), receipt.Status)

// 			case 1:
// 				// Handle first event update
// 				hasUpdates(t, msg, 1)
// 				hasEventUpdate(t, msg, 0, map[string]interface{}{
// 					"contractAddress": sourceAddress.Hex(),
// 					"name":           event1Name,
// 					"args": map[string]interface{}{
// 						"value": float64(100),
// 					},
// 				})

// 			case 2:
// 				// Handle second event update
// 				hasUpdates(t, msg, 1)
// 				hasEventUpdate(t, msg, 0, map[string]interface{}{
// 					"contractAddress": sourceAddress.Hex(),
// 					"name":           event2Name,
// 					"args": map[string]interface{}{
// 						"value": float64(200),
// 					},
// 				})
// 				done <- true
// 				return
// 			}
// 			messageCount++
// 		}
// 	}()

// 	// Send subscription message for both events
// 	subMsg := ClientMessage{
// 		Type: TypeSubscribe,
// 		Data: json.RawMessage(fmt.Sprintf(`{
// 			"pendingId": "%s",
// 			"eventSubscription": {
// 				"contractAddress": "%s",
// 				"events": [
// 					{"name": "%s"},
// 					{"name": "%s"}
// 				]
// 			}
// 		}`, pendingID, sourceAddress.Hex(), event1Name, event2Name)),
// 	}
// 	require.NoError(t, ws.WriteJSON(subMsg))

// 	// Wait for completion or timeout
// 	select {
// 	case <-done:
// 	case err := <-errChan:
// 		t.Fatal(err)
// 	case <-time.After(30 * time.Second):
// 		t.Fatal("test timeout")
// 	}
// }


func TestStateSubscriptionSourceAddUpdateRemoveState(t *testing.T) {
	env, err := setupTestEnv(t)
	require.NoError(t, err)
	defer env.cleanup()

	pendingID1 := "test-pending-id-3"
	contract := env.contracts["EtherbaseSource"].Contract.(*etherbasesource.EtherbaseSource)
	contractAddr := env.contracts["EtherbaseSource"].ContractAddress

	aliceBool := true
	aliceBoolData, err := encodeAbiParameters([]string{"bool"}, []interface{}{aliceBool})
	require.NoError(t, err)
	aliceBytes := []byte("hello")
	aliceBytesData, err := encodeAbiParameters([]string{"bytes"}, []interface{}{aliceBytes})
	require.NoError(t, err)
	aliceString := "hello"
	aliceStringData, err := encodeAbiParameters([]string{"string"}, []interface{}{aliceString})
	require.NoError(t, err)
	aliceUint256 := float64(1)
	newAliceUint256 := float64(2)
	newAliceUint256Data, err := encodeAbiParameters([]string{"uint256"}, []interface{}{big.NewInt(int64(newAliceUint256))})
	require.NoError(t, err)
	aliceUint256Data, err := encodeAbiParameters([]string{"uint256"}, []interface{}{big.NewInt(int64(aliceUint256))})
	require.NoError(t, err)
	aliceInt256 := float64(-1)
	aliceInt256Data, err := encodeAbiParameters([]string{"int256"}, []interface{}{big.NewInt(int64(aliceInt256))})
	require.NoError(t, err)
	bobBool := false
	bobBoolData, err := encodeAbiParameters([]string{"bool"}, []interface{}{bobBool})
	require.NoError(t, err)
	bobBytes := []byte("world")
	bobBytesData, err := encodeAbiParameters([]string{"bytes"}, []interface{}{bobBytes})
	require.NoError(t, err)
	

	// tx, err := contract.SetValues(env.auth, []etherbasesource.EtherDatabaseLibBatchSetValue{
	// 	{
	// 		Segments: []string{"users", "alice", "bool"},
	// 		DataType: uint8(types.DataTypeBool),
	// 		Data:     aliceBoolData,
	// 	},
	// })
	// require.NoError(t, err)
	// receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
	// require.NoError(t, err)
	// require.Equal(t, uint64(1), receipt.Status)

	require.NoError(t, env.wsWriter.WriteJSON(ClientMessage{
		Type: TypeSetValue,
		Data: json.RawMessage(fmt.Sprintf(`{
			"contractAddress": "%s",
			"state": {
				"users": {
					"alice": {
						"bool": %t
					}
				}
			}
		}`, contractAddr.Hex(), aliceBool)),
	}))

	// todo test on getting back tx hash from tx pool to client sender
	// todo change all the other implementations to use this endpoint, same with emit event
	time.Sleep(5 * time.Second)

	// Setup message handlers
	env.handleMessages([]func(ClientMessage) error{
		// Handle subscription confirmation
		func(msg ClientMessage) error {
			isSubscriptionSuccess(t, msg, pendingID1)
			return nil
		},
		// Handle initial state
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"users": map[string]interface{}{
					"alice": map[string]interface{}{
						"bool": aliceBool,
					},
				},
			})

			aliceBoolTypeActual, aliceBoolDataActual, err := contract.GetValue(nil, []string{"users", "alice", "bool"})
			require.NoError(t, err)
			require.Equal(t, uint8(types.DataTypeBool), aliceBoolTypeActual)
			require.Equal(t, aliceBoolData, aliceBoolDataActual)

			require.NoError(t, env.wsWriter.WriteJSON(ClientMessage{
				Type: TypeSetValue,
				Data: json.RawMessage(fmt.Sprintf(`{
					"contractAddress": "%s",
					"state": {
						"users": {
							"alice": {
								"bytes": "%s",
								"string": "%s",
								"uint256": %f,
								"int256": %f,
							},
							"bob": {
								"bytes": "%s",
								"bool": %t
							}
						}
					}
				}`, contractAddr.Hex(), base64.StdEncoding.EncodeToString(aliceBytes), aliceString, aliceUint256, aliceInt256, base64.StdEncoding.EncodeToString(bobBytes), bobBool)),
			}))
			return nil
		},
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"users": map[string]interface{}{
					"alice": map[string]interface{}{
						"bytes": base64.StdEncoding.EncodeToString(aliceBytes),
						"string": aliceString,
						"uint256": aliceUint256,
						"int256": aliceInt256,
					},
					"bob": map[string]interface{}{
						"bytes": base64.StdEncoding.EncodeToString(bobBytes),
						"bool": bobBool,
					},
				},
			})
			
			aliceBytesTypeActual, aliceBytesDataActual, err := contract.GetValue(nil, []string{"users", "alice", "bytes"})
			require.NoError(t, err)
			require.Equal(t, uint8(types.DataTypeBytes), aliceBytesTypeActual)
			require.Equal(t, aliceBytesData, aliceBytesDataActual)

			aliceStringTypeActual, aliceStringDataActual, err := contract.GetValue(nil, []string{"users", "alice", "string"})
			require.NoError(t, err)
			require.Equal(t, uint8(types.DataTypeString), aliceStringTypeActual)
			require.Equal(t, aliceStringData, aliceStringDataActual)

			aliceUint256TypeActual, aliceUint256DataActual, err := contract.GetValue(nil, []string{"users", "alice", "uint256"})
			require.NoError(t, err)
			require.Equal(t, uint8(types.DataTypeUint256), aliceUint256TypeActual)
			require.Equal(t, aliceUint256Data, aliceUint256DataActual)

			aliceInt256TypeActual, aliceInt256DataActual, err := contract.GetValue(nil, []string{"users", "alice", "int256"})
			require.NoError(t, err)
			require.Equal(t, uint8(types.DataTypeInt256), aliceInt256TypeActual)
			require.Equal(t, aliceInt256Data, aliceInt256DataActual)
			
			bobBytesTypeActual, bobBytesDataActual, err := contract.GetValue(nil, []string{"users", "bob", "bytes"})
			require.NoError(t, err)
			require.Equal(t, uint8(types.DataTypeBytes), bobBytesTypeActual)
			require.Equal(t, bobBytesData, bobBytesDataActual)

			bobBoolTypeActual, bobBoolDataActual, err := contract.GetValue(nil, []string{"users", "bob", "bool"})
			require.NoError(t, err)
			require.Equal(t, uint8(types.DataTypeBool), bobBoolTypeActual)
			require.Equal(t, bobBoolData, bobBoolDataActual)

			// update a field in alice
			// tx, err := contract.SetValues(env.auth, []etherbasesource.EtherDatabaseLibBatchSetValue{
			// 	{
			// 		Segments: []string{"users", "alice", "uint256"},
			// 		DataType: uint8(types.DataTypeUint256),
			// 		Data:     newAliceUint256Data,
			// 	},
			// })
			// require.NoError(t, err)
			// receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
			// require.NoError(t, err)
			// require.Equal(t, uint64(1), receipt.Status)

			require.NoError(t, env.wsWriter.WriteJSON(ClientMessage{
				Type: TypeSetValue,
				Data: json.RawMessage(fmt.Sprintf(`{
					"contractAddress": "%s",
					"state": {
						"users": {
							"alice": {
								"uint256": %f,
							}
						}
					}
				}`, contractAddr.Hex(), newAliceUint256)),
			}))
			
			return nil
		},
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"users": map[string]interface{}{
					"alice": map[string]interface{}{
						"uint256": newAliceUint256,
					},
				},
			})

			aliceUint256TypeActual, aliceUint256DataActual, err := contract.GetValue(nil, []string{"users", "alice", "uint256"})
			require.NoError(t, err)
			require.Equal(t, uint8(types.DataTypeUint256), aliceUint256TypeActual)
			require.Equal(t, newAliceUint256Data, aliceUint256DataActual)

			// remove a field from alice
			// tx, err := contract.SetValues(env.auth, []etherbasesource.EtherDatabaseLibBatchSetValue{
			// 	{
			// 		Segments: []string{"users", "alice", "bool"},
			// 		DataType: uint8(types.DataTypeNone),
			// 		Data:     nil,
			// 	},
			// })
			// require.NoError(t, err)
			// receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
			// require.NoError(t, err)
			// require.Equal(t, uint64(1), receipt.Status)

			require.NoError(t, env.wsWriter.WriteJSON(ClientMessage{
				Type: TypeSetValue,
				Data: json.RawMessage(fmt.Sprintf(`{
					"contractAddress": "%s",
					"state": {
						"users": {
							"alice": {
								"bool": null,
							}
						}
					}
				}`, contractAddr.Hex())),
			}))
			
			return nil
		},
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"users": map[string]interface{}{
					"alice": map[string]interface{}{
						"bool": nil,
					},
				},
			})

			aliceBoolTypeActual, aliceBoolDataActual, err := contract.GetValue(nil, []string{"users", "alice", "bool"})
			require.NoError(t, err)
			require.Equal(t, uint8(types.DataTypeNone), aliceBoolTypeActual)
			require.Equal(t, aliceBoolDataActual, []byte{})


			bobBytesTypeActual, bobBytesDataActual, err := contract.GetValue(nil, []string{"users", "bob", "bytes"})
			require.NoError(t, err)
			require.Equal(t, uint8(types.DataTypeBytes), bobBytesTypeActual)
			require.Equal(t, bobBytesData, bobBytesDataActual)

			// remove alice entirely
			// tx, err := contract.SetValues(env.auth, []etherbasesource.EtherDatabaseLibBatchSetValue{
			// 	{
			// 		Segments: []string{"users", "alice"},
			// 		DataType: uint8(types.DataTypeNone),
			// 		Data:     nil,
			// 	},
			// })
			// require.NoError(t, err)
			// receipt, err := bind.WaitMined(context.Background(), env.ethClient, tx)
			// require.NoError(t, err)
			// require.Equal(t, uint64(1), receipt.Status)

			env.wsWriter.WriteJSON(ClientMessage{
				Type: TypeSetValue,
				Data: json.RawMessage(fmt.Sprintf(`{
					"contractAddress": "%s",
					"state": {
						"users": {
							"alice": null,
						}
					}
				}`, contractAddr.Hex())),
			})
			
			
			return nil
		},
		func(msg ClientMessage) error {
			hasUpdates(t, msg, 1)
			hasStateUpdate(t, msg, 0, map[string]interface{}{
				"users": map[string]interface{}{
					"alice": nil,
				},
			})
			aliceBytesTypeActual, aliceBytesDataActual, err := contract.GetValue(nil, []string{"users", "alice", "bytes"})
			require.NoError(t, err)
			require.Equal(t, uint8(types.DataTypeNone), aliceBytesTypeActual)
			require.Equal(t, aliceBytesDataActual, []byte{})

			aliceStringTypeActual, aliceStringDataActual, err := contract.GetValue(nil, []string{"users", "alice", "string"})
			require.NoError(t, err)
			require.Equal(t, uint8(types.DataTypeNone), aliceStringTypeActual)
			require.Equal(t, aliceStringDataActual, []byte{})

			aliceUint256TypeActual, aliceUint256DataActual, err := contract.GetValue(nil, []string{"users", "alice", "uint256"})
			require.NoError(t, err)
			require.Equal(t, uint8(types.DataTypeNone), aliceUint256TypeActual)
			require.Equal(t, aliceUint256DataActual, []byte{})			

			return nil
		},
	})

	// Send subscription message
	subMsg := ClientMessage{
		Type: TypeSubscribe,
		Data: json.RawMessage(fmt.Sprintf(`{
			"pendingId": "%s",
			"stateSubscription": {
				"contractAddress": "%s",
				"statePath": ["users"]
			}
		}`, pendingID1, contractAddr.Hex())),
	}
	require.NoError(t, env.ws.WriteJSON(subMsg))

	// Wait for completion or timeout
	select {
	case <-env.done:
	case err := <-env.errChan:
		t.Fatal(err)
	case <-time.After(120 * time.Second):
		t.Fatal("test timeout")
	}
}