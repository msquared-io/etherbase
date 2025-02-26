package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/msquared-io/etherbase/backend/internal/abiparser"
	"github.com/msquared-io/etherbase/backend/internal/integration/go/etherbase"
	"github.com/msquared-io/etherbase/backend/internal/integration/go/etherbasesource"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EventSchema holds metadata about an event
type EventSchema struct {
    Name           string
    ID             string
    Args           []Argument
    EventTopic     common.Hash
    NumIndexedArgs int
	Abi            *abi.Event
}

// Argument represents a single argument in an event.
type Argument struct {
    Name      string
    ArgType   string
    IsIndexed bool  // renamed from Indexed to match TS
}

// SchemaProvider is responsible for tracking known contract ABIs and event schemas
type SchemaProvider struct {
    contracts      sync.Map // map[common.Address]*ContractSchema
    globalSchemas  sync.Map // map[string]*EventSchema
    etherbaseAddress common.Address
    schemaWatchers sync.Map // map[common.Address]func() // for cleanup functions
    watchStop      context.CancelFunc
    watchDone      sync.WaitGroup
    ethClient      *ethclient.Client
}

type ContractSchema struct {
    Name   string
    ABI    *abi.ABI
    Events map[string]*EventSchema  // eventID -> schema
    Names  map[string]string        // eventName -> eventID
}

var instance *SchemaProvider
var once sync.Once

var validAbiTypes = map[string]bool{
    "address":     true,
    "bool":        true,
    "bytes":       true,
    "bytes32":     true,
    "bytes32[]":   true,
    "int256":      true,
    "int256[]":    true,
    "string":      true,
    "string[]":    true,
    "uint256":     true,
    "uint256[]":   true,
    "uint8":       true,
    "uint8[]":     true,
    "uint16":      true,
    "uint16[]":    true,
    "uint32":      true,
    "uint32[]":    true,
    "uint64":      true,
    "uint64[]":    true,
    "uint128":     true,
    "uint128[]":   true,
}

func NewSchemaProvider(etherbaseAddr common.Address, ethClient *ethclient.Client) *SchemaProvider {
    once.Do(func() {
        sp := &SchemaProvider{
            etherbaseAddress: etherbaseAddr,
            ethClient:     ethClient,
        }
        instance = sp
        sp.startEtherbaseWatcher()
    })
    return instance
}

func GetSchemaProvider() *SchemaProvider {
    return instance
}

// First, add this helper method to decode event data
func decodeEventData(data []byte) (string, error) {
    // Decode a single string parameter from event data
    stringType, _ := abi.NewType("string", "", nil)
    arguments := abi.Arguments{{Type: stringType}}
    decoded, err := arguments.Unpack(data)
    if err != nil {
        return "", err
    }
    if len(decoded) == 0 {
        return "", fmt.Errorf("no data decoded")
    }
    return decoded[0].(string), nil
}

func getAbiFromEventSchema(eventSchema *etherbasesource.EventSchema) (*abi.Event, error) {
	eventAbi := fmt.Sprintf("event %s(", eventSchema.Name)
	for _, arg := range eventSchema.Args {
		eventAbi += fmt.Sprintf("%s ", arg.ArgType)
		if arg.IsIndexed {
			eventAbi += "indexed "
		}
		eventAbi += fmt.Sprintf("%s, ", arg.Name)
	}
	eventAbi = eventAbi[:len(eventAbi)-2] + ")"
	abiElement, err := abiparser.ParseEventSignature(eventAbi)
	if err != nil {
		return nil, err
	}
	abiString, err := json.Marshal(abiElement)
	if err != nil {
		return nil, err
	}

	var outAbiEvent abi.Event
	err = json.Unmarshal(abiString, &outAbiEvent)
	if err != nil {
		return nil, err
	}

	return &outAbiEvent, nil
}

