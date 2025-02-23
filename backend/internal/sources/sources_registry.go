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
    customContracts       sync.Map // key: common.Address (contract), value: CustomContract

    sources     sync.Map // key: common.Address (source), value: common.Address (owner)
    sourcesSubs   []chan []SourceInfo
    customContractsSubs   []chan []etherbase.EtherbaseCustomContract
    stopWatch    context.CancelFunc
    watchDone    sync.WaitGroup
    watchOnce    sync.Once
    mu           sync.Mutex // kept only for updateSubs slice operations
}

var sourceRegistryInstance *SourceRegistry
var sourceRegistryOnce sync.Once

func NewSourceRegistry(etherbaseAddress common.Address) *SourceRegistry {
    sourceRegistryOnce.Do(func() {
        sourceRegistryInstance = &SourceRegistry{
            etherbaseAddress:      etherbaseAddress,
        }
        sourceRegistryInstance.startWatch()
		sourceRegistryInstance.loadInitialSources()
        sourceRegistryInstance.loadInitialCustomContracts()
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
	sources, err := etherbaseContract.GetSources(&bind.CallOpts{})
	if err != nil {
		log.Printf("Failed to get sources: %v", err)
		return
	}
	for _, source := range sources {
		sr.addSource(source.SourceAddress, source.Owner)
	}
}

func (sr *SourceRegistry) loadInitialCustomContracts() {
	manager := client.GetManager()
	etherbaseContract, err := etherbase.NewEtherbase(sr.etherbaseAddress, manager.Client())
	if err != nil {
		log.Printf("Failed to create etherbase contract: %v", err)
		return
	}
	customContracts, err := etherbaseContract.GetCustomContracts(&bind.CallOpts{})
	if err != nil {
		log.Printf("Failed to get custom contracts: %v", err)
		return
	}
	for _, contract := range customContracts {
		sr.addCustomContract(contract)
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
            log.Printf("Error subscribing to Etherbase logs for source registry: %v", err)
            return
        }
        defer sub.Unsubscribe()

        // Define the event topic hash once
        sourceCreatedTopic := crypto.Keccak256Hash([]byte("SourceCreated(address,address)"))
        customContractAddedTopic := crypto.Keccak256Hash([]byte("CustomContractAdded(address)"))
        customContractDeletedTopic := crypto.Keccak256Hash([]byte("CustomContractDeleted(address)"))

        for {
            select {
            case l := <-logsChan:
                if l.Topics[0] == sourceCreatedTopic {
                    sourceAddr := common.BytesToAddress(l.Topics[1].Bytes()) 
                    owner := common.BytesToAddress(l.Topics[2].Bytes())
                    sr.addSource(sourceAddr, owner)
                } else if l.Topics[0] == customContractAddedTopic {
                    contractAddr := common.BytesToAddress(l.Topics[1].Bytes())
                    manager := client.GetManager()
                    etherbaseContract, err := etherbase.NewEtherbase(sr.etherbaseAddress, manager.Client())
                    if err != nil {
                        log.Printf("Failed to create etherbase contract: %v", err)
                        return
                    }
                    contract, err := etherbaseContract.GetCustomContract(&bind.CallOpts{}, contractAddr)
                    if err != nil {
                        log.Printf("failed to get custom contract: %v", err)
                    }
                        
                    sr.addCustomContract(contract)
                } else if l.Topics[0] == customContractDeletedTopic {
                    contractAddr := common.BytesToAddress(l.Topics[1].Bytes())
                    sr.deleteCustomContract(contractAddr)
                }
            case err := <-sub.Err():
                log.Printf("SourceRegistry subscription error: %v", err)
                // return // dont return as we want to keep the subscription alive and retry
            case <-ctx.Done():
                return
            }
        }
    }()
}

func (sr *SourceRegistry) subscribeEtherbaseLogs(ctx context.Context, logsChan chan<- types.Log) (ethereum.Subscription, error) {
    manager := client.GetManager()
    filter := ethereum.FilterQuery{
        Addresses: []common.Address{sr.etherbaseAddress},
        Topics:    [][]common.Hash{},
    }
    return manager.Client().SubscribeFilterLogs(ctx, filter, logsChan)
}

func (sr *SourceRegistry) addSource(addr common.Address, owner common.Address) {
    if _, loaded := sr.sources.LoadOrStore(addr, owner); !loaded {
        // Only notify if this was a new entry
        sr.notifySources()
    }
}

func (sr *SourceRegistry) addCustomContract(customContract etherbase.EtherbaseCustomContract) {
    if _, loaded := sr.customContracts.LoadOrStore(customContract.ContractAddress, customContract); !loaded {
        // Only notify if this was a new entry
        sr.notifyCustomContracts()
    }
}

func (sr *SourceRegistry) deleteCustomContract(addr common.Address) {
    sr.customContracts.Delete(addr)
    // Only notify if this was a new entry
    sr.notifyCustomContracts()
}

func (sr *SourceRegistry) OnUpdateSources() chan []SourceInfo {
    ch := make(chan []SourceInfo, 1) // buffered channel to prevent blocking
    sr.mu.Lock()
    defer sr.mu.Unlock()
    sr.sourcesSubs = append(sr.sourcesSubs, ch)
    // Send initial state
    ch <- sr.GetSources()
    return ch
}

func (sr *SourceRegistry) OnUpdateCustomContracts() chan []etherbase.EtherbaseCustomContract {
    ch := make(chan []etherbase.EtherbaseCustomContract, 1) // buffered channel to prevent blocking
    sr.mu.Lock()
    defer sr.mu.Unlock()
    sr.customContractsSubs = append(sr.customContractsSubs, ch)
    // Send initial state
    ch <- sr.GetCustomContracts()
    return ch
}

func (sr *SourceRegistry) notifySources() {
    sources := []SourceInfo{}
    sr.sources.Range(func(key, value interface{}) bool {
        sources = append(sources, SourceInfo{
            SourceAddress: key.(common.Address),
            Owner:         value.(common.Address),
        })
        return true
    })
    
    sr.mu.Lock()
    defer sr.mu.Unlock()
    for _, ch := range sr.sourcesSubs {
        select {
            case ch <- sources:
        default:
            // skip if channel is full
        }
    }
}

func (sr *SourceRegistry) notifyCustomContracts() {
    customContracts := []etherbase.EtherbaseCustomContract{}
    sr.customContracts.Range(func(key, value interface{}) bool {
        customContracts = append(customContracts, value.(etherbase.EtherbaseCustomContract))
        return true
    })

    sr.mu.Lock()
    defer sr.mu.Unlock()
    for _, ch := range sr.customContractsSubs {
        select {
            case ch <- customContracts:
        default:
            // skip if channel is full
        }
    }
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

func (sr *SourceRegistry) GetSources() []SourceInfo {
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

func (sr *SourceRegistry) GetCustomContracts() []etherbase.EtherbaseCustomContract {
    var out []etherbase.EtherbaseCustomContract
    sr.customContracts.Range(func(key, value interface{}) bool {
        out = append(out, value.(etherbase.EtherbaseCustomContract))
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
