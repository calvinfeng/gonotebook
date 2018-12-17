package smart

import (
	"context"
	"encoding/json"
	"fmt"
	"go-academy/chatroom/util"
	"time"

	"github.com/gorilla/websocket"
)

// Client can read and write to demultiplexer.
type Client interface {
	ID() string
	Listen()
	Activate(ctx context.Context)
	WriteQueue() chan<- Payload
	SetBroadcast(chan Payload)
}

// newWebSocketClient returns a consumer with channels initialized.
func newWebSocketClient(c *websocket.Conn) Client {
	return &WebSocketClient{
		id:    util.RandStringID(15),
		conn:  c,
		read:  make(chan json.RawMessage, 100),
		write: make(chan Payload, 100),
	}
}

// WebSocketClient reads and writes to broker.
type WebSocketClient struct {
	id        string
	conn      *websocket.Conn
	read      chan json.RawMessage // acting as a read queue
	write     chan Payload         // acting as a write queue
	broadcast chan Payload
}

// ID is a getter for client's ID.
func (c *WebSocketClient) ID() string {
	return c.id
}

// Activate sets client loop(s) running.
func (c *WebSocketClient) Activate(ctx context.Context) {
	go c.readLoop(ctx)
	go c.writeLoop(ctx)
}

// Listen runs a loop to listen to websocket connection.
func (c *WebSocketClient) Listen() {
	for {
		_, bytes, err := c.conn.ReadMessage()
		if err != nil &&
			websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
			return
		}

		// This is an abnormal error, important to log it for debugging.
		if err != nil {
			util.LogErr("ReadJSON", err)
			return
		}

		c.read <- bytes
	}
}

// SetBroadcast is a setter for broadcast.
func (c *WebSocketClient) SetBroadcast(ch chan Payload) {
	c.broadcast = ch
}

// WriteQueue is a getter for write channel.
func (c *WebSocketClient) WriteQueue() chan<- Payload {
	return c.write
}

// readLoop pumps from read queue and process the message by unmarshaling it into struct.
func (c *WebSocketClient) readLoop(ctx context.Context) {
	c.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	c.conn.SetPongHandler(func(s string) error {
		c.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		return nil
	})

	for {
		select {
		case <-ctx.Done():
			util.LogInfo(fmt.Sprintf("client %s has termianted read loop", c.id))
			return
		case raw := <-c.read:
			c.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
			p := Payload{}

			err := json.Unmarshal(raw, &p)
			if err != nil {
				util.LogErr("JSON Unmarshal", err)
				continue
			}

			c.broadcast <- p
		}
	}
}

// writeLoop pumps from write channel and processes message by marshaling it into bytes and writes
// bytes to websocket connection.
func (c *WebSocketClient) writeLoop(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			util.LogInfo(fmt.Sprintf("client %s has terminated write loop", c.id))
			return

		case p := <-c.write:
			c.conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
			bytes, err := json.Marshal(p)
			if err != nil {
				util.LogErr("JSON Marshal", err)
				continue
			}

			err = c.conn.WriteMessage(websocket.TextMessage, bytes)
			if err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
			err := c.conn.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				return
			}
		}
	}
}