// Then update the event handler in AddSource
func (sp *SchemaProvider) AddSource(addr common.Address) error {
    // Check if already exists
    if _, exists := sp.contracts.Load(addr); !exists {
		sourceABI, err := abi.JSON(strings.NewReader(etherbasesource.EtherbaseSourceABI))
        if err != nil {
            return fmt.Errorf("failed to parse source ABI: %w", err)
        }
		ethDbAbi, err := abi.JSON(strings.NewReader(etherbasesource.IEthDBUpdaterABI))
		if err != nil {
			return fmt.Errorf("failed to parse ethDB ABI: %w", err)
		}
		// combine the two event abis
		for name, event := range ethDbAbi.Events {
			sourceABI.Events[name] = event
		}

        c := &ContractSchema{
            Name:   "EtherbaseSource",
            ABI:    &sourceABI,
            Events: make(map[string]*EventSchema),
            Names:  make(map[string]string),
        }

        // Add built-in events
        for _, event := range sourceABI.Events {
            es, err := buildEventSchema(event)
            if err != nil {
                log.Printf("Error building event schema for %s: %v", event.Name, err)
                continue
            }
            c.Events[es.ID] = es
            c.Names[es.Name] = es.ID
        }

        sp.contracts.Store(addr, c)

        // Set up schema poller if not already watching
        if _, exists := sp.schemaWatchers.Load(addr); !exists {
            log.Printf("Starting schema poller for source %s", addr.Hex())
            
            // Create context for this poller
            ctx, cancel := context.WithCancel(context.Background())
            
            // Store cleanup function
            sp.schemaWatchers.Store(addr, cancel)

            // Start polling for schemas
            go sp.startSchemaPoller(ctx, addr)
        }
    }
    return nil
}

// startSchemaPoller polls for new schemas every second
func (sp *SchemaProvider) startSchemaPoller(ctx context.Context, addr common.Address) {
    sp.watchDone.Add(1)
    defer sp.watchDone.Done()

    // Create contract binding for the source
    source, err := etherbasesource.NewEtherbaseSource(addr, sp.ethClient)
    if err != nil {
        log.Printf("Failed to create source binding: %v", err)
        return
    }

    // Get initial event names and schemas
    knownEventNames := make(map[string]bool)
    
    // Load initial schemas
    eventNames, err := source.GetRegisteredEventNames(&bind.CallOpts{})
    if err != nil {
        log.Printf("Failed to get registered event names: %v", err)
    } else {
        for _, eventName := range eventNames {
            knownEventNames[eventName] = true
            schema, err := source.GetEventSchemaFromName(&bind.CallOpts{}, eventName)
            if err != nil {
                log.Printf("Failed to get schema for event %s: %v", eventName, err)
                continue
            }

            abi, err := getAbiFromEventSchema(&schema)
            if err != nil {
                log.Printf("Failed to get ABI for event %s: %v", eventName, err)
                continue
            }
            
            eventSchema := &EventSchema{
                Name:           schema.Name,
                ID:             schema.Id,
                EventTopic:     schema.EventTopic,
                NumIndexedArgs: int(schema.NumIndexedArgs),
                Abi:            abi,
            }

            // Convert args
            for _, arg := range schema.Args {
                eventSchema.Args = append(eventSchema.Args, Argument{
                    Name:      arg.Name,
                    ArgType:   arg.ArgType,
                    IsIndexed: arg.IsIndexed,
                })
            }

            // Store the schema
            if c, ok := sp.contracts.Load(addr); ok {
                contract := c.(*ContractSchema)
                contract.Events[eventSchema.ID] = eventSchema
                contract.Names[eventSchema.Name] = eventSchema.ID
            }
        }
    }

    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            log.Printf("Schema poller for %s stopped", addr.Hex())
            return
        case <-ticker.C:
            // Poll for new event names
            eventNames, err := source.GetRegisteredEventNames(&bind.CallOpts{})
            if err != nil {
                log.Printf("Failed to poll registered event names: %v", err)
                continue
            }

            // Check for new event names
            for _, eventName := range eventNames {
                if !knownEventNames[eventName] {
                    // New event found
                    knownEventNames[eventName] = true
                    log.Printf("New event schema detected: %s", eventName)
                    
                    schema, err := source.GetEventSchemaFromName(&bind.CallOpts{}, eventName)
                    if err != nil {
                        log.Printf("Failed to get schema for event %s: %v", eventName, err)
                        continue
                    }

                    abi, err := getAbiFromEventSchema(&schema)
                    if err != nil {
                        log.Printf("Failed to get ABI for event %s: %v", eventName, err)
                        continue
                    }
                    
                    eventSchema := &EventSchema{
                        Name:           schema.Name,
                        ID:             schema.Id,
                        EventTopic:     schema.EventTopic,
                        NumIndexedArgs: int(schema.NumIndexedArgs),
                        Abi:            abi,
                    }

                    // Convert args
                    for _, arg := range schema.Args {
                        eventSchema.Args = append(eventSchema.Args, Argument{
                            Name:      arg.Name,
                            ArgType:   arg.ArgType,
                            IsIndexed: arg.IsIndexed,
                        })
                    }

                    // Store the schema
                    if c, ok := sp.contracts.Load(addr); ok {
                        contract := c.(*ContractSchema)
                        contract.Events[eventSchema.ID] = eventSchema
                        contract.Names[eventSchema.Name] = eventSchema.ID
                    }
                }
            }
        }
    }
}

