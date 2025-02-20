package txpool

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type RPCClient struct {
	conn      *websocket.Conn
	mu        sync.Mutex // Protects reqID
	writeMu   sync.Mutex // Protects websocket writes
	reqID     uint64
	pending   map[uint64]chan json.RawMessage
	pendingMu sync.Mutex
}

// NewRPCClient dials the given WebSocket URL and returns an RPCClient.
func NewRPCClient(url string) (*RPCClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	client := &RPCClient{
		conn:    conn,
		pending: make(map[uint64]chan json.RawMessage),
	}
	go client.readLoop()
	return client, nil
}

// readLoop continuously reads messages from the WebSocket and dispatches responses.
func (c *RPCClient) readLoop() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("read error: %v", err)
			return
		}
		var resp struct {
			ID     uint64          `json:"id"`
			Result json.RawMessage `json:"result"`
			Error  *struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			} `json:"error"`
		}
		if err := json.Unmarshal(message, &resp); err != nil {
			log.Printf("failed to unmarshal response: %v", err)
			continue
		}
		if ch, ok := c.pending[resp.ID]; ok {
			if resp.Error != nil {
				log.Fatalf("RPC Error: code=%d, message=%s", resp.Error.Code, resp.Error.Message)
			}
			ch <- resp.Result
			delete(c.pending, resp.ID)
		}
	}
}

// Call sends a JSON‑RPC request and waits for the response.
func (c *RPCClient) Call(method string, params interface{}) (json.RawMessage, error) {
	// Increment request ID.
	c.reqID++
	reqID := c.reqID

	req := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      reqID,
	}

	ch := make(chan json.RawMessage)
	c.pending[reqID] = ch

	// Protect the write with writeMu.
	err := c.conn.WriteJSON(req)
	if err != nil {
		return nil, err
	}
	// In production you might want to add a timeout here.
	result := <-ch
	return result, nil
}

func (c *RPCClient) CallLocked(method string, params interface{}) (json.RawMessage, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Call(method, params)
}

// Send sends a JSON‑RPC request without waiting for a response.
func (c *RPCClient) Send(method string, params interface{}) error {
	// Increment request ID.
	c.reqID++
	reqID := c.reqID

	req := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      reqID,
	}
	err := c.conn.WriteJSON(req)
	return err
} 