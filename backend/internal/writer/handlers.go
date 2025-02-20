package writer

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/msquared-io/etherbase/backend/internal/schema"
	"github.com/msquared-io/etherbase/backend/internal/txpool"
	"github.com/msquared-io/etherbase/backend/internal/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/websocket"
)

var abiArgsCache sync.Map

func handleSetValue(conn *websocket.Conn, privateKey string, data json.RawMessage) {
	var msg types.SetValueMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		sendError(conn, "Invalid set_value message format")
		return
	}

	// Process state into batch set values
	batchSetValues := []types.BatchSetValue{}
	if err := processState(msg.State, nil, &batchSetValues); err != nil {
		sendError(conn, "Failed to process state: "+err.Error())
		return
	}

	txReq := &txpool.TxRequest{
		ContractAddress: msg.ContractAddress,
		PrivateKey:     privateKey,
		Payload: txpool.TxPayload{
			FunctionName: "setValues",
			GetArgs: func() ([]interface{}, error) {
				return []interface{}{batchSetValues}, nil
			},
		},
	}

	// Add to transaction pool
	txpool.GetTransactionPool().AddTransaction(*txReq)
}

func handleEmitEvent(conn *websocket.Conn, privateKey string, data json.RawMessage) {
	var msg types.EmitEventMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		sendError(conn, "Invalid emit_event message format")
		return
	}

	// Create transaction request
	txReq := &txpool.TxRequest{
		ContractAddress: msg.ContractAddress,
		PrivateKey:     privateKey,
		Payload: txpool.TxPayload{
			FunctionName: "emitEvent",
			GetArgs: func() ([]interface{}, error) {
				parameterTopics, parameterData, err := encodeEvent(msg.ContractAddress, msg.Name, msg.Args)
				if err != nil {
					return nil, err
				}
				return []interface{}{msg.Name, parameterTopics, parameterData}, nil
			},
		},
	}

	// Add to transaction pool
	txpool.GetTransactionPool().AddTransaction(*txReq)
}

// Helper function to process state recursively
func processState(obj map[string]interface{}, path []string, result *[]types.BatchSetValue) error {
	for key, value := range obj {
		currentPath := make([]string, len(path)+1)
		copy(currentPath, path)
		currentPath[len(path)] = key

		if value == nil {
			*result = append(*result, types.BatchSetValue{
				Segments: currentPath,
				DataType: uint8(types.DataTypeNone),
				Data:     []byte{},
			})
			continue
		}

		switch v := value.(type) {
		case map[string]interface{}:
			if err := processState(v, currentPath, result); err != nil {
				return err
			}
		case string:
			data, err := encodeAbiParameters([]string{"string"}, []interface{}{v})
			if err != nil {
				return err
			}
			*result = append(*result, types.BatchSetValue{
				Segments: currentPath,
				DataType: uint8(types.DataTypeString),
				Data:     data,
			})
		case bool:
			data, err := encodeAbiParameters([]string{"bool"}, []interface{}{v})
			if err != nil {
				return err
			}
			*result = append(*result, types.BatchSetValue{
				Segments: currentPath,
				DataType: uint8(types.DataTypeBool),
				Data:     data,
			})
		case float64:
			if float64(int64(v)) == v {
				// It's an integer
				dataType := uint8(types.DataTypeUint256)
				if v < 0 {
					dataType = uint8(types.DataTypeInt256)
				}
				data, err := encodeAbiParameters([]string{"uint256"}, []interface{}{big.NewInt(int64(v))})
				if err != nil {
					return err
				}
				*result = append(*result, types.BatchSetValue{
					Segments: currentPath,
					DataType: dataType,
					Data:     data,
				})
			} else {
				return fmt.Errorf("only integer numbers are supported")
			}
		default:
			return fmt.Errorf("unsupported value type: %T", value)
		}
	}
	return nil
}


type AbiArgument struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Indexed bool `json:"indexed"`
}

