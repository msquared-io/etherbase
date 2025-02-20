package sources

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/msquared-io/etherbase/backend/internal/client"
	"github.com/msquared-io/etherbase/backend/internal/integration/go/etherbase"
	"github.com/msquared-io/etherbase/backend/internal/integration/go/etherbasesource"
)

// SourceRegistry listens to SourceCreated events from the etherbase contract
// and keeps track of additional addresses. It also notifies watchers about new sources.
type SourceRegistry struct {
    etherbaseAddress      common.Address
    additionalContracts []common.Address

    sources     sync.Map // key: common.Address (source), value: common.Address (owner)
    updateSubs   []chan []common.Address
    stopWatch    context.CancelFunc
    watchDone    sync.WaitGroup
    watchOnce    sync.Once
    mu           sync.Mutex // kept only for updateSubs slice operations
}

var sourceRegistryInstance *SourceRegistry
var sourceRegistryOnce sync.Once

func NewSourceRegistry(etherbaseAddress common.Address, additional []common.Address) *SourceRegistry {
    sourceRegistryOnce.Do(func() {
        sourceRegistryInstance = &SourceRegistry{
            etherbaseAddress:      etherbaseAddress,
            additionalContracts: additional,
        }
        sourceRegistryInstance.startWatch()
		sourceRegistryInstance.loadInitialSources()
    })
    return sourceRegistryInstance
}

func GetSourceRegistry() *SourceRegistry {
    return sourceRegistryInstance
}

func (sr *SourceRegistry) loadInitialSources() {
	manager := client.GetManager()
	etherbaseContract, err := etherbase.NewEtherbase(sr.etherbaseAddress, manager.Client())
	if err != nil {
		log.Printf("Failed to create etherbase contract: %v", err)
		return
	}
	sources, err := etherbaseContract.GetSources(&bind.CallOpts{
		From: sr.etherbaseAddress,
		BlockNumber: nil,
	})
	if err != nil {
		log.Printf("Failed to get sources: %v", err)
		return
	}
	for _, source := range sources {
		sr.addSource(source.SourceAddress, source.Owner)
	}
}

func (sr *SourceRegistry) startWatch() {
    ctx, cancel := context.WithCancel(context.Background())
    sr.stopWatch = cancel
    sr.watchDone.Add(1)

    go func() {
        defer sr.watchDone.Done()

        logsChan := make(chan types.Log)
        sub, err := sr.subscribeEtherbaseLogs(ctx, logsChan)
        if err != nil {
            log.Printf("Error subscribing to SourceCreated logs: %v", err)
            return
        }
        defer sub.Unsubscribe()

        // Define the event topic hash once
        sourceCreatedTopic := crypto.Keccak256Hash([]byte("SourceCreated(address,address)"))

        for {
            select {
            case l := <-logsChan:
                if len(l.Topics) < 2 {
                    continue
                }
                // make sure the event is SourceCreated
                if l.Topics[0] != sourceCreatedTopic {
                    continue
                }
                // parse SourceCreated event
                sourceAddr := common.BytesToAddress(l.Topics[1].Bytes()) 
                sr.addSource(sourceAddr, common.Address{})
            case err := <-sub.Err():
                log.Printf("SourceRegistry subscription error: %v", err)
                return
            case <-ctx.Done():
                return
            }
        }
    }()
}

func (sr *SourceRegistry) subscribeEtherbaseLogs(ctx context.Context, logsChan chan<- types.Log) (ethereum.Subscription, error) {
    manager := client.GetManager()
    sourceCreatedTopic := crypto.Keccak256Hash([]byte("SourceCreated(address,address)"))
    filter := ethereum.FilterQuery{
        Addresses: []common.Address{sr.etherbaseAddress},
        Topics:    [][]common.Hash{{sourceCreatedTopic}}, // Add the topic filter
    }
    return manager.Client().SubscribeFilterLogs(ctx, filter, logsChan)
}

