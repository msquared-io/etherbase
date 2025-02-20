package subscription

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"

	"github.com/msquared-io/etherbase/backend/internal/schema"
	"github.com/msquared-io/etherbase/backend/internal/types"
)

// Client is per-WS-connection state
type Client struct {
    ID                types.ClientID
    eventSubs         map[types.SubscriptionID]*types.EventSubscription
    stateSubs         map[types.SubscriptionID]*types.StateSubscription
    mu                sync.RWMutex
}


type ListenEvent struct {
    Name string
    Args map[string]interface{}
}

func NewSubscriptionClient(id types.ClientID) *Client {
    return &Client{
        ID:        id,
        eventSubs: make(map[types.SubscriptionID]*types.EventSubscription),
        stateSubs: make(map[types.SubscriptionID]*types.StateSubscription),
    }
}

func (sc *Client) SetEventSubscription(subID types.SubscriptionID, sub *types.EventSubscription) {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    sc.eventSubs[subID] = sub
}

func (sc *Client) SetStateSubscription(subID types.SubscriptionID, sub *types.StateSubscription) {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    sc.stateSubs[subID] = sub
}

func (sc *Client) RemoveSubscription(subID types.SubscriptionID) {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    delete(sc.eventSubs, subID)
    delete(sc.stateSubs, subID)
}

func (sc *Client) GetEventSubscriptions() map[types.SubscriptionID]*types.EventSubscription {
    sc.mu.RLock()
    defer sc.mu.RUnlock()
    copy := make(map[types.SubscriptionID]*types.EventSubscription, len(sc.eventSubs))
    for k, v := range sc.eventSubs {
        copy[k] = v
    }
    return copy
}

func (sc *Client) GetStateSubscriptions() map[types.SubscriptionID]*types.StateSubscription {
    sc.mu.RLock()
    defer sc.mu.RUnlock()
    copy := make(map[types.SubscriptionID]*types.StateSubscription, len(sc.stateSubs))
    for k, v := range sc.stateSubs {
        copy[k] = v
    }
    return copy
}

func (sc *Client) GetStateSubscription(subID types.SubscriptionID) *types.StateSubscription {
    sc.mu.RLock()
    defer sc.mu.RUnlock()
    return sc.stateSubs[subID]
}

func (sc *Client) EventMatchesSubscription(event types.DecodedEvent, subscription *types.EventSubscription) bool {
    if len(subscription.ContractAddresses) > 0 {
        contractAddr := event.ContractAddress
        found := false
        for _, addr := range subscription.ContractAddresses {
            if addr.Hex() == contractAddr {
                found = true
                break
            }
        }
        if !found {
            return false
        }
    }

    if len(subscription.Events) > 0 {
        eventName := event.Name
        eventID := event.EventID
        found := false
        for _, subEvent := range subscription.Events {
            if subEvent.Name == eventName || subEvent.ID == eventID {
                found = true
                break
            }
        }
        if !found {
            return false
        }
    }

    return true
}

func (sc *Client) StateMatchesSubscription(update types.StateUpdate, subscription *types.StateSubscription) bool {
    for _, subPath := range subscription.StatePaths {
        if subscription.ContractAddress.Hex() == strings.ToLower(update.ContractAddress) &&
            reflect.DeepEqual(subPath, update.State) {
            return true
        }
    }
    return false
}

func (sc *Client) ClearSubscriptions() {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    sc.eventSubs = make(map[types.SubscriptionID]*types.EventSubscription)
    sc.stateSubs = make(map[types.SubscriptionID]*types.StateSubscription)
}

// ValidateEventSubscription checks if all events in the subscription have valid schemas
func (sc *Client) ValidateEventSubscription(subscription *types.MessageEventSubscription) (bool, error) {
    if len(subscription.Events) > 0 {
        for _, event := range subscription.Events {
            // check all contract addresses we're subscribing to have a schema for the event that matches the requested event
            for _, contractAddress := range subscription.ContractAddresses {
                if !schema.GetSchemaProvider().HasSchemaForEvent(contractAddress, event.Name) {
                    return false, fmt.Errorf(
                        "No schema found for event %q on contract %s. The event may not exist or may not be registered.",
                        event.Name,
                        contractAddress.Hex(),
                    )
                }
            }
            // todo validate args/ids
        }
    }
    return true, nil
}

// ValidateStateSubscriptionEvents checks if all listen events in the state subscription have valid schemas
func (sc *Client) ValidateStateSubscriptionEvents(subscription *types.MessageStateSubscription) (bool, error) {
    if subscription.Options != nil && subscription.Options.Repoll != nil && subscription.Options.Repoll.ListenEvents != nil && len(*subscription.Options.Repoll.ListenEvents) > 0 {
        for _, event := range *subscription.Options.Repoll.ListenEvents {
            log.Printf("validating event: %v", event)
            if !schema.GetSchemaProvider().HasSchemaForEvent(subscription.ContractAddress, event.Name) {
                return false, fmt.Errorf(
                    "No schema found for event %q on contract %s. The event may not exist or may not be registered.",
                    event.Name,
                    subscription.ContractAddress.Hex(),
                )
            }
        }
    }
    return true, nil
}