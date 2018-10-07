package stream

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

// WebsocketConsumer reads and writes to broker.
type WebsocketConsumer struct {
	conn  *websocket.Conn
	read  chan json.RawMessage
	write chan json.RawMessage
}

// ReadPump pummps from read channel and perform some logic.
func (wc *WebsocketConsumer) ReadPump(ctx context.Context) {
	wc.conn.SetReadDeadline(time.Now().Add(time.Minute))
	wc.conn.SetPongHandler(func(s string) error {
		wc.conn.SetReadDeadline(time.Now().Add(time.Minute))
		return nil
	})

	for {
		select {
		case <-ctx.Done():
			return
		case raw := <-wc.read:
			wc.conn.SetReadDeadline(time.Now().Add(time.Minute))
			fmt.Println(string(raw))
		}
	}
}

// WritePump pumps from write channel and write to websocket connection.
func (wc *WebsocketConsumer) WritePump(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case raw := <-wc.write:
			wc.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			err := wc.conn.WriteMessage(websocket.TextMessage, raw)
			if err != nil {
				return
			}
		case <-ticker.C:
			wc.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			err := wc.conn.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				return
			}
		}
	}
}
