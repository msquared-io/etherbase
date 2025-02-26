package txpool

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type RPCClient struct {
	url       string        // Store the URL for reconnection
	conn      *websocket.Conn
	mu        sync.Mutex    // Protects reqID
	writeMu   sync.Mutex    // Protects websocket writes
	connMu    sync.Mutex    // Protects connection state
	reqID     uint64
	pending   map[uint64]chan json.RawMessage
	pendingMu sync.Mutex
	closed    bool          // Flag to indicate if client is closed
	reconnect chan struct{} // Channel to trigger reconnection
}

// NewRPCClient dials the given WebSocket URL and returns an RPCClient.
func NewRPCClient(url string) (*RPCClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	client := &RPCClient{
		url:       url,
		conn:      conn,
		pending:   make(map[uint64]chan json.RawMessage),
		reconnect: make(chan struct{}, 1),
	}
	go client.readLoop()
	go client.reconnectLoop()
	return client, nil
}

// reconnectLoop handles reconnection attempts when the connection is lost
func (c *RPCClient) reconnectLoop() {
	for range c.reconnect {
		c.connMu.Lock()
		if c.closed {
			c.connMu.Unlock()
			return
		}
		
		// Close the existing connection if it's still open
		if c.conn != nil {
			c.conn.Close()
		}
		
		// Attempt to reconnect with exponential backoff
		maxRetries := 5
		backoff := 1 * time.Second
		
		var conn *websocket.Conn
		var err error
		
		for i := 0; i < maxRetries; i++ {
			log.Printf("Attempting to reconnect to WebSocket (attempt %d/%d)...", i+1, maxRetries)
			conn, _, err = websocket.DefaultDialer.Dial(c.url, nil)
			if err == nil {
				break
			}
			log.Printf("Reconnection failed: %v. Retrying in %v...", err, backoff)
			time.Sleep(backoff)
			backoff *= 2 // Exponential backoff
		}
		
		if err != nil {
			log.Printf("Failed to reconnect after %d attempts: %v", maxRetries, err)
			c.connMu.Unlock()
			// Try again later
			time.Sleep(10 * time.Second)
			c.triggerReconnect()
			continue
		}
		
		log.Printf("Successfully reconnected to WebSocket")
		c.conn = conn
		c.connMu.Unlock()
		
		// Start a new read loop for the new connection
		go c.readLoop()
	}
}

// triggerReconnect signals that a reconnection should be attempted
func (c *RPCClient) triggerReconnect() {
	select {
	case c.reconnect <- struct{}{}:
		// Signal sent
	default:
		// Channel already has a pending signal
	}
}

// readLoop continuously reads messages from the WebSocket and dispatches responses.
func (c *RPCClient) readLoop() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			c.triggerReconnect()
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
			log.Printf("Failed to unmarshal response: %v", err)
			continue
		}
		
		c.pendingMu.Lock()
		ch, ok := c.pending[resp.ID]
		if ok {
			if resp.Error != nil {
				log.Printf("RPC Error: code=%d, message=%s", resp.Error.Code, resp.Error.Message)
			}
			ch <- resp.Result
			delete(c.pending, resp.ID)
		}
		c.pendingMu.Unlock()
	}
}

// getConnection safely gets the current connection or returns an error if disconnected
func (c *RPCClient) getConnection() (*websocket.Conn, error) {
	c.connMu.Lock()
	defer c.connMu.Unlock()
	
	if c.closed {
		return nil, errors.New("client is closed")
	}
	
	if c.conn == nil {
		return nil, errors.New("not connected")
	}
	
	return c.conn, nil
}

// Call sends a JSON‑RPC request and waits for the response.
func (c *RPCClient) Call(method string, params interface{}) (json.RawMessage, error) {
	c.mu.Lock()
	c.reqID++
	reqID := c.reqID
	c.mu.Unlock()

	req := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      reqID,
	}

	ch := make(chan json.RawMessage, 1)
	c.pendingMu.Lock()
	c.pending[reqID] = ch
	c.pendingMu.Unlock()

	// Try to send with retries
	maxRetries := 3
	var err error
	
	for i := 0; i < maxRetries; i++ {
		conn, connErr := c.getConnection()
		if connErr != nil {
			// Wait for reconnection
			time.Sleep(time.Second)
			continue
		}
		
		c.writeMu.Lock()
		err = conn.WriteJSON(req)
		c.writeMu.Unlock()
		
		if err == nil {
			break
		}
		
		log.Printf("Error sending RPC request (attempt %d/%d): %v", i+1, maxRetries, err)
		c.triggerReconnect()
		time.Sleep(time.Second * time.Duration(i+1)) // Increasing delay between retries
	}
	
	if err != nil {
		c.pendingMu.Lock()
		delete(c.pending, reqID)
		c.pendingMu.Unlock()
		return nil, err
	}

	// Wait for response with timeout
	select {
	case result := <-ch:
		return result, nil
	case <-time.After(30 * time.Second):
		c.pendingMu.Lock()
		delete(c.pending, reqID)
		c.pendingMu.Unlock()
		return nil, errors.New("RPC request timed out")
	}
}

func (c *RPCClient) CallLocked(method string, params interface{}) (json.RawMessage, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Call(method, params)
}

// Send sends a JSON‑RPC request without waiting for a response.
func (c *RPCClient) Send(method string, params interface{}) error {
	c.mu.Lock()
	c.reqID++
	reqID := c.reqID
	c.mu.Unlock()

	req := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      reqID,
	}
	
	// Try to send with retries
	maxRetries := 3
	var err error
	
	for i := 0; i < maxRetries; i++ {
		conn, connErr := c.getConnection()
		if connErr != nil {
			// Wait for reconnection
			time.Sleep(time.Second)
			continue
		}
		
		c.writeMu.Lock()
		err = conn.WriteJSON(req)
		c.writeMu.Unlock()
		
		if err == nil {
			return nil
		}
		
		log.Printf("Error sending RPC request (attempt %d/%d): %v", i+1, maxRetries, err)
		c.triggerReconnect()
		time.Sleep(time.Second * time.Duration(i+1)) // Increasing delay between retries
	}
	
	return err
}

// Close closes the WebSocket connection and cleans up resources.
func (c *RPCClient) Close() error {
	c.connMu.Lock()
	defer c.connMu.Unlock()
	
	if c.closed {
		return nil
	}
	
	c.closed = true
	if c.conn != nil {
		return c.conn.Close()
	}
	
	return nil
} 