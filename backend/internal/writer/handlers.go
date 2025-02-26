package writer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"

	"github.com/msquared-io/etherbase/backend/internal/client"
	"github.com/msquared-io/etherbase/backend/internal/integration/go/etherbasesource"
	"github.com/msquared-io/etherbase/backend/internal/schema"
	"github.com/msquared-io/etherbase/backend/internal/txpool"
	"github.com/msquared-io/etherbase/backend/internal/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/websocket"
)

var abiArgsCache sync.Map

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

func handleSetValue(conn *websocket.Conn, privateKey string, data json.RawMessage) {
	var msg types.SetValueMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		sendError(conn, "Invalid set_value message format")
		return
	}
	log.Printf("handleSetValue: %v", msg)
 
	sourceContract, err := etherbasesource.NewEtherbaseSource(msg.ContractAddress, client.GetManager().Client())
	if err != nil {
		sendError(conn, "failed to create etherbase source: "+err.Error())
		return
	}

	// Handle source contract
	allSegments, err := sourceContract.GetAllSegments(&bind.CallOpts{})
	if err != nil {
		sendError(conn, "failed to get segments: "+err.Error())
		return
	}

	entries, err := sourceContract.GetEntries(&bind.CallOpts{})
	if err != nil {
		sendError(conn, "failed to get entries: "+err.Error())
		return
	}	
	
	// Process state into batch set values
	batchSetValues := []types.BatchSetValue{}
	if err := processState(msg.State, nil, &batchSetValues); err != nil {
		sendError(conn, "Failed to process state: "+err.Error())
		return
	}

	if err := handleRemovingSubEntries(&batchSetValues, allSegments, entries); err != nil {
		sendError(conn, "Failed to handle removing sub entries: "+err.Error())
		return
	}

	log.Printf("batchSetValues: %v", batchSetValues)

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

func encodeSegmentIndices(indices []int) []byte {
	result := make([]byte, len(indices)*4) // Allow space for varint encoding
	offset := 0

	for _, x := range indices {
		if x == 0 {
			result[offset] = 0
			offset++
			continue
		}

		temp := x
		for temp >= 0x80 {
			result[offset] = byte(temp&0x7f) | 0x80
			temp >>= 7
			offset++
		}
		result[offset] = byte(temp & 0x7f)
		offset++
	}

	// Add terminating zero
	result[offset] = 0
	offset++

	return result[:offset]
}

func decodeSegmentIndices(data []byte) []int {
	var result []int
	i := 0

	for i < len(data) {
		if data[i] == 0 {
			if i == len(data)-1 {
				break
			}
			result = append(result, 0)
			i++
			continue
		}

		value := 0
		shift := 0

		for {
			b := data[i]
			value |= int(b&0x7f) << shift

			if b&0x80 == 0 {
				break
			}

			shift += 7
			i++
		}

		result = append(result, value)
		i++
	}

	return result
}


func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func indexOf(slice []string, item string) int {
	for i, s := range slice {
		if s == item {
			return i
		}
	}
	return -1
}

