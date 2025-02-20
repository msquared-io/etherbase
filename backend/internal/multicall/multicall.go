package multicall

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/msquared-io/etherbase/backend/internal/client"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// MULTICALL_3_ABI represents the ABI for the Multicall3 contract
var MULTICALL_3_ABI = `[{
    "inputs": [
        {"internalType": "bool","name": "requireSuccess","type": "bool"},
        {
            "components": [
                {"internalType": "address","name": "target","type": "address"},
                {"internalType": "bytes","name": "callData","type": "bytes"}
            ],
            "internalType": "struct Multicall3.Call[]",
            "name": "calls",
            "type": "tuple[]"
        }
    ],
    "name": "tryBlockAndAggregate",
    "outputs": [
        {"internalType": "uint256","name": "blockNumber","type": "uint256"},
        {"internalType": "bytes32","name": "blockHash","type": "bytes32"},
        {
            "components": [
                {"internalType": "bool","name": "success","type": "bool"},
                {"internalType": "bytes","name": "returnData","type": "bytes"}
            ],
            "internalType": "struct Multicall3.Result[]",
            "name": "returnData",
            "type": "tuple[]"
        }
    ],
    "stateMutability": "payable",
    "type": "function"
}]`

// Example: bridging your existing TS multicall logic.
// In practice, you'd need a compiled "multicall" contract ABI/binding
// and do aggregated calls.

func ExecuteMulticall(ctx context.Context, queries []MulticallQuery, requireSuccess bool) (*MulticallResult, error) {
	ethClient := client.GetManager().Client()
	
	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(MULTICALL_3_ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse multicall ABI: %w", err)
	}

	// Prepare the call arguments
	type Call struct {
		Target   common.Address
		CallData []byte
	}
	
	calls := make([]Call, len(queries))
	for i, query := range queries {
		calls[i] = Call{
			Target:   query.Target,
			CallData: query.CallData,
		}
	}

	// Pack the function call
	data, err := parsedABI.Pack("tryBlockAndAggregate", requireSuccess, calls)
	if err != nil {
		return nil, fmt.Errorf("failed to pack function call: %w", err)
	}

	// Create the call message
	msg := ethereum.CallMsg{
		To:   &client.GetManager().MulticallAddress,
		Data: data,
	}

	// Execute the call
	output, err := ethClient.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, fmt.Errorf("multicall failed: %w", err)
	}

	// Unpack the result
	result, err := parsedABI.Unpack("tryBlockAndAggregate", output)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %w", err)
	}

	// Parse the results
	blockNumber := result[0].(*big.Int)
	blockHash := result[1].([32]byte)
	returnData := result[2].([]struct {
		Success    bool    `json:"success"`
		ReturnData []uint8 `json:"returnData"`
	})

	// Convert to our result format
	callResults := make([]CallResult, len(returnData))
	for i, data := range returnData {
		callResults[i] = CallResult{
			Success:    data.Success,
			ReturnData: data.ReturnData,
		}
	}

	return &MulticallResult{
		BlockNumber: blockNumber,
		BlockHash:   blockHash,    // Add BlockHash to the result
		Results:     callResults,
	}, nil
}

type MulticallQuery struct {
	Target   common.Address
	CallData []byte
}

type MulticallResult struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	Results     []CallResult
}

type CallResult struct {
	Success    bool
	ReturnData []byte
}

func ReadContractExample(addr common.Address, data []byte) ([]byte, error) {
	// do a "ethclient.CallContract()" call
	return []byte{}, nil
}