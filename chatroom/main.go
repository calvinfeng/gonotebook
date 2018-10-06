package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	b := &Broker{
		ConnMap:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan Payload),
		Upgrader:  &websocket.Upgrader{},
	}

	go b.handleBroadcast()

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/ws", b.getHandleConnections())

	fmt.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