func handleRemovingSubEntries(batchSetValues *[]types.BatchSetValue, allSegments []string, entries []etherbasesource.EtherDatabaseLibNode) error {
	if len(entries) == 0 {
		return nil
	}

	for _, batchSetValue := range *batchSetValues {
		// for each path, we need to check to see if the path below it is set. we need to get all entries that may be affected by this change
		// e.g. if the existing path is ["users", "alice", ["name, "age"]] and we set ["users"] to NONE, we need to remove all entries under ["users", "alice"]

		// first for each path e.g. ["users"] or ["users", "alice"] or ["users", "alice", "name"]
		// we check if all segments exist. if they do, we need to consider removing the path.
		foundAllSegments := true
		for _, segment := range batchSetValue.Segments {
			if !contains(allSegments, segment) {
				foundAllSegments = false
				break
			}
		}

		// if all the segments are present, we need to consider removing this path or subpaths.
		// if not, it means the path is new and we dont need to consider removing anything.
		if !foundAllSegments {		
			continue
		}
			
		// find the entry
		segmentIndices := make([]int, len(batchSetValue.Segments))
		for i, segment := range batchSetValue.Segments {
			idx := indexOf(allSegments, segment)
			if idx == -1 {
				return fmt.Errorf("segment %s not found", segment)
			}
			segmentIndices[i] = idx + 1
		}

		encodedPath := encodeSegmentIndices(segmentIndices)

		// if this is an exact match, we need to remove all subpaths so skip
		skip := false
		for _, entry := range entries {
			if entry.Exists && bytes.Equal(entry.Path, encodedPath) {
				skip = true
				break
			}
		}

		if skip {
			continue
		}

		// remove last byte from encodedPath as its null-terminated but we're doing prefix match
		encodedPath = encodedPath[:len(encodedPath)-1]

		// if this is a partial match, it means we're modifying a subpath of an existing path and so the existing path needs to be removed
		for _, entry := range entries {
			if !entry.Exists || !bytes.HasPrefix(entry.Path, encodedPath) {
				continue
			}

			segments := decodeSegmentIndices(entry.Path)
			segmentsString := make([]string, len(segments))
			for i, segment := range segments {
				segmentsString[i] = allSegments[segment-1]
			}

			// remove the entry
			*batchSetValues = append(*batchSetValues, types.BatchSetValue{
				Segments: segmentsString,
				DataType: uint8(types.DataTypeNone),
				Data:     []byte{},
			})
		}
	}

	return nil
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

// Helper function to convert string to big.Int
func stringToBigInt(s string, allowNegative bool) (*big.Int, error) {
	bigInt := new(big.Int)
	_, success := bigInt.SetString(s, 0) // 0 means auto-detect base
	if !success {
		return nil, fmt.Errorf("invalid number format: %s", s)
	}
	
	// Check if negative numbers are allowed
	if !allowNegative && bigInt.Sign() < 0 {
		return nil, fmt.Errorf("negative value not allowed for unsigned integer: %s", s)
	}
	
	return bigInt, nil
}

// Helper function to convert string to bool
func stringToBool(s string) (bool, error) {
	s = strings.ToLower(s)
	switch s {
	case "true", "1", "yes", "y":
		return true, nil
	case "false", "0", "no", "n":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean value: %s", s)
	}
}

// Helper function to convert any value to the appropriate type based on ABI type
func convertValueByType(value interface{}, typeName string, paramName string) (interface{}, error) {
	// If value is nil, return an error
	if value == nil {
		return nil, fmt.Errorf("nil value provided for parameter %s", paramName)
	}

	// For numeric types, always convert to *big.Int
	if strings.HasPrefix(typeName, "uint") || strings.HasPrefix(typeName, "int") {
		// If already a *big.Int, return as is
		if bigInt, ok := value.(*big.Int); ok {
			return bigInt, nil
		}
		
		// Handle string values
		if strValue, ok := value.(string); ok {
			allowNegative := strings.HasPrefix(typeName, "int")
			return stringToBigInt(strValue, allowNegative)
		}
		
		// Handle float64 values
		if floatValue, ok := value.(float64); ok {
			return big.NewInt(int64(floatValue)), nil
		}
		
		// Handle int values
		if intValue, ok := value.(int); ok {
			return big.NewInt(int64(intValue)), nil
		}
		
		// Handle int64 values
		if int64Value, ok := value.(int64); ok {
			return big.NewInt(int64Value), nil
		}
		
		return nil, fmt.Errorf("cannot convert value of type %T to %s for parameter %s", value, typeName, paramName)
	}
	
	// Handle other types
	switch typeName {
	case "string":
		if strValue, ok := value.(string); ok {
			return strValue, nil
		}
		// Convert non-string to string if possible
		return fmt.Sprintf("%v", value), nil
		
	case "bool":
		// If already a bool, return as is
		if boolVal, ok := value.(bool); ok {
			return boolVal, nil
		}
		// If string, try to convert to bool
		if strValue, ok := value.(string); ok {
			return stringToBool(strValue)
		}
		// If number, 0 is false, anything else is true
		if floatValue, ok := value.(float64); ok {
			return floatValue != 0, nil
		}
		
	case "address":
		// If already an address, return as is
		if addr, ok := value.(common.Address); ok {
			return addr, nil
		}
		// If string, convert to address
		if strValue, ok := value.(string); ok {
			return common.HexToAddress(strValue), nil
		}
	}
	
	// For any other type, return as is and let the ABI packer handle it
	return value, nil
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

	// Convert values to appropriate types based on ABI type
	convertedValues := make([]interface{}, len(values))
	for i, value := range values {
		if i < len(args) {
			converted, err := convertValueByType(value, args[i], "")
			if err != nil {
				return nil, err
			}
			convertedValues[i] = converted
		} else {
			convertedValues[i] = value
		}
	}

	return realArgs.Pack(convertedValues...)
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
			
			// Convert value to appropriate type
			convertedValue, err := convertValueByType(value, arg.ArgType, arg.Name)
			if err != nil {
				return nil, nil, err
			}
			
			switch arg.ArgType {
			case "string":
				strVal := convertedValue.(string)
				topicData = crypto.Keccak256([]byte(strVal))
			case "address":
				addr := convertedValue.(common.Address)
				topicData = common.LeftPadBytes(addr.Bytes(), 32)
			case "uint256", "int256":
				bigInt := convertedValue.(*big.Int)
				topicData = common.LeftPadBytes(bigInt.Bytes(), 32)
			case "bool":
				boolVal := convertedValue.(bool)
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
					topicData = crypto.Keccak256(bytesVal)
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

	// Encode non-indexed parameters
	data, err := encodeAbiParameters(dataTypes, dataValues)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to encode parameters: %v", err)
	}

	return topics, data, nil
}

func handleExecuteContractMethod(conn *websocket.Conn, privateKey string, data json.RawMessage) {
	var msg types.ExecuteContractMethodMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		sendError(conn, "Invalid execute_contract_method message format")
		return
	}

	abi := schema.GetSchemaProvider().GetContractABI(msg.ContractAddress)
	if abi == nil {
		sendError(conn, "Failed to get contract ABI as it has not been added. Please add the contract ABI to the etherbase contract")
		return
	}

	method, ok := abi.Methods[msg.MethodName]
	if !ok {
		sendError(conn, "Method not found: "+msg.MethodName)
		return
	}
	
	// execute method
	txReq := &txpool.TxRequest{
		ContractAddress: msg.ContractAddress,
		PrivateKey:     privateKey,
		Payload: txpool.TxPayload{
			FunctionName: msg.MethodName,
			GetArgs: func() ([]interface{}, error) {
				// Convert map to ordered slice of args based on method inputs
				inputArgs := make([]interface{}, len(method.Inputs))
				for i, input := range method.Inputs {
					argValue, exists := msg.Args[input.Name]
					if !exists {
						return nil, fmt.Errorf("missing parameter %s", input.Name)
					}

					log.Printf("argValue: %v, input: %v", argValue, input)
					
					// Convert value to appropriate type based on input type
					converted, err := convertValueByType(argValue, input.Type.String(), input.Name)
					if err != nil {
						return nil, err
					}
					inputArgs[i] = converted
				}
				
				log.Printf("inputArgs: %v", inputArgs)
				
				// Pack arguments directly using the method
				packedData, err := method.Inputs.Pack(inputArgs...)
				if err != nil {
					log.Printf("Error packing arguments: %v", err)
					return nil, fmt.Errorf("failed to pack arguments: %v", err)
				}
				
				log.Printf("packed data: %v", packedData)

				// so it spreads correctly
				finalArgs := []interface{}{}
				for _, arg := range packedData {
					finalArgs = append(finalArgs, arg)
				}

				return finalArgs, nil
			},
		},
		Abi: abi,
	}

	txpool.GetTransactionPool().AddTransaction(*txReq)
}