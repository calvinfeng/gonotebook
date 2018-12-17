package main

import (
	"go-academy/chatroom/naive"
	"go-academy/chatroom/util"
	"net/http"

	"github.com/gorilla/websocket"
)

func naiveSetup() error {
	naive.RunBroker()

	streamHandler, err := naive.NewMessageStreamHandler(&websocket.Upgrader{})
	if err != nil {
		return err
	}

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.Handle("/streams/messages", streamHandler)

	return nil
}

func main() {
	naiveSetup()
	util.LogInfo("starting server on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		util.LogErr("ListenAndServe", err)
	}
}
