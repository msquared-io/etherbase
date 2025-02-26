package sources

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/msquared-io/etherbase/backend/internal/client"
	"github.com/msquared-io/etherbase/backend/internal/integration/go/etherbase"
	"github.com/msquared-io/etherbase/backend/internal/integration/go/etherbasesource"
)

// SourceRegistry polls the etherbase contract for sources and custom contracts
// and keeps track of them. It also notifies watchers about new sources.
type SourceRegistry struct {
    etherbaseAddress      common.Address
    customContracts       sync.Map // key: common.Address (contract), value: CustomContract

    sources     sync.Map // key: common.Address (source), value: common.Address (owner)
    sourcesSubs   []chan []SourceInfo
    customContractsSubs   []chan []etherbase.EtherbaseCustomContract
    stopPolling    context.CancelFunc
    pollingDone    sync.WaitGroup
    pollingOnce    sync.Once
    mu           sync.Mutex // kept only for updateSubs slice operations
}

var sourceRegistryInstance *SourceRegistry
var sourceRegistryOnce sync.Once

func NewSourceRegistry(etherbaseAddress common.Address) *SourceRegistry {
    sourceRegistryOnce.Do(func() {
        sourceRegistryInstance = &SourceRegistry{
            etherbaseAddress:      etherbaseAddress,
        }
        sourceRegistryInstance.startPolling()
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

func (sr *SourceRegistry) startPolling() {
    ctx, cancel := context.WithCancel(context.Background())
    sr.stopPolling = cancel
    sr.pollingDone.Add(1)

    go func() {
        defer sr.pollingDone.Done()
        ticker := time.NewTicker(1 * time.Second)
        defer ticker.Stop()

        log.Printf("[source registry] Starting polling for sources and custom contracts every 1 second")

        for {
            select {
            case <-ctx.Done():
                log.Printf("[source registry] Polling stopped")
                return
            case <-ticker.C:
                sr.pollForUpdates()
            }
        }
    }()
}

func (sr *SourceRegistry) pollForUpdates() {
    // Poll for new sources
    manager := client.GetManager()
    etherbaseContract, err := etherbase.NewEtherbase(sr.etherbaseAddress, manager.Client())
    if err != nil {
        log.Printf("[source registry] Failed to create etherbase contract: %v", err)
        return
    }

    // Check for new sources
    sources, err := etherbaseContract.GetSources(&bind.CallOpts{})
    if err != nil {
        log.Printf("[source registry] Failed to get sources: %v", err)
        return
    }

    // Track if we found any new sources
    foundNewSources := false
    for _, source := range sources {
        if _, loaded := sr.sources.LoadOrStore(source.SourceAddress, source.Owner); !loaded {
            log.Printf("[source registry] Found new source: address=%s owner=%s", 
                source.SourceAddress.Hex(), 
                source.Owner.Hex())
            foundNewSources = true
        }
    }

    // Notify if we found new sources
    if foundNewSources {
        sr.notifySources()
    }

    // Check for new custom contracts
    customContracts, err := etherbaseContract.GetCustomContracts(&bind.CallOpts{})
    if err != nil {
        log.Printf("[source registry] Failed to get custom contracts: %v", err)
        return
    }

    // Track if we found any new custom contracts
    foundNewCustomContracts := false
    
    // Create a map of existing contract addresses for quick lookup
    existingContracts := make(map[common.Address]bool)
    sr.customContracts.Range(func(key, value interface{}) bool {
        existingContracts[key.(common.Address)] = true
        return true
    })
    
    // Check for new contracts
    for _, contract := range customContracts {
        if !existingContracts[contract.ContractAddress] {
            sr.customContracts.Store(contract.ContractAddress, contract)
            log.Printf("[source registry] Found new custom contract: %s", contract.ContractAddress.Hex())
            foundNewCustomContracts = true
        }
    }
    
    // Check for deleted contracts
    sr.customContracts.Range(func(key, value interface{}) bool {
        contractAddr := key.(common.Address)
        found := false
        for _, contract := range customContracts {
            if contract.ContractAddress == contractAddr {
                found = true
                break
            }
        }
        if !found {
            sr.customContracts.Delete(contractAddr)
            log.Printf("[source registry] Custom contract deleted: %s", contractAddr.Hex())
            foundNewCustomContracts = true
        }
        return true
    })

    // Notify if we found new or deleted custom contracts
    if foundNewCustomContracts {
        sr.notifyCustomContracts()
    }
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
    sr.pollingOnce.Do(func() {
        sr.stopPolling()
        sr.pollingDone.Wait()
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