func encodeAbiParameters(args []string, values []interface{}) ([]byte, error) {
	// Create cache key by joining args with a delimiter
	cacheKey := strings.Join(args, "|")
	
	var realArgs abi.Arguments
	if cachedArgs, found := abiArgsCache.Load(cacheKey); found {
		realArgs = cachedArgs.(abi.Arguments)
	} else {
		// Create new AbiArgument list if not in cache
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

		err = json.Unmarshal(argListJson, &realArgs)
		if err != nil {
			return nil, err
		}
		
		// Store in cache
		abiArgsCache.Store(cacheKey, realArgs)
	}

	return realArgs.Pack(values...)
}

// Helper function to encode event parameters
func encodeEvent(contractAddr common.Address, eventName string, args map[string]interface{}) ([][32]byte, []byte, error) {
	eventSchema := schema.GetSchemaProvider().GetEventSchemaByName(contractAddr, eventName)
	if eventSchema == nil {
		return nil, nil, fmt.Errorf("event schema not found for %s", eventName)
	}

	// Create parameter values array in same order as schema args
	paramValues := make([]interface{}, len(eventSchema.Args))
	for i, arg := range eventSchema.Args {
		val, ok := args[arg.Name]
		if !ok {
			return nil, nil, fmt.Errorf("missing parameter %s", arg.Name)
		}
		paramValues[i] = val
	}

	topics := make([][32]byte, 0)
	dataTypes := make([]string, 0)
	dataValues := make([]interface{}, 0)

	// Process each parameter
	for i, arg := range eventSchema.Args {
		value := paramValues[i]

		if arg.IsIndexed {
			// Handle indexed parameters
			var topicData []byte
			switch arg.ArgType {
			case "string":
				strVal, ok := value.(string)
				if !ok {
					return nil, nil, fmt.Errorf("invalid string value for %s", arg.Name)
				}
				hash := crypto.Keccak256([]byte(strVal))
				topicData = hash
			case "address":
				addr, ok := value.(common.Address)
				if !ok {
					if strAddr, ok := value.(string); ok {
						addr = common.HexToAddress(strAddr)
					} else {
						return nil, nil, fmt.Errorf("invalid address value for %s", arg.Name)
					}
				}
				topicData = common.LeftPadBytes(addr.Bytes(), 32)
			case "uint256", "int256":
				numVal, ok := value.(float64)
				if !ok {
					return nil, nil, fmt.Errorf("invalid number value for %s", arg.Name)
				}
				bigInt := big.NewInt(int64(numVal))
				topicData = common.LeftPadBytes(bigInt.Bytes(), 32)
			case "bool":
				boolVal, ok := value.(bool)
				if !ok {
					return nil, nil, fmt.Errorf("invalid bool value for %s", arg.Name)
				}
				if boolVal {
					topicData = common.LeftPadBytes([]byte{1}, 32)
				} else {
					topicData = common.LeftPadBytes([]byte{0}, 32)
				}
			default:
				if strings.HasPrefix(arg.ArgType, "bytes") {
					bytesVal, ok := value.([]byte)
					if !ok {
						return nil, nil, fmt.Errorf("invalid bytes value for %s", arg.Name)
					}
					hash := crypto.Keccak256(bytesVal)
					topicData = hash
				} else {
					return nil, nil, fmt.Errorf("unsupported indexed parameter type: %s", arg.ArgType)
				}
			}
			topics = append(topics, [32]byte(topicData))
		} else {
			// Handle non-indexed parameters
			dataTypes = append(dataTypes, arg.ArgType)
			dataValues = append(dataValues, value)
		}
	}

	// if a data value is float64, convert to big.Int
	for i, value := range dataValues {
		if floatValue, ok := value.(float64); ok {
			dataValues[i] = big.NewInt(int64(floatValue))
		}
	}

	// Encode non-indexed parameters
	data, err := encodeAbiParameters(dataTypes, dataValues)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to encode parameters: %v", err)
	}

	return topics, data, nil
}