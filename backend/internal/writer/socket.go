package writer

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/oklog/ulid/v2"
)

// Add the upgrader as a package-level variable that can be set from server.go
var Upgrader *websocket.Upgrader

// Message types
type MessageType string

const (
	TypeSetValue    MessageType = "set_value"
	TypeEmitEvent   MessageType = "emit_event"
	TypeError       MessageType = "error"
)

type ClientMessage struct {
	Type MessageType     `json:"type"`
	Data json.RawMessage `json:"data"`
}

// HandleWebSocket handles new WebSocket connections
func handleWriteWebSocket(conn *websocket.Conn, r *http.Request) {
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
					log.Printf("[websocket-writer] Ping failed: %v", err)
					close(done)
					return
				}
			case <-done:
				return
			}
		}
	}()

	clientID := ulid.Make().String()
	// Parse private key from URL if present
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Printf("[websocket-writer] Error parsing URL for client %s: %v", clientID, err)
		return
	}
	privateKey := u.Query().Get("privateKey")

	log.Printf("[websocket-writer] Client connected: id=%s, address=%s", clientID, r.RemoteAddr)

	defer func() {
		close(done) // Stop ping goroutine
		if err := conn.Close(); err != nil {
			log.Printf("[websocket-writer] Error closing connection for client %s: %v", clientID, err)
		}
		log.Printf("[websocket-writer] Client disconnected: id=%s, address=%s", clientID, r.RemoteAddr)
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("[websocket-writer] Unexpected close for client %s: %v", clientID, err)
			} else if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Printf("[websocket-writer] Client %s closed connection normally", clientID)
			} else {
				log.Printf("[websocket-writer] Read error for client %s: %v", clientID, err)
			}
			break
		}

		var clientMsg ClientMessage
		if err := json.Unmarshal(message, &clientMsg); err != nil {
			log.Printf("[websocket-writer] Invalid message format from client %s: %v\nMessage: %s", clientID, err, string(message))
			sendError(conn, "Invalid message format")
			continue
		}

		// log.Printf("[websocket-writer] Received message type '%s' from client %s", clientMsg.Type, clientID)

		switch clientMsg.Type {
		case TypeSetValue:
			if privateKey == "" {
				log.Printf("[websocket-writer] Client %s attempted set_value without private key", clientID)
				sendError(conn, "No private key for set_value")
				continue
			}
			handleSetValue(conn, privateKey, clientMsg.Data)
		case TypeEmitEvent:
			if privateKey == "" {
				log.Printf("[websocket-writer] Client %s attempted emit_event without private key", clientID)
				sendError(conn, "No private key for emit_event")
				continue
			}
			handleEmitEvent(conn, privateKey, clientMsg.Data)
		default:
			log.Printf("[websocket-writer] Unhandled message type '%s' from client %s", clientMsg.Type, clientID)
			sendMessage(conn, "unhandled_message", nil)
		}
	}
}

func sendError(conn *websocket.Conn, errMsg string) {
	log.Printf("[websocket-writer] Sending error to client: %s", errMsg)
	sendMessage(conn, "error", errMsg)
}

func sendMessage(conn *websocket.Conn, msgType string, data interface{}) {
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	msg := map[string]interface{}{
		"type": msgType,
		"data": data,
	}
	if err := conn.WriteJSON(msg); err != nil {
		log.Printf("[websocket-writer] Error sending message type '%s': %v", msgType, err)
	}
}