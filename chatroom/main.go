package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gorilla/websocket"
)

var charRunes = []rune("0123456789")

// Payload is the expected JSON format for a websocket payload.
type Payload struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
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
			for ID, c := range mb.ConnByID {
				err := c.WriteJSON(payload)
				if err != nil {
					fmt.Printf("failed to write JSON to client:%s, %s", ID, err)
				}
			}
		case r := <-mb.AddConn:
			mb.ConnByID[r.ID] = r.Conn
			fmt.Printf("client:%s has joined chatroom\n", r.ID)

		case r := <-mb.RemoveConn:
			delete(mb.ConnByID, r.ID)
			fmt.Printf("client:%s has left chatroom\n", r.ID)
		}
	}
}

// NewStreamHandler returns a streams endpoint handler.
func NewStreamHandler(u *websocket.Upgrader, mb *MessageBroker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := u.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("failed to upgrade connection", err)
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
				fmt.Println("expected error", err)
			default:
				fmt.Println("broadcasting message:", p.Message)
				mb.Broadcast <- p
			}
		}
	}
}

func main() {
	mb := &MessageBroker{
		ConnByID:   make(map[string]*websocket.Conn),
		Broadcast:  make(chan Payload),
		AddConn:    make(chan Registration),
		RemoveConn: make(chan Registration),
	}

	u := &websocket.Upgrader{}

	go mb.ListenBroadcast()

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/streams/", NewStreamHandler(u, mb))

	fmt.Println("starting server on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
