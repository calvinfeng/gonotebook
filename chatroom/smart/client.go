package smart

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

// Client can read and write to demultiplexer.
type Client interface {
	ID() string
	Activate(ctx context.Context)
	Destroy()
}

// NewWebSocketClient returns a consumer with channels initialized.
func NewWebSocketClient(c *websocket.Conn) *WebSocketClient {
	return &WebSocketClient{
		conn:  c,
		read:  make(chan json.RawMessage),
		write: make(chan json.RawMessage),
	}
}

// WebSocketClient reads and writes to broker.
type WebSocketClient struct {
	conn  *websocket.Conn
	read  chan json.RawMessage
	write chan json.RawMessage
}

// ReadPump pummps from read channel and processes the message..
func (c *WebSocketClient) ReadPump(ctx context.Context) {
	c.conn.SetReadDeadline(time.Now().Add(time.Minute))
	c.conn.SetPongHandler(func(s string) error {
		c.conn.SetReadDeadline(time.Now().Add(time.Minute))
		return nil
	})

	for {
		select {
		case <-ctx.Done():
			return
		case raw := <-c.read:
			c.conn.SetReadDeadline(time.Now().Add(time.Minute))
			fmt.Println(string(raw))
		}
	}
}

// WritePump pumps from write channel and write to websocket connection.
func (c *WebSocketClient) WritePump(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case raw := <-c.write:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			err := c.conn.WriteMessage(websocket.TextMessage, raw)
			if err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			err := c.conn.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				return
			}
		}
	}
}
