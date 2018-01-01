package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type Payload struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type Broker struct {
	ConnMap   map[*websocket.Conn]bool
	Broadcast chan Payload
	Upgrader  *websocket.Upgrader
}

func (b *Broker) handleBroadcast() {
	for payload := range b.Broadcast {
		for conn := range b.ConnMap {
			err := conn.WriteJSON(payload)
			if err != nil {
				fmt.Println("Failed to write JSON to websocket connection.")
				conn.Close()
				delete(b.ConnMap, conn)
			}
		}
	}
}

func (b *Broker) getHandleConnections() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := b.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Failed to upgrade GET to a websocket connection.")
		}

		defer conn.Close()

		b.ConnMap[conn] = true
		for {
			var p Payload

			if err := conn.ReadJSON(&p); err != nil {
				if websocket.IsUnexpectedCloseError(err) {
					fmt.Println("Client connection has closed")
				} else {
					fmt.Println("Encountered websocket error")
				}
				delete(b.ConnMap, conn)
				return
			} else {
				fmt.Println("Incoming message:", p.Message)
				b.Broadcast <- p
			}
		}
	}
}