// Example method for adding custom contract
func (sp *SchemaProvider) AddCustomContract(addr common.Address, name string, abiJSON string) error {
    a, err := abi.JSON(strings.NewReader(abiJSON))
    if err != nil {
        return err
    }
    
    c := &ContractSchema{
        Name:   name,
        ABI:    &a,
        Events: make(map[string]*EventSchema),
        Names:  make(map[string]string),
    }

    // Parse all events from the ABI
    for _, event := range c.ABI.Events {
        es, err := buildEventSchema(event)
        if err != nil {
            log.Printf("Error building event schema for %s: %v", event.Name, err)
            continue
        }
        c.Events[es.ID] = es
        c.Names[es.Name] = es.ID
		fmt.Printf("Added event %s\n", es.Name)
    }

    sp.contracts.Store(addr, c)
	fmt.Printf("Added custom contract %s\n", addr.Hex())
    return nil
}

func buildEventSchema(e abi.Event) (*EventSchema, error) {
    var args []Argument
    for _, input := range e.Inputs {
        args = append(args, Argument{
            Name:    input.Name,
            ArgType: input.Type.String(),
            IsIndexed: input.Indexed,
        })
    }


	indexedCount := 0
    for _, input := range e.Inputs {
        if input.Indexed {
            indexedCount++
        }
    }
    eventID := fmt.Sprintf("%s:%d", e.ID.Hex(), indexedCount)

    return &EventSchema{
        Name:           e.Name,
        ID:             eventID,
        Args:           args,
        EventTopic:     e.ID,
        NumIndexedArgs: indexedCount,
		Abi: &e,
    }, nil
}

func (sp *SchemaProvider) GetContractABI(addr common.Address) *abi.ABI {
    if c, ok := sp.contracts.Load(addr); ok {
        return c.(*ContractSchema).ABI
    }
    return nil
}

// Example for retrieving an event schema from name
func (sp *SchemaProvider) GetEventSchemaByName(addr common.Address, eventName string) *EventSchema {
    if c, ok := sp.contracts.Load(addr); ok {
        contract := c.(*ContractSchema)
        if eventID, exists := contract.Names[eventName]; exists {
            return contract.Events[eventID]
        }
    }
    return nil
}

func (sp *SchemaProvider) startEtherbaseWatcher() {
    // Create a context for watching
    ctx, cancel := context.WithCancel(context.Background())
    sp.watchStop = cancel
    
    sp.watchDone.Add(1)
    go func() {
        defer sp.watchDone.Done()
        
        // Start polling for custom contracts
        ticker := time.NewTicker(1 * time.Second)
        defer ticker.Stop()
        
        // Get initial custom contracts
        if err := sp.pollCustomContracts(); err != nil {
            log.Printf("Error in initial custom contracts poll: %v", err)
        }
        
        for {
            select {
            case <-ctx.Done():
                log.Printf("Etherbase watcher stopped")
                return
            case <-ticker.C:
                if err := sp.pollCustomContracts(); err != nil {
                    log.Printf("Error polling custom contracts: %v", err)
                }
            }
        }
    }()
}

