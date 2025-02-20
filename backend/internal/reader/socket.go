package reader

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/oklog/ulid/v2"

	"github.com/msquared-io/etherbase/backend/internal/subscription"
	"github.com/msquared-io/etherbase/backend/internal/types"
)

var (
	subscribers = &sync.Map{} // maps *websocket.Conn to *subscription.Client
)

// Add the upgrader as a package-level variable that can be set from server.go
var Upgrader *websocket.Upgrader

// Message types
type MessageType string

const (
	TypeSubscribe    MessageType = "subscribe"
	TypeUnsubscribe MessageType = "unsubscribe"
	TypeUpdates     MessageType = "updates"
	TypeError       MessageType = "error"
	TypeInitialState MessageType = "initial_state"
)

type ClientMessage struct {
	Type MessageType     `json:"type"`
	Data json.RawMessage `json:"data"`
}

type SubscribeMessage struct {
	EventSubscription *types.MessageEventSubscription    `json:"eventSubscription,omitempty"`
	StateSubscription *types.MessageStateSubscription    `json:"stateSubscription,omitempty"`
	PendingID        string                `json:"pendingId"`
}


func HandleReadWebSocket(conn *websocket.Conn, r *http.Request) {
	// Set read limit
	conn.SetReadLimit(512 * 1024) // 512KB

	// Configure read deadline and pong handler
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	// Start ping ticker
	done := make(chan struct{})
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(10*time.Second)); err != nil {
					log.Printf("[websocket-reader] Ping failed: %v", err)
					close(done)
					return
				}
			case <-done:
				return
			}
		}
	}()

	// Existing client setup code...
	clientID := ulid.Make().String()
	subClient := subscription.NewSubscriptionClient(types.ClientID(clientID))
	subscribers.Store(conn, subClient)
	subscription.GetManager().AddClient(subClient)

	log.Printf("[websocket-reader] Client connected: id=%s, address=%s", clientID, r.RemoteAddr)

	defer func() {
		close(done) // Stop ping goroutine
		if err := conn.Close(); err != nil {
			log.Printf("[websocket-reader] Error closing connection for client %s: %v", clientID, err)
		}
		subscribers.Delete(conn)
		subscription.GetManager().RemoveClient(types.ClientID(clientID))
		log.Printf("[websocket-reader] Client disconnected: id=%s, address=%s", clientID, r.RemoteAddr)
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("[websocket-reader] Unexpected close for client %s: %v", clientID, err)
			} else if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Printf("[websocket-reader] Client %s closed connection normally", clientID)
			} else {
				log.Printf("[websocket-reader] Read error for client %s: %v", clientID, err)
			}
			break
		}

		var clientMsg ClientMessage
		if err := json.Unmarshal(message, &clientMsg); err != nil {
			log.Printf("[websocket-reader] Invalid message format from client %s: %v\nMessage: %s", clientID, err, string(message))
			sendError(conn, "Invalid message format")
			continue
		}

		// log.Printf("[websocket-reader] Received message type '%s' from client %s", clientMsg.Type, clientID)

		switch clientMsg.Type {
		case TypeSubscribe:
			handleSubscribe(conn, subClient, clientMsg.Data)
		case TypeUnsubscribe:
			handleUnsubscribe(conn, subClient, clientMsg.Data)
		default:
			log.Printf("[websocket-reader] Unhandled message type '%s' from client %s", clientMsg.Type, clientID)
			sendMessage(conn, "unhandled_message", nil)
		}
	}
}

func sendError(conn *websocket.Conn, errMsg string) {
	log.Printf("[websocket-reader] Sending error to client: %s", errMsg)
	sendMessage(conn, "error", errMsg)
}

func sendMessage(conn *websocket.Conn, msgType string, data interface{}) {
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	msg := map[string]interface{}{
		"type": msgType,
		"data": data,
	}
	if err := conn.WriteJSON(msg); err != nil {
		log.Printf("[websocket-reader] Error sending message type '%s': %v", msgType, err)
	}
}