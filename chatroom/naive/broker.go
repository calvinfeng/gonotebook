package naive

import (
	"go-academy/chatroom/util"

	"github.com/gorilla/websocket"
)

// Payload is the expected JSON format for a websocket payload.
type Payload struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var broker *MessageBroker

// RunBroker configures a broker and runs it.
func RunBroker() {
	broker = &MessageBroker{
		connByID:   make(map[string]*websocket.Conn),
		broadcast:  make(chan Payload),
		addConn:    make(chan registration),
		removeConn: make(chan registration),
	}

	go broker.loop()
}

// MessageBroker carries the responsibility of broadcasting.
type MessageBroker struct {
	connByID   map[string]*websocket.Conn
	addConn    chan registration
	removeConn chan registration
	broadcast  chan Payload
}

// registration is a payload that is sent to broker for registering a new connection.
type registration struct {
	id   string
	conn *websocket.Conn
}

// ListenBroadcast listens to broadcast and write it to every client.
func (mb *MessageBroker) loop() {
	for {
		select {
		case payload := <-mb.broadcast:
			for ID := range mb.connByID {
				err := mb.connByID[ID].WriteJSON(payload) // This is bad practice, will explain why.
				if err != nil {
					util.LogErr("WriteJSON", err)
				}
			}
		case r := <-mb.addConn:
			mb.connByID[r.id] = r.conn

		case r := <-mb.removeConn:
			delete(mb.connByID, r.id)
		}
	}
}
