package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

// NewConsumer returns a consumer with channels initialized.
func NewConsumer(c *websocket.Conn) *Consumer {
	return &Consumer{
		conn:  c,
		read:  make(chan json.RawMessage),
		write: make(chan json.RawMessage),
	}
}

// Consumer reads and writes to broker.
type Consumer struct {
	conn  *websocket.Conn
	read  chan json.RawMessage
	write chan json.RawMessage
}

// ReadPump pummps from read channel and perform some logic.
func (c *Consumer) ReadPump(ctx context.Context) {
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
func (c *Consumer) WritePump(ctx context.Context) {
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