func (sr *SourceRegistry) addSource(addr common.Address, owner common.Address) {
    // No need for mutex lock anymore
    if _, loaded := sr.sources.LoadOrStore(addr, owner); !loaded {
        // Only notify if this was a new entry
        sr.notify()
    }
}

func (sr *SourceRegistry) OnUpdate() chan []common.Address {
    ch := make(chan []common.Address, 1) // buffered channel to prevent blocking
    sr.mu.Lock()
    defer sr.mu.Unlock()
    sr.updateSubs = append(sr.updateSubs, ch)
    // Send initial state
    ch <- sr.GetSources()
    return ch
}

func (sr *SourceRegistry) notify() {
    addresses := []common.Address{}
    sr.sources.Range(func(key, value interface{}) bool {
        addresses = append(addresses, key.(common.Address))
        return true
    })
    addresses = append(addresses, sr.additionalContracts...)
    
    sr.mu.Lock()
    defer sr.mu.Unlock()
    for _, ch := range sr.updateSubs {
        select {
        case ch <- addresses:
        default:
            // skip if channel is full
        }
    }
}

func (sr *SourceRegistry) GetSources() []common.Address {
    var out []common.Address
    sr.sources.Range(func(key, value interface{}) bool {
        out = append(out, key.(common.Address))
        return true
    })
    out = append(out, sr.additionalContracts...)
    return out
}

func (sr *SourceRegistry) IsSource(addr common.Address) bool {
    _, found := sr.sources.Load(addr)
    return found
}

func (sr *SourceRegistry) Stop() {
    sr.watchOnce.Do(func() {
        sr.stopWatch()
        sr.watchDone.Wait()
    })
}

type SourceInfo struct {
    SourceAddress common.Address
    Owner         common.Address
}

func (sr *SourceRegistry) GetSourcesWithOwners() []SourceInfo {
    var out []SourceInfo
    sr.sources.Range(func(key, value interface{}) bool {
        out = append(out, SourceInfo{
            SourceAddress: key.(common.Address),
            Owner:         value.(common.Address),
        })
        return true
    })
    return out
}

type EventDefinition struct {
    Name           string     `json:"name"`
    Id             string     `json:"id"`
    Args           []Argument `json:"args"`
    EventTopic     string     `json:"eventTopic"`
    NumIndexedArgs uint8      `json:"numIndexedArgs"`
}

type Argument struct {
    Name      string `json:"name"`
    ArgType   string `json:"argType"`
    IsIndexed bool   `json:"isIndexed"`
}

func (sr *SourceRegistry) GetEventDefinitions(sourceAddr string) ([]EventDefinition, error) {
    address := common.HexToAddress(sourceAddr)
    if !sr.IsSource(address) {
        return nil, fmt.Errorf("address %s is not a registered source", sourceAddr)
    }

    manager := client.GetManager()
    sourceContract, err := etherbasesource.NewEtherbaseSource(address, manager.Client())
    if err != nil {
        return nil, fmt.Errorf("failed to create source contract: %v", err)
    }

    // Get registered event names
    names, err := sourceContract.GetRegisteredEventNames(&bind.CallOpts{})
    if err != nil {
        return nil, fmt.Errorf("failed to get event names: %v", err)
    }

    // Get schema for each event
    definitions := make([]EventDefinition, 0, len(names))
    for _, name := range names {
        schema, err := sourceContract.GetEventSchemaFromName(&bind.CallOpts{}, name)
        if err != nil {
            return nil, fmt.Errorf("failed to get schema for event %s: %v", name, err)
        }

        // Convert contract arguments to our Argument type
        args := make([]Argument, len(schema.Args))
        for i, arg := range schema.Args {
            args[i] = Argument{
                Name:      arg.Name,
                ArgType:   arg.ArgType,
                IsIndexed: arg.IsIndexed,
            }
        }

        definitions = append(definitions, EventDefinition{
            Name:           schema.Name,
            Id:             schema.Id,
            Args:           args,
            EventTopic:     common.Bytes2Hex(schema.EventTopic[:]),
            NumIndexedArgs: schema.NumIndexedArgs,
        })
    }

    return definitions, nil
}
