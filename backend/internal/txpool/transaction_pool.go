package txpool

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/msquared-io/etherbase/backend/internal/config"
	"github.com/msquared-io/etherbase/backend/internal/integration/go/etherbasesource"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type TxRequest struct {
    ContractAddress common.Address
    PrivateKey      string // hex
    Payload         TxPayload
}

type TxPayload struct {
    FunctionName string
    GetArgs      func() ([]interface{}, error)
}

// Add new type for tracking transaction results
type TxResult struct {
    Hash     common.Hash
    Request  TxRequest
    SentTime time.Time
}

// Add new type for prepared transaction
type PreparedTx struct {
    txHex     string
    txHash    common.Hash
    txRequest TxRequest
}

// TransactionPool is a naive queue that processes transaction requests in batches
type TransactionPool struct {
    pool           []TxRequest
    processingLock sync.Mutex
    addingLock     sync.Mutex
    sendingLock    sync.Mutex
    stopCh         chan struct{}
    sourceABI     abi.ABI
    keyCache       sync.Map    // Changed from map to sync.Map
}

var txPoolInstance *TransactionPool
var txPoolOnce sync.Once
var client *RPCClient

func MakeTransactionPool(config *config.Config) {
    txPoolOnce.Do(func() {
		var err error
		client, err = NewRPCClient(config.RpcURL)
		if err != nil {
			log.Fatalf("Failed to connect to Ethereum client: %v", err)
		}

        // Parse the ABI once during initialization
        sourceABI, err := abi.JSON(strings.NewReader(etherbasesource.EtherbaseSourceABI))
        if err != nil {
            log.Fatalf("failed to parse source ABI: %v", err)
        }

        txPoolInstance = &TransactionPool{
            pool:       make([]TxRequest, 0, 100),
            stopCh:     make(chan struct{}),
            sourceABI: sourceABI,
        }
        go txPoolInstance.loop()
    })
}

func GetTransactionPool() *TransactionPool {
    return txPoolInstance
}

func (tp *TransactionPool) AddTransaction(req TxRequest) {
    tp.addingLock.Lock()
    defer tp.addingLock.Unlock()
    tp.pool = append(tp.pool, req)
}

func (tp *TransactionPool) loop() {
    ticker := time.NewTicker(50 * time.Millisecond)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            tp.checkAndProcess()
        case <-tp.stopCh:
            return
        }
    }
}

func (tp *TransactionPool) checkAndProcess() {
    tp.processingLock.Lock()
    defer tp.processingLock.Unlock()

    tp.addingLock.Lock()
    if len(tp.pool) == 0 {
        tp.addingLock.Unlock()
        return
    }

    // Example batch logic
    batchSize := 100
    if len(tp.pool) < batchSize {
        batchSize = len(tp.pool)
    }

    batch := tp.pool[:batchSize]
    tp.pool = tp.pool[batchSize:]

    tp.addingLock.Unlock()
    tp.processBatch(batch)
}

// getOrCreatePrivateKey safely retrieves or creates an ECDSA private key
func (tp *TransactionPool) getOrCreatePrivateKey(hexKey string) (*ecdsa.PrivateKey, error) {
    // Check if key exists
    if privKey, exists := tp.keyCache.Load(hexKey); exists {
        return privKey.(*ecdsa.PrivateKey), nil
    }

    // Create new key
    privKey, err := crypto.HexToECDSA(hexKey[2:])
    if err != nil {
        return nil, fmt.Errorf("failed to parse private key: %v", err)
    }

    // Store using LoadOrStore to handle race conditions
    actual, _ := tp.keyCache.LoadOrStore(hexKey, privKey)
    return actual.(*ecdsa.PrivateKey), nil
}

// debug mode slows all transactions but checks all errors and receipts
const debug = true

type TimingStats struct {
    GetArgsTime        time.Duration
    GetPrivKeyTime     time.Duration
    PackDataTime       time.Duration
    SigningTime        time.Duration
    MarshalTime        time.Duration
    SendTransactionTime time.Duration
}

