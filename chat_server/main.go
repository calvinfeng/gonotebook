package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))

	b := &Broker{
		ConnMap:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan Payload),
		Upgrader:  &websocket.Upgrader{},
	}

	go b.handleBroadcast()

	http.Handle("/", fs)
	http.HandleFunc("/ws", b.getHandleConnections())

	fmt.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
