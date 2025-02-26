package state

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"sync"

	"github.com/msquared-io/etherbase/backend/internal/subscription"
	"github.com/msquared-io/etherbase/backend/internal/types"

	"github.com/ethereum/go-ethereum/common"
)

// StateManager caches state in memory and merges updates
type StateManager struct {
	subscriptionsContractState sync.Map // map[types.SubscriptionID]map[common.Address]map[string]interface{}
}

var (
	instance *StateManager
	once     sync.Once
)

// GetManager returns the singleton instance of StateManager
func GetManager() *StateManager {
	once.Do(func() {
		instance = &StateManager{}
	})
	return instance
}

// FetchStateForSubscriptions fetches and returns state for given subscriptions
func (sm *StateManager) FetchStateForSubscriptions(ctx context.Context, stateSubscriptions map[types.ClientID][]types.SubscriptionID, blockNumber *big.Int) (map[types.SubscriptionID]map[string]interface{}, uint64, error) {
	uniqueSubscriptions := sm.getCompactUniqueStateSubscriptions(stateSubscriptions)

	// Fetch state from chain
	result, err := sm.fetchStateFromChain(ctx, uniqueSubscriptions, blockNumber)
	if err != nil {
		return nil, 0, err
	}

	if result == nil {
		return nil, 0, nil
	}

	// log.Printf("Fetched %d states, %v, %v\n", len(result.States), stateSubscriptions, result.States)

	// Build state updates for each subscription
	stateUpdates := sm.buildStateUpdatesForSubscriptions(stateSubscriptions, result.States)

	// log.Printf("stateUpdates %v\n", stateUpdates)

	return stateUpdates, result.BlockNumber, nil
}

// getUniqueStateSubscriptions returns a map of contract addresses to sets of state paths
func (sm *StateManager) getCompactUniqueStateSubscriptions(stateSubscriptions map[types.ClientID][]types.SubscriptionID) map[common.Address][][]string {
	uniqueSubs := make(map[common.Address]map[string][]string)
	
	subscriptionManager := subscription.GetManager()
	
	for clientID, subIDs := range stateSubscriptions {
		client := subscriptionManager.GetClient(clientID)
		if client == nil {
			continue
		}

		for _, subID := range subIDs {
			stateSub := client.GetStateSubscription(subID)
			if stateSub == nil {
				continue
			}

			if uniqueSubs[stateSub.ContractAddress] == nil {
				uniqueSubs[stateSub.ContractAddress] = make(map[string][]string)
			}
			
			for _, path := range stateSub.StatePaths {
				// Use string representation of path slice as map key
				pathKey := fmt.Sprintf("%v", path)
				uniqueSubs[stateSub.ContractAddress][pathKey] = path
			}
		}
	}

	// Convert to final format
	result := make(map[common.Address][][]string)
	for addr, paths := range uniqueSubs {
		result[addr] = make([][]string, 0, len(paths))
		for _, path := range paths {
			result[addr] = append(result[addr], path)
		}
	}

	return result
}

func (sm *StateManager) fetchStateFromChain(ctx context.Context, stateSubscriptions map[common.Address][][]string, blockNumber *big.Int) (*BatchFetchResult, error) {
	requests := make([]StateFetchRequest, 0)
	
	for contractAddr, paths := range stateSubscriptions {
		for _, path := range paths {
			requests = append(requests, StateFetchRequest{
				ContractAddress: contractAddr,
				Path:           path,
			})
		}
	}

	return BatchFetchContractStates(ctx, requests, blockNumber)
}

func (sm *StateManager) buildStateUpdatesForSubscriptions(stateSubscriptions map[types.ClientID][]types.SubscriptionID, states []StateFetchResponse) map[types.SubscriptionID]map[string]interface{} {
	updates := make(map[types.SubscriptionID]map[string]interface{})
	
	subscriptionManager := subscription.GetManager()

	for clientID, subIDs := range stateSubscriptions {
		client := subscriptionManager.GetClient(clientID)
		if client == nil {
			continue
		}

		for _, subID := range subIDs {
			stateSub := client.GetStateSubscription(subID)
			if stateSub == nil {
				continue
			}

			fullStateUpdateOnChange := stateSub.Options.FullStateUpdateOnChange

			subscriberContractsStateVal, ok := sm.subscriptionsContractState.Load(subID)
			var contractsState map[common.Address]map[string]interface{}
			if ok {
				contractsState = subscriberContractsStateVal.(map[common.Address]map[string]interface{})
			} else {
				contractsState = make(map[common.Address]map[string]interface{})
			}

			subscriberContractStateVal, ok := contractsState[stateSub.ContractAddress]
			if !ok {
				subscriberContractStateVal = make(map[string]interface{})
			}

			for _, path := range stateSub.StatePaths {
				// log.Printf("sub: %v, path: %v", stateSub, path)
				contractCurrentValue := getValueAtPath(subscriberContractStateVal, path)
				isValueUpdated := false
				var updatedValue interface{}

				for _, state := range states {
					if state.ContractAddress == stateSub.ContractAddress {
						updatedValue = getValueAtPath(state.State, path)
						// log.Printf("state: %v, updatedValue: %v, currentValue: %v", state, updatedValue, contractCurrentValue)
						if updatedValue != nil {
							isValueUpdated = !reflect.DeepEqual(updatedValue, contractCurrentValue)
							if isValueUpdated {
								break
							}
						}
					}
				}

				// log.Printf("isValueUpdated: %v, fullStateUpdateOnChange: %v", isValueUpdated, fullStateUpdateOnChange)
				if fullStateUpdateOnChange || isValueUpdated {
					if updates[subID] == nil {
						updates[subID] = make(map[string]interface{})
					}
					updates[subID] = setValueAtPath(updates[subID], path, updatedValue)
					subscriberContractStateVal = setValueAtPath(subscriberContractStateVal, path, updatedValue)
				}
			}

			contractsState[stateSub.ContractAddress] = subscriberContractStateVal
			sm.subscriptionsContractState.Store(subID, contractsState)
		}
	}

	return updates
}

func getValueAtPath(obj map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return obj
	}

	current := obj
	for i, key := range path {
		if i == len(path)-1 {
			return current[key]
		}
		
		next, ok := current[key].(map[string]interface{})
		if !ok {
			return nil
		}
		current = next
	}
	return nil
}

func setValueAtPath(obj map[string]interface{}, path []string, value interface{}) map[string]interface{} {
	if len(path) == 0 {
		if valueMap, ok := value.(map[string]interface{}); ok {
			return valueMap
		}
		log.Printf("Warning: Cannot set non-map value %v to root object", value)
		return obj
	}

	current := obj
	for i, key := range path {
		if i == len(path)-1 {
			current[key] = value
			return obj
		}
		
		if current[key] == nil {
			current[key] = make(map[string]interface{})
		}
		current = current[key].(map[string]interface{})
	}
	return obj
}

func deepMerge(target, source map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	
	for k, v := range target {
		result[k] = v
	}
	
	for k, v := range source {
		if targetMap, isMap := result[k].(map[string]interface{}); isMap {
			if sourceMap, ok := v.(map[string]interface{}); ok {
				result[k] = deepMerge(targetMap, sourceMap)
				continue
			}
		}
		result[k] = v
	}
	
	return result
}
