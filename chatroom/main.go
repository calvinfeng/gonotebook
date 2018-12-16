package main

import (
	"go-academy/chatroom/naive"
	"go-academy/chatroom/util"
	"net/http"

	"github.com/gorilla/websocket"
)

func naiveSetup() {
	mb := naive.NewBroker()
	u := &websocket.Upgrader{}
	go mb.ListenBroadcast()
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.Handle("/streams/messages", naive.NewMessageStreamHandler(u, mb))
}

func main() {
	naiveSetup()
	util.LogInfo("starting server on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		util.LogErr("ListenAndServe", err)
	}
}
