package subscription

import (
	"log"
	"reflect"
	"sync"

	"github.com/msquared-io/etherbase/backend/internal/types"

	"github.com/ethereum/go-ethereum/common"
)

type Manager struct {
    clients sync.Map
}

var managerInstance *Manager
var managerOnce sync.Once

func NewManager() *Manager {
    managerOnce.Do(func() {
        managerInstance = &Manager{}
    })
    return managerInstance
}

func GetManager() *Manager {
    return managerInstance
}

func (m *Manager) AddClient(sc *Client) {
    m.clients.Store(sc.ID, sc)
}

func (m *Manager) RemoveClient(id types.ClientID) {
    m.clients.Delete(id)
}

func (m *Manager) GetClient(id types.ClientID) *Client {
    if client, ok := m.clients.Load(id); ok {
        return client.(*Client)
    }
    return nil
}

// Example method that returns all clients
func (m *Manager) AllClients() []*Client {
    clients := make([]*Client, 0)
    m.clients.Range(func(key, value interface{}) bool {
        clients = append(clients, value.(*Client))
        return true
    })
    return clients
}

// Example usage
func (m *Manager) DebugPrint() {
    count := 0
    m.clients.Range(func(key, value interface{}) bool {
        count++
        return true
    })
    log.Printf("Total subscription clients: %d", count)
}

// GetSubscription returns a specific subscription for a client
func (m *Manager) GetSubscription(clientID types.ClientID, subscriptionID types.SubscriptionID) interface{} {
    if client := m.GetClient(clientID); client != nil {
        // Check both event and state subscriptions
        if sub := client.GetEventSubscriptions()[subscriptionID]; sub != nil {
            return sub
        }
        if sub := client.GetStateSubscriptions()[subscriptionID]; sub != nil {
            return sub
        }
    }
    return nil
}

type StringSliceSet struct {
	data [][]string
}

func NewStringSliceSet() *StringSliceSet {
	return &StringSliceSet{
		data: make([][]string, 0),
	}
}

// Contains checks for existence using deep equality.
func (s *StringSliceSet) Contains(slice []string) bool {
	for _, item := range s.data {
		if reflect.DeepEqual(item, slice) {
			return true
		}
	}
	return false
}

// Add inserts the slice if it doesn't already exist.
func (s *StringSliceSet) Add(slice []string) {
	if !s.Contains(slice) {
		s.data = append(s.data, slice)
	}
}


// GetAllUniqueStateSubscriptions returns all unique path/contract combinations
func (m *Manager) GetAllUniqueStateSubscriptions() map[common.Address]*StringSliceSet {
    uniqueSubs := make(map[common.Address]*StringSliceSet)
    
    m.clients.Range(func(_, value interface{}) bool {
        client := value.(*Client)
        for _, sub := range client.GetStateSubscriptions() {
            if _, exists := uniqueSubs[sub.ContractAddress]; !exists {
                uniqueSubs[sub.ContractAddress] = NewStringSliceSet()
            }
            for _, path := range sub.StatePaths {
                uniqueSubs[sub.ContractAddress].Add(path)
            }
        }
        return true
    })
    
    return uniqueSubs
}

// GetCompactUniqueStateSubscriptions returns the most compact and unique paths
func (m *Manager) GetCompactUniqueStateSubscriptions(stateSubscriptions map[types.ClientID][]types.SubscriptionID) map[common.Address]*StringSliceSet {
    uniquePaths := make(map[common.Address]*StringSliceSet)

    // Collect all paths
    for clientID, subIDs := range stateSubscriptions {
        for _, subID := range subIDs {
            if sub := m.GetSubscription(clientID, subID); sub != nil {
                if stateSub, ok := sub.(*types.StateSubscription); ok {
                    if _, exists := uniquePaths[stateSub.ContractAddress]; !exists {
                        uniquePaths[stateSub.ContractAddress] = NewStringSliceSet()
                    }
                    for _, path := range stateSub.StatePaths {
                        uniquePaths[stateSub.ContractAddress].Add(path)
                    }
                }
            } else {
                log.Printf("Subscription %s not found", subID)
            }
        }
    }

    // Remove paths that are prefixes of other paths
    for addr := range uniquePaths {
        paths := make([]types.StatePath, 0)
        for _, path := range uniquePaths[addr].data {
            paths = append(paths, path)
        }

        compactPaths := NewStringSliceSet()
        for i, path := range paths {
            isPrefix := false
            for j, otherPath := range paths {
                if i == j {
                    continue
                }
                if isPathPrefix(path, otherPath) {
                    isPrefix = true
                    break
                }
            }
            if isPrefix {
                compactPaths.Add(path)
            }
        }
        uniquePaths[addr] = compactPaths
    }

    return uniquePaths
}

func isPathPrefix(path1, path2 types.StatePath) bool {
    if len(path1) > len(path2) {
        return false
    }
    for i := range path1 {
        if path1[i] != path2[i] {
            return false
        }
    }
    return true
}

// Reset clears all clients
func (m *Manager) Reset() {
    m.clients = sync.Map{}
}
