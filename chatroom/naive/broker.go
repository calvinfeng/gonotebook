package naive

import (
	"fmt"
	"go-academy/chatroom/util"
	"math/rand"
	"net/http"

	"github.com/gorilla/websocket"
)

var charRunes = []rune("0123456789abcdef")

// Payload is the expected JSON format for a websocket payload.
type Payload struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

// NewBroker returns an initialized broker.
func NewBroker() *MessageBroker {
	mb := &MessageBroker{
		ConnByID:   make(map[string]*websocket.Conn),
		Broadcast:  make(chan Payload),
		AddConn:    make(chan Registration),
		RemoveConn: make(chan Registration),
	}

	return mb
}

// MessageBroker carries the responsibility of broadcasting.
type MessageBroker struct {
	ConnByID   map[string]*websocket.Conn
	AddConn    chan Registration
	RemoveConn chan Registration
	Broadcast  chan Payload
}

// Registration is a payload that is sent to broker for registration of a connection.
type Registration struct {
	ID   string
	Conn *websocket.Conn
}

// RandStringID returns a random string of size n which is composed of digits from 0-9.
func RandStringID(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = charRunes[rand.Intn(len(charRunes))]
	}

	return string(b)
}

// ListenBroadcast listens to broadcast and write it to every client.
func (mb *MessageBroker) ListenBroadcast() {
	for {
		select {
		case payload := <-mb.Broadcast:
			for ID := range mb.ConnByID {
				err := mb.ConnByID[ID].WriteJSON(payload)
				if err != nil {
					util.LogErr("WriteJSON", err)
				}
			}
		case r := <-mb.AddConn:
			mb.ConnByID[r.ID] = r.Conn
			util.LogInfo(fmt.Sprintf("client:%s has joined chatroom", r.ID))

		case r := <-mb.RemoveConn:
			delete(mb.ConnByID, r.ID)
			util.LogInfo(fmt.Sprintf("client:%s has left chatroom", r.ID))
		}
	}
}

// NewMessageStreamHandler returns a streams endpoint handler.
func NewMessageStreamHandler(u *websocket.Upgrader, mb *MessageBroker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := u.Upgrade(w, r, nil)
		if err != nil {
			util.LogErr("Upgrade", err)
			return
		}

		defer conn.Close()
		rg := Registration{Conn: conn, ID: RandStringID(15)}
		mb.AddConn <- rg

		for {
			p := Payload{}
			err := conn.ReadJSON(&p)
			switch {
			case websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway):
				mb.RemoveConn <- rg
				return
			case err != nil:
				util.LogErr("ReadJSON", err)
			default:
				util.LogInfo(fmt.Sprintf("broadcasting message: %s", p.Message))
				mb.Broadcast <- p
			}
		}
	}
}
