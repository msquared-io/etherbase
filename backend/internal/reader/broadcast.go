package reader

import (
	"log"

	"github.com/msquared-io/etherbase/backend/internal/streaming"
	"github.com/msquared-io/etherbase/backend/internal/subscription"
	"github.com/msquared-io/etherbase/backend/internal/types"

	"github.com/gorilla/websocket"
)

type StateUpdate struct {
	ContractAddress   string                 `json:"contractAddress"`
	State            map[string]interface{} `json:"state"`
	TransactionIndex uint                   `json:"transactionIndex,omitempty"`
}

func buildUpdates(client *subscription.Client, payload *streaming.ViewCallback) []types.Update {
	var updates []types.Update

	// Process events
	decodedEvents := make([]types.DecodedEvent, 0, len(payload.EventUpdates))
	for _, eventUpdate := range payload.EventUpdates {
		// combine indexed and non-indexed args - the user can process them their way since they know they're indexed or non-indexed
		args := make(map[string]interface{})
		for k, v := range eventUpdate.Event.IndexedArgs {
			args[k] = v
		}
		for k, v := range eventUpdate.Event.NonIndexedArgs {
			args[k] = v
		}

		decodedEvents = append(decodedEvents, types.DecodedEvent{
			ContractAddress: eventUpdate.Event.ContractAddress.String(),
			Name:           eventUpdate.Event.Name,
			Args:           args,
			EventID:        eventUpdate.Event.ID,
			Block:          payload.BlockNumber,
		})
	}

	// Add matching events to updates
	for subscriptionID, subscription := range client.GetEventSubscriptions() {
		for _, event := range decodedEvents {
			if client.EventMatchesSubscription(event, subscription) {
				updates = append(updates, types.Update{
					Type:           "event",
					SubscriptionID: string(subscriptionID),
					Data:          event,
				})
			}
		}
	}

	// Process state updates
	if len(payload.StateUpdates) > 0 {
		// Group updates by subscriptionId and transaction
		updatesBySubAndTx := make(map[types.SubscriptionID]map[uint][]streaming.SubscriberStateUpdate)
		for _, update := range payload.StateUpdates {
			subID := types.SubscriptionID(update.SubscriptionID)
			if _, exists := updatesBySubAndTx[subID]; !exists {
				updatesBySubAndTx[subID] = make(map[uint][]streaming.SubscriberStateUpdate)
			}
			txUpdates := updatesBySubAndTx[subID]
			txUpdates[update.TransactionIndex] = append(txUpdates[update.TransactionIndex], update)
		}

		// Process each subscription's updates
		for subscriptionID, txUpdates := range updatesBySubAndTx {
			subscription := client.GetStateSubscription(subscriptionID)
			if subscription == nil {
				continue
			}
 
			if subscription.Options.CombineStateUpdatesFromSameContract {
				// Combine updates from the same contract
				combinedState := make(map[string]interface{})
				for _, stateUpdates := range txUpdates {
					for _, update := range stateUpdates {
						deepMerge(combinedState, update.State)
					}
				}

				updates = append(updates, types.Update{
					Type:           "state",
					SubscriptionID: string(subscriptionID),
					Data: types.StateUpdate{
						ContractAddress:   subscription.ContractAddress.String(),
						State:            combinedState,
					},
				})
			} else if subscription.Options.CombineStateUpdatesFromSameTransaction {
				// Combine updates from the same transaction
				for txIndex, stateUpdates := range txUpdates {
					combinedState := make(map[string]interface{})
					for _, update := range stateUpdates {
						deepMerge(combinedState, update.State)
					}
					
					updates = append(updates, types.Update{
						Type:           "state",
						SubscriptionID: string(subscriptionID),
						Data: types.StateUpdate{
							ContractAddress:   stateUpdates[0].ContractAddress.String(),
							State:            combinedState,
							TransactionIndex: txIndex,
						},
					})
				}
			} else {
				// Add each update individually
				for _, updateGroup := range txUpdates {
					for _, update := range updateGroup {
						updates = append(updates, types.Update{
							Type:           "state",
							SubscriptionID: string(subscriptionID),
							Data: types.StateUpdate{
								ContractAddress:   update.ContractAddress.String(),
								State:            update.State,
								TransactionIndex: update.TransactionIndex,
							},
						})
					}
				}
			}
		}
	}

	return updates
}

// deepMerge merges source map into target map, recursively
func deepMerge(target, source map[string]interface{}) {
	for key, sourceVal := range source {
		if targetVal, exists := target[key]; exists {
			// If both values are maps, merge them recursively
			if sourceMap, isSourceMap := sourceVal.(map[string]interface{}); isSourceMap {
				if targetMap, isTargetMap := targetVal.(map[string]interface{}); isTargetMap {
					deepMerge(targetMap, sourceMap)
					continue
				}
			}
		}
		// For non-map values or if target key doesn't exist, just copy the value
		target[key] = sourceVal
	}
}

// BroadcastUpdates sends updates to all connected clients
func BroadcastUpdates(payload *streaming.ViewCallback) {
	subscribers.Range(func(key, value interface{}) bool {
		conn := key.(*websocket.Conn)
		client := value.(*subscription.Client)

		updates := buildUpdates(client, payload)
		if len(updates) > 0 {
			log.Printf("Sending %d updates to client %s for block %d", len(updates), conn.RemoteAddr(), payload.BlockNumber)
			sendMessage(conn, "updates", map[string]interface{}{
				"updates": updates,
				"block":   payload.BlockNumber,
			})
			log.Printf("Sent %d updates to client %s for block %d", len(updates), conn.RemoteAddr(), payload.BlockNumber)
		}

		return true
	})

	log.Printf("Broadcasted updates for block %d", payload.BlockNumber)
}
