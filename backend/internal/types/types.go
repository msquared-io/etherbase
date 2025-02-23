package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

// Subscription Types
type ClientID string
type SubscriptionID string
type StatePath []string

// WebSocket Message Types
type MessageType string

const (
	TypeSubscribe     MessageType = "subscribe"
	TypeUnsubscribe   MessageType = "unsubscribe"
	TypeSetValue      MessageType = "set_value"
	TypeEmitEvent     MessageType = "emit_event"
	TypeExecuteContractMethod MessageType = "execute_contract_method"
	TypeUpdates       MessageType = "updates"
	TypeError         MessageType = "error"
	TypeInitialState  MessageType = "initial_state"
)

// WebSocket Messages
type ClientMessage struct {
	Type MessageType     `json:"type"`
	Data json.RawMessage `json:"data"`
}

type EventSubscriptionRequest struct {
	ContractAddresses []common.Address `json:"contractAddresses,omitempty"`
	ContractAddress  *common.Address  `json:"contractAddress,omitempty"`
	EventNames       []string         `json:"eventNames,omitempty"`
	EventIDs         []string         `json:"eventIds,omitempty"`
}

type StateSubscriptionRequest struct {
	ContractAddress common.Address `json:"contractAddress"`
	StatePath      []interface{}      `json:"statePath"` // can be string or []string
	Options        *MessageStateSubscriptionOptions `json:"options,omitempty"`
}

type MessageStateSubscriptionOptions struct {
	Repoll                              *MessageRepollOptions `json:"repoll,omitempty"`
	CombineStateUpdatesFromSameContract *bool        `json:"combineStateUpdatesFromSameContract,omitempty"`
	CombineStateUpdatesFromSameTransaction *bool        `json:"combineStateUpdatesFromSameTransaction,omitempty"`
	FullStateUpdateOnChange               *bool        `json:"fullStateUpdateOnChange,omitempty"`
}

type MessageRepollOptions struct {
	OnAnyContractEvent *bool 		 `json:"onAnyContractEvent,omitempty"`
	ListenEvents       *[]MessageListenEvent `json:"listenEvents,omitempty"`
}

type MessageListenEvent struct {
	Name string `json:"name"`
	Args map[string][]interface{} `json:"args"`
}

type StateSubscriptionOptions struct {
	Repoll                              RepollOptions
	CombineStateUpdatesFromSameContract bool
	CombineStateUpdatesFromSameTransaction bool
	FullStateUpdateOnChange               bool
}

type ListenEvent struct {
	Name string
	IndexedArgs map[string][]interface{}
	NonIndexedArgs map[string][]interface{}
}


type RepollOptions struct {
	OnAnyContractEvent bool
	ListenEvents       []ListenEvent
}


type UnsubscribeMessage struct {
	SubscriptionID string `json:"subscriptionId"`
}

// Subscription Definitions
type MessageEventSubscription struct {
	ContractAddress   *common.Address  `json:"contractAddress,omitempty"`
	ContractAddresses []common.Address `json:"contractAddresses,omitempty"`
	Events            []MessageEvent    `json:"events,omitempty"`
}

type MessageStateSubscription struct {
	ContractAddress common.Address     `json:"contractAddress"`
	StatePath      []interface{}       `json:"statePath"` // can be string or []string
	Options        *MessageStateSubscriptionOptions `json:"options"`
}

type EventSubscription struct {
	ContractAddresses []common.Address
	Events            []Event
}


type StateSubscription struct {
	ContractAddress common.Address
	StatePaths      []StatePath
	Options        StateSubscriptionOptions
}

// Update Types
type Update struct {
	Type           string      `json:"type"`
	SubscriptionID string      `json:"subscriptionId,omitempty"`
	Data           interface{} `json:"data"`
}

type StateUpdate struct {
	ContractAddress   string                 `json:"contractAddress"`
	State            map[string]interface{} `json:"state"`
	TransactionIndex uint                   `json:"transactionIndex,omitempty"`
}

// Event Types
type DecodedEvent struct {
	ContractAddress string                 `json:"contractAddress"`
	Name           string                 `json:"name"`
	Args           map[string]interface{} `json:"args"`
	EventID        string                 `json:"eventId"`
	Block          uint64                 `json:"block"`
	Sequence       int                    `json:"-"`
	TransactionHash common.Hash           `json:"-"`
}

// Streaming Types
type ViewCallback struct {
	EventUpdates  []SubscriberEventUpdate
	StateUpdates  []SubscriberStateUpdate
	BlockNumber   uint64
}

type SubscriberEventUpdate struct {
	Event            DecodedEvent
	Subscriptions    map[ClientID][]SubscriptionID
	TransactionIndex uint
}

type SubscriberStateUpdate struct {
	SubscriptionID    SubscriptionID
	ContractAddress   common.Address
	State            map[string]interface{}
	TransactionIndex uint
}

// Transaction Types
type SetValueMessage struct {
	ContractAddress common.Address            `json:"contractAddress"`
	State          map[string]interface{}     `json:"state"`
}

type EmitEventMessage struct {
	ContractAddress common.Address            `json:"contractAddress"`
	Name           string                     `json:"name"`
	Args           map[string]interface{}     `json:"args"`
}

type ExecuteContractMethodMessage struct {
	ContractAddress common.Address            `json:"contractAddress"`
	MethodName     string                     `json:"methodName"`
	Args           map[string]interface{}     `json:"args"`
}

type MessageEvent struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
	Args map[string][]interface{} `json:"args,omitempty"`
}

type Event struct {
	Name string
	ID   string
	NonIndexedArgs map[string][]interface{}
	IndexedArgs    map[string][]interface{}
}

type BatchSetValue struct {
	Segments []string
	DataType uint8
	Data     []byte
}

// DataType represents the type of data stored in the EthDB
type DataType uint8

const (
	DataTypeNone DataType = iota
	DataTypeString
	DataTypeBool
	DataTypeUint256
	DataTypeInt256
	DataTypeBytes
)

func DataTypeToString(dataType DataType) string {
	return []string{
		"none", "string", "bool", "uint256", "int256", "bytes"}[dataType]
}
