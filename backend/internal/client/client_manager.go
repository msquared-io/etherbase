package client

import (
	"context"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/msquared-io/etherbase/backend/internal/config"
	"golang.org/x/sync/semaphore"
)

type Manager struct {
    client  *ethclient.Client
	MulticallAddress common.Address	
    PollingIntervalMS int

    mu sync.Mutex
}

var instance *Manager
var once sync.Once

// NewManager creates (or returns existing) Manager
func NewManager(cfg *config.Config) (*Manager, error) {
    var err error
    once.Do(func() {
        cli, err := ethclient.Dial(cfg.RpcURL)
        if err != nil {
            log.Printf("Error dialing Ethereum client: %v", err)
            return
        }
        instance = &Manager{
            client:          cli,
            PollingIntervalMS: cfg.PollingIntervalMS,
            MulticallAddress:   cfg.MulticallAddress,
        }
    })
    return instance, err
}

func GetManager() *Manager {
    return instance
}

func (m *Manager) Client() *ethclient.Client {
    return m.client
}

// A simple concurrency throttle for batch calls
var batchSem = semaphore.NewWeighted(10) // up to 10 concurrent calls

func (m *Manager) BatchCall(ctx context.Context, fn func() error) error {
    err := batchSem.Acquire(ctx, 1)
    if err != nil {
        return err
    }
    defer batchSem.Release(1)
    return fn()
}

// Example get block number
func (m *Manager) CurrentBlockNumber(ctx context.Context) (*big.Int, error) {
    num, err := m.client.BlockNumber(ctx)
    return new(big.Int).SetUint64(num), err
}

// Example checking code at an address (for debugging)
func (m *Manager) CodeAt(ctx context.Context, addr common.Address) ([]byte, error) {
    return m.client.CodeAt(ctx, addr, nil)
}

func (m *Manager) Close() {
    if m.client != nil {
        m.client.Close()
    }
    log.Println("Ethereum client connections closed.")
}