// pollCustomContracts fetches all custom contracts from the Etherbase contract
func (sp *SchemaProvider) pollCustomContracts() error {
    etherbase, err := etherbase.NewEtherbase(sp.etherbaseAddress, sp.ethClient)
    if err != nil {
        return fmt.Errorf("failed to create etherbase binding: %w", err)
    }
    
    // Get all custom contracts
    customContracts, err := etherbase.GetCustomContracts(&bind.CallOpts{})
    if err != nil {
        return fmt.Errorf("failed to get custom contracts: %w", err)
    }
    
    // Track current contracts to detect deletions
    currentContracts := make(map[common.Address]bool)
    
    // Process each contract
    for _, contract := range customContracts {
        contractAddr := contract.ContractAddress
        currentContracts[contractAddr] = true
        
        // Check if already exists
        if _, exists := sp.contracts.Load(contractAddr); exists {
            continue
        }
        
        // Add the new contract
        if err := sp.AddCustomContract(contractAddr, "CustomContract", contract.ContractABI); err != nil {
            log.Printf("Failed to add custom contract %s: %v", contractAddr.Hex(), err)
        } else {
            log.Printf("Added custom contract %s", contractAddr.Hex())
        }
    }
    
    // Check for deleted contracts
    sp.contracts.Range(func(key, value interface{}) bool {
        addr := key.(common.Address)
        // Skip etherbase address and source contracts
        if addr == sp.etherbaseAddress || sp.isSourceContract(addr) {
            return true
        }
        
        if !currentContracts[addr] {
            // Contract was deleted
            sp.contracts.Delete(addr)
            log.Printf("Deleted custom contract %s", addr.Hex())
        }
        return true
    })
    
    return nil
}

// isSourceContract checks if the address is a source contract
func (sp *SchemaProvider) isSourceContract(addr common.Address) bool {
    // Iterate through all contracts to find sources
    var isSource bool
    sp.contracts.Range(func(key, value interface{}) bool {
        contract := value.(*ContractSchema)
        if contract.Name == "EtherbaseSource" {
            isSource = true
            return false // stop iteration
        }
        return true
    })
    return isSource
}

// Modify the schema watcher in AddSource
func (sp *SchemaProvider) startSchemaWatcher(addr common.Address) error {
    // Check if already watching
    if _, exists := sp.schemaWatchers.Load(addr); exists {
        return nil
    }

    // Create context for this watcher
    ctx, cancel := context.WithCancel(context.Background())
    
    // Store cleanup function
    sp.schemaWatchers.Store(addr, cancel)
    
    // Start the schema poller
    go sp.startSchemaPoller(ctx, addr)
    
    return nil
}

func (sp *SchemaProvider) StopWatching() {
    if sp.watchStop != nil {
        sp.watchStop()
        sp.watchDone.Wait()
    }
    
    // Clean up all contract watchers
    sp.schemaWatchers.Range(func(key, value interface{}) bool {
        if cleanup, ok := value.(func()); ok {
            cleanup()
        }
        return true
    })
    sp.schemaWatchers = sync.Map{}
}

func (sp *SchemaProvider) Reset() {
    sp.StopWatching()
    sp.contracts = sync.Map{}
    sp.globalSchemas = sync.Map{}
    instance = nil
}

func (sp *SchemaProvider) validateEventSchema(schema *EventSchema) bool {
    for _, arg := range schema.Args {
        if !validAbiTypes[arg.ArgType] {
            log.Printf("Invalid ABI type: %s", arg.ArgType)
            return false
        }
    }
    return true
}

func (sp *SchemaProvider) getEventTopic(name string, args []Argument) common.Hash {
    // Build event signature
    sig := fmt.Sprintf("%s(%s)", name, sp.getEventArgsSignature(args))
    return crypto.Keccak256Hash([]byte(sig))
}

func (sp *SchemaProvider) getEventArgsSignature(args []Argument) string {
    types := make([]string, len(args))
    for i, arg := range args {
        types[i] = arg.ArgType
    }
    return strings.Join(types, ",")
}

func (sp *SchemaProvider) getEventID(topic common.Hash, numIndexedArgs int) string {
    return fmt.Sprintf("%s:%d", topic.Hex(), numIndexedArgs)
}

// GetEventSchemaFromID retrieves an event schema by its ID, checking both contract-specific
// and global schemas. If the address is provided, it first checks the contract schemas,
// then falls back to global schemas.
func (sp *SchemaProvider) GetEventSchemaFromID(addr common.Address, eventID string) *EventSchema {
    // First check contract-specific schemas
    if c, ok := sp.contracts.Load(addr); ok {
        contract := c.(*ContractSchema)
        if schema, exists := contract.Events[eventID]; exists {
            return schema
        }
    }

    // Fall back to global schemas
    if schema, ok := sp.globalSchemas.Load(eventID); ok {
        return schema.(*EventSchema)
    }

    return nil
}

func (sp *SchemaProvider) HasSchemaForEvent(addr common.Address, eventName string) bool {
    if c, ok := sp.contracts.Load(addr); ok {
        contract := c.(*ContractSchema)
        return contract.Names[eventName] != ""
    }
    return false
}