func (tp *TransactionPool) processBatch(batch []TxRequest) {
    startTime := time.Now()
    chainID := big.NewInt(50311)
    signer := types.LatestSignerForChainID(chainID)
    
    stats := TimingStats{}
    preparedTxs := make([]PreparedTx, 0, len(batch))

    fmt.Printf("Processing batch of %d transactions\n", len(batch))

    // First prepare all transactions without locking
    for i, txReq := range batch {
        // Time GetArgs
        argsStart := time.Now()
        args, err := txReq.Payload.GetArgs()
        if err != nil {
            fmt.Printf("Error building args: %v\n", err)
            continue
        }
        stats.GetArgsTime += time.Since(argsStart)

        // Time getting private key
        keyStart := time.Now()
        privKey, err := tp.getOrCreatePrivateKey(txReq.PrivateKey)
        if err != nil {
            log.Printf("Failed to get private key for transaction %d: %v", i, err)
            continue
        }
        stats.GetPrivKeyTime += time.Since(keyStart)

        // Time packing data
        packStart := time.Now()
        data, err := tp.sourceABI.Pack(txReq.Payload.FunctionName, args...)
        if err != nil {
            log.Printf("Failed to pack data for transaction %d: %v", i, err)
            continue
        }
        stats.PackDataTime += time.Since(packStart)

        nonceVal := uint64(time.Now().UnixNano() / 1000)
        tx := types.NewTx(&types.DynamicFeeTx{
            ChainID:   chainID,
            Nonce:     nonceVal,
            GasTipCap: big.NewInt(0),
            GasFeeCap: big.NewInt(36000000000),
            Gas:       uint64(2000000),
            To:        &txReq.ContractAddress,
            Value:     big.NewInt(0),
            Data:      data,
        })

        // Time signing
        signStart := time.Now()
        signedTx, err := types.SignTx(tx, signer, privKey)
        if err != nil {
            log.Fatalf("Failed to sign transaction %d: %v", i, err)
        }
        stats.SigningTime += time.Since(signStart)

        // Time marshaling
        marshalStart := time.Now()
        rawTx, err := signedTx.MarshalBinary()
        if err != nil {
            log.Fatalf("Failed to marshal transaction %d: %v", i, err)
        }
        txHex := "0x" + hex.EncodeToString(rawTx)
        stats.MarshalTime += time.Since(marshalStart)

        preparedTxs = append(preparedTxs, PreparedTx{
            txHex:     txHex,
            txHash:    signedTx.Hash(),
            txRequest: txReq,
        })
    }

    txResults := make([]TxResult, 0, len(preparedTxs))
    
    // Send all prepared transactions while holding the lock
    for _, prepared := range preparedTxs {
        sendStart := time.Now()

        if debug {
            // In debug mode, use Call to get response and verify hash
            result, err := client.Call("eth_sendRawTransaction", []interface{}{prepared.txHex})
            if err != nil {
                log.Printf("Failed to send transaction with hash %s: %v", prepared.txHash.Hex(), err)
                continue
            }
            var respHash common.Hash
            if err := json.Unmarshal(result, &respHash); err != nil {
                log.Printf("Failed to parse transaction hash for %s: %v", prepared.txHash.Hex(), err)
                continue
            }
            // Verify the hash matches what we calculated
            if respHash != prepared.txHash {
                log.Printf("Warning: Calculated hash %s doesn't match response hash %s", prepared.txHash.Hex(), respHash.Hex())
            }
        } else {
            // In non-debug mode, just send without waiting for response
            if err := client.Send("eth_sendRawTransaction", []interface{}{prepared.txHex}); err != nil {
                log.Printf("Failed to send transaction with hash %s: %v", prepared.txHash.Hex(), err)
                continue
            }
        }
        stats.SendTransactionTime += time.Since(sendStart)

        // Track transaction result
        txResults = append(txResults, TxResult{
            Hash:     prepared.txHash,
            Request:  prepared.txRequest,
            SentTime: time.Now(),
        })
    }

    duration := time.Since(startTime)
    fmt.Printf("Sent %d transactions in %v\n", len(preparedTxs), duration)


    // Start goroutine to monitor transaction receipts only in debug mode
    if debug {
        go tp.monitorTransactionReceipts(txResults)
    }
}

func (tp *TransactionPool) monitorTransactionReceipts(txResults []TxResult) {
    // Wait a bit before starting to check receipts
    time.Sleep(2 * time.Second)

    maxAttempts := 5
    for attempt := 0; attempt < maxAttempts; attempt++ {
        pendingTxs := make([]TxResult, 0)

        for _, txResult := range txResults {
            var receipt map[string]interface{}
            result, err := client.Call("eth_getTransactionReceipt", []interface{}{txResult.Hash.Hex()})
            if err != nil {
                log.Printf("Error checking receipt for tx %s: %v", txResult.Hash.Hex(), err)
                pendingTxs = append(pendingTxs, txResult)
                continue
            }

            if err := json.Unmarshal(result, &receipt); err != nil {
                log.Printf("Failed to parse receipt for tx %s: %v", txResult.Hash.Hex(), err)
                continue
            }

            if receipt == nil {
                // Transaction still pending
                pendingTxs = append(pendingTxs, txResult)
                continue
            }

            // Check transaction status
            status, ok := receipt["status"].(string)
            if !ok {
                log.Printf("Invalid status format for tx %s", txResult.Hash.Hex())
                continue
            }

            // status "0x1" means success, "0x0" means failure
            if status == "0x0" {
                log.Printf("Transaction failed - Hash: %s, Function: %s", 
                    txResult.Hash.Hex(), txResult.Request.Payload.FunctionName)
            }
        }

        // If no pending transactions left, we're done
        if len(pendingTxs) == 0 {
            break
        }

        // Update txResults for next iteration
        txResults = pendingTxs

        // Wait before next check
        time.Sleep(3 * time.Second)
    }
}

func (tp *TransactionPool) Stop() {
    close(tp.stopCh)
}
