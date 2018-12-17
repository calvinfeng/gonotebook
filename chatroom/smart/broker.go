package smart

import (
	"encoding/json"
	"fmt"
	"go-academy/chatroom/util"
)

// Payload is the expected JSON format for a websocket payload.
type Payload struct {
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

// Broker takes an input and fans it out to all consumers.
type Broker struct {
	Broadcast    chan json.RawMessage
	ClientByID   map[string]Client
	RoomByID     map[string][]string
	AddClient    chan Client
	RemoveClient chan Client
}

// Loop listens for all updates.
func (b *Broker) Loop() {
	for {
		select {
		case rm := <-b.Broadcast:
			b.handleBroadcast(rm)
		case c := <-b.AddClient:
			b.handleAddClient(c)
		case c := <-b.RemoveClient:
			b.handleRemoveClient(c)
		}
	}
}

func (b *Broker) handleBroadcast(rm json.RawMessage) {
	p := Payload{}
	json.Unmarshal(rm, &p)
	util.LogInfo(fmt.Sprintf("received message %s", p.Message))
}

func (b *Broker) handleAddClient(c Client) {
	// Create a new room for consumer if the room does not exist.
}

func (b *Broker) handleRemoveClient(c Client) {
	// Delete the room if it is empty.
}
