package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

func main() {
	fs := http.FileServer(http.Dir("../../static"))

	s := &Server{
		ConnMap: make(map[*websocket.Conn]bool),
		Broadcast: make(chan Payload),
		Upgrader: &websocket.Upgrader{},
	}

	http.Handle("/", fs)
	http.HandleFunc("/ws", s.getHandleConnections())

	fmt.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
