package state

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"math/big"

	"github.com/msquared-io/etherbase/backend/internal/client"
	"github.com/msquared-io/etherbase/backend/internal/integration/go/etherbasesource"
	"github.com/msquared-io/etherbase/backend/internal/multicall"
	"github.com/msquared-io/etherbase/backend/internal/schema"
	"github.com/msquared-io/etherbase/backend/internal/sources"
	"github.com/msquared-io/etherbase/backend/internal/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// StateFetchRequest represents a request to fetch state
type StateFetchRequest struct {
	ContractAddress common.Address
	Path           types.StatePath
}

// StateFetchResponse contains the fetched state
type StateFetchResponse struct {
	ContractAddress common.Address
	State          map[string]interface{}
}

type BatchFetchResult struct {
	States      []StateFetchResponse
	BlockNumber uint64
}

type requestTracking struct {
	ContractAddress common.Address
	RequestIndex   int
	QueryIndex     int
	SegmentsString []string
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

// convertArgToType converts a string argument to the appropriate type based on the ABI type
func convertArgToType(arg string, abiType abi.Type) (interface{}, error) {
	switch abiType.T {
	case abi.IntTy, abi.UintTy:
		// Handle both int and uint types
		bigInt := new(big.Int)
		// Check if hex
		if strings.HasPrefix(arg, "0x") {
			_, ok := bigInt.SetString(arg[2:], 16)
			if !ok {
				return nil, fmt.Errorf("failed to parse hex number: %s", arg)
			}
		} else {
			_, ok := bigInt.SetString(arg, 10)
			if !ok {
				return nil, fmt.Errorf("failed to parse decimal number: %s", arg)
			}
		}
		return bigInt, nil
	case abi.AddressTy:
		if !strings.HasPrefix(arg, "0x") {
			return nil, fmt.Errorf("address must start with 0x: %s", arg)
		}
		return common.HexToAddress(arg), nil
	case abi.BytesTy, abi.FixedBytesTy:
		if !strings.HasPrefix(arg, "0x") {
			return nil, fmt.Errorf("bytes must start with 0x: %s", arg)
		}
		return common.FromHex(arg), nil
	case abi.BoolTy:
		return strings.ToLower(arg) == "true", nil
	case abi.StringTy:
		return arg, nil
	default:
		return nil, fmt.Errorf("unsupported type: %v", abiType)
	}
}

// BatchFetchContractStates fetches state from multiple contracts
func BatchFetchContractStates(ctx context.Context, requests []StateFetchRequest) (*BatchFetchResult, error) {
	// Group requests by contract address
	grouped := make(map[common.Address][]StateFetchRequest)
	for _, r := range requests {
		grouped[r.ContractAddress] = append(grouped[r.ContractAddress], r)
	}

	var multicallQueries []multicall.MulticallQuery
	var requestTracker []requestTracking

	// Process each contract's requests
	for addr, contractRequests := range grouped {
		abi := schema.GetSchemaProvider().GetContractABI(addr)
		if abi == nil {
			return nil, fmt.Errorf("no ABI found for contract %s", addr.Hex())
		}

		isSourceContract := sources.GetSourceRegistry().IsSource(addr)

		for reqIdx, request := range contractRequests {
			if isSourceContract {
                sourceContract, err := etherbasesource.NewEtherbaseSource(addr, client.GetManager().Client())
				if err != nil {
					return nil, fmt.Errorf("failed to create etherbase source: %w", err)
				}

				// Handle source contract
				allSegments, err := sourceContract.GetAllSegments(&bind.CallOpts{
					Context: ctx,
				})
				if err != nil {
					return nil, fmt.Errorf("failed to get segments: %w", err)
				}

				// Check if all segments exist
				foundAllSegments := true
				for _, segment := range request.Path {
					if !contains(allSegments, segment) {
						foundAllSegments = false
						break
					}
				}

				if !foundAllSegments {
					// log.Printf("not all segments found for %s", addr.Hex())
					continue
				}

				// Convert path to segment indices
				segmentIndices := make([]int, len(request.Path))
				for i, segment := range request.Path {
					idx := indexOf(allSegments, segment)
					if idx == -1 {
						return nil, fmt.Errorf("segment %s not found", segment)
					}
					segmentIndices[i] = idx + 1
				}

				encodedPath := encodeSegmentIndices(segmentIndices)
				// remove last byte from encodedPath as its null-terminated but we're doing prefix match
				encodedPath = encodedPath[:len(encodedPath)-1]
				entries, err := sourceContract.GetEntries(&bind.CallOpts{
					Context: ctx,
				})
				if err != nil {
					return nil, fmt.Errorf("failed to get entries: %w", err)
				}

				// Find exact match
				exactMatch := false
				for _, entry := range entries {
					if entry.Exists && bytes.Equal(entry.Path, encodedPath) {
						exactMatch = true
						callData, err := abi.Pack("getValue", request.Path)
						if err != nil {
							return nil, fmt.Errorf("failed to pack getValue call: %w", err)
						}

						multicallQueries = append(multicallQueries, multicall.MulticallQuery{
							Target:   addr,
							CallData: callData,
						})

						requestTracker = append(requestTracker, requestTracking{
							ContractAddress: addr,
							RequestIndex:   reqIdx,
							QueryIndex:     len(multicallQueries) - 1,
							SegmentsString: request.Path,
						})
						break
					}
				}

				if exactMatch {
					continue
				}

				// Find partial matches
				for _, entry := range entries {
					if !entry.Exists {
						continue
					}

					if bytes.HasPrefix(entry.Path, encodedPath) {
						// Decode entry path back to segments
						segments := decodeSegmentIndices(entry.Path)
						segmentsString := make([]string, len(segments))
						for i, segment := range segments {
							segmentsString[i] = allSegments[segment-1]
						}

						callData, err := abi.Pack("getValue", segmentsString)
						if err != nil {
							return nil, fmt.Errorf("failed to pack getValue call for partial match: %w", err)
						}

						multicallQueries = append(multicallQueries, multicall.MulticallQuery{
							Target:   addr,
							CallData: callData,
						})

						requestTracker = append(requestTracker, requestTracking{
							ContractAddress: addr,
							RequestIndex:   reqIdx,
							QueryIndex:     len(multicallQueries) - 1,
							SegmentsString: segmentsString,
						})
					}
				}
			} else {
				// Handle regular contract
				segments := request.Path
				if len(segments) == 0 {
					continue
				}

				variableName := segments[0]
				args := segments[1:]

				// Get the function from ABI to check input types
				method, exists := abi.Methods[variableName]
				if !exists {
					return nil, fmt.Errorf("method %s not found in ABI", variableName)
				}

				if len(args) != len(method.Inputs) {
					return nil, fmt.Errorf("argument count mismatch for %s: expected %d, got %d", 
						variableName, len(method.Inputs), len(args))
				}

				// Convert each argument to its proper type
				interfaceArgs := make([]interface{}, len(args))
				for i, arg := range args {
					convertedArg, err := convertArgToType(arg, method.Inputs[i].Type)
					if err != nil {
						return nil, fmt.Errorf("failed to convert argument %d for %s: %w", i, variableName, err)
					}
					interfaceArgs[i] = convertedArg
				}

				callData, err := abi.Pack(variableName, interfaceArgs...)
				if err != nil {
					return nil, fmt.Errorf("failed to pack call data: %w", err)
				}

				multicallQueries = append(multicallQueries, multicall.MulticallQuery{
					Target:   addr,
					CallData: callData,
				})

				requestTracker = append(requestTracker, requestTracking{
					ContractAddress: addr,
					RequestIndex:   reqIdx,
					QueryIndex:     len(multicallQueries) - 1,
					SegmentsString: segments,
				})
			}
		}
	}

	if len(multicallQueries) == 0 {
		log.Printf("No multicall queries")
		return nil, nil
	}

	// Execute multicall
	results, err := multicall.ExecuteMulticall(ctx, multicallQueries, true)
	if err != nil {
		return nil, fmt.Errorf("multicall failed: %w", err)
	}

	// Process results
	responses := make([]StateFetchResponse, len(requests))
	for _, tracker := range requestTracker {
		result := results.Results[tracker.QueryIndex]
		if !result.Success {
			continue
		}

		request := requests[tracker.RequestIndex]
		isSourceContract := sources.GetSourceRegistry().IsSource(request.ContractAddress)

		abi := schema.GetSchemaProvider().GetContractABI(tracker.ContractAddress)
		if abi == nil {
			return nil, fmt.Errorf("no ABI found for contract %s", tracker.ContractAddress.Hex())
		}

		var decodedState interface{}
		if isSourceContract {
			// Decode source contract response
			decodedResult, err := abi.Unpack("getValue", result.ReturnData)
			if err != nil {
				return nil, fmt.Errorf("failed to unpack getValue result: %w", err)
			}

			if len(decodedResult) < 2 {
				return nil, fmt.Errorf("invalid getValue result length")
			}

			dataType := types.DataType(decodedResult[0].(uint8))
			encodedValue := decodedResult[1].([]byte)

			// Decode value based on data type
			switch dataType {
			case types.DataTypeString:
				decoded, err := decodeAbiParameters([]string{"string"}, encodedValue)
				if err != nil {
					return nil, fmt.Errorf("failed to decode string value: %w", err)
				}
				decodedState = decoded[0]
			case types.DataTypeBool:
				decoded, err := decodeAbiParameters([]string{"bool"}, encodedValue)
				if err != nil {
					return nil, fmt.Errorf("failed to decode bool value: %w", err)
				}
				decodedState = decoded[0]
			case types.DataTypeUint256:
				decoded, err := decodeAbiParameters([]string{"uint256"}, encodedValue)
				if err != nil {
					return nil, fmt.Errorf("failed to decode uint256 value: %w", err)
				}
				decodedState = decoded[0]
			case types.DataTypeInt256:
				decoded, err := decodeAbiParameters([]string{"int256"}, encodedValue)
				if err != nil {
					return nil, fmt.Errorf("failed to decode int256 value: %w", err)
				}
				decodedState = decoded[0]
			case types.DataTypeBytes:
				decoded, err := decodeAbiParameters([]string{"bytes"}, encodedValue)
				if err != nil {
					return nil, fmt.Errorf("failed to decode bytes value: %w", err)
				}
				decodedState = decoded[0]
			default:
				return nil, fmt.Errorf("unsupported data type: %d", dataType)
			}
		} else {
			// Decode regular contract response
			method, err := getViewFunctionForVariable(abi, request.Path[0])
			if err != nil {
				return nil, fmt.Errorf("method not found for return data: %w", err)
			}

			decodedResult, err := method.Outputs.Unpack(result.ReturnData)
			if err != nil {
				return nil, fmt.Errorf("failed to unpack return data: %w", err)
			}

			if len(decodedResult) > 0 {
				decodedState = decodedResult[0]
				// Convert big numbers to regular numbers
				if bigInt, ok := decodedState.(*big.Int); ok {
					decodedState = bigInt.Int64()
				}
			}
		}

		// Create nested state object
		state := createNestedObject(tracker.SegmentsString, decodedState)

		// Get or create response
		existingResponse := responses[tracker.RequestIndex]
		if existingResponse.ContractAddress == (common.Address{}) {
			existingResponse = StateFetchResponse{
				ContractAddress: tracker.ContractAddress,
				State:          make(map[string]interface{}),
			}
		}

		// Merge states
		responses[tracker.RequestIndex] = StateFetchResponse{
			ContractAddress: tracker.ContractAddress,
			State:          deepMerge(existingResponse.State, state),
		}
	}

	return &BatchFetchResult{
		States:      responses,
		BlockNumber: results.BlockNumber.Uint64(),
	}, nil
}

func getViewFunctionForVariable(abi *abi.ABI, variableName string) (*abi.Method, error) {
	method, ok := abi.Methods[variableName]
	if !ok {
		return nil, fmt.Errorf("method not found: %s", variableName)
	}
	return &method, nil
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

// Example per-contract fetch. In reality, you'd do an actual contract call or multicall logic
func fetchSingleContractStates(ctx context.Context, addr common.Address, batch []StateFetchRequest) ([]StateFetchResponse, error) {
    // We can do: read the ABI from schema provider, decode the paths, call the contract, etc.
    var out []StateFetchResponse
    for _, r := range batch {
        // For demonstration, let's just store a dummy:
        // "path" => "value"
        response := StateFetchResponse{
            ContractAddress: r.ContractAddress,
            State: map[string]interface{}{
                joinPath(r.Path): "dummyValue",
            },
        }
        out = append(out, response)
    }
    return out, nil
}

func joinPath(p types.StatePath) string {
    return fmt.Sprintf("%v", p)
}

// createNestedObject creates a nested map from a path and value
func createNestedObject(path []string, value interface{}) map[string]interface{} {
	if len(path) == 0 {
		return nil
	}

	result := make(map[string]interface{})
	current := result
	for i, key := range path[:len(path)-1] {
		current[key] = make(map[string]interface{})
		current = current[key].(map[string]interface{})
		if i == len(path)-2 {
			current[path[len(path)-1]] = value
		}
	}

	if len(path) == 1 {
		result[path[0]] = value
	}

	return result
}
