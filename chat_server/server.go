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

type Server struct {
	Conns     map[*websocket.Conn]bool
	Broadcast chan Message
	Upgrader  *websocket.Upgrader
}

func (s *Server) handleBroadcast() {
	for payload := range s.Broadcast {
		for conn := range s.Conns {
			err := conn.WriteJSON(payload)
			if err != nil {
				fmt.Println("Failed to write JSON to websocket connection.")
				conn.Close()
				delete(s.Conns, conn)
			}
		}
	}
}

func (s *Server) getHandleConnections() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := s.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Failed to upgrade GET to a websocket connection.")
		}

		defer conn.Close()

		s.Conns[conn] = true
		for {
			var p Payload

			if err := conn.ReadJSON(&p); err != nil {
				fmt.Println("Encountered websocket error")
				delete(s.Conns, conn)
				return
			} else {
				fmt.Println("Incoming message:", p.Message)
				s.Broadcast <- p
			}
		}
	}
}
