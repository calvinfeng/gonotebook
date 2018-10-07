package chat

import (
	"encoding/json"
	"fmt"
)

// Payload is the expected JSON format for a websocket payload.
type Payload struct {
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

// Demultiplexer takes an input and fans it out to all consumers.
type Demultiplexer struct {
	In             chan json.RawMessage
	ConsumerByID   map[string]*Consumer
	RoomByID       map[string][]string
	AddConsumer    chan *Consumer
	RemoveConsumer chan *Consumer
}

// Loop listens for all updates.
func (d *Demultiplexer) Loop() {
	for {
		select {
		case rm := <-d.In:
			d.handleInput(rm)
		case c := <-d.AddConsumer:
			d.handleAddConsumer(c)
		case c := <-d.RemoveConsumer:
			d.handleRemoveConsumer(c)
		}
	}
}

func (d *Demultiplexer) handleInput(rm json.RawMessage) {
	p := Payload{}
	json.Unmarshal(rm, &p)
	loginfo(fmt.Sprintf("received message %s", p.Message))
}

func (d *Demultiplexer) handleAddConsumer(c *Consumer) {
	// Create a new room for consumer if the room does not exist.
}

func (d *Demultiplexer) handleRemoveConsumer(c *Consumer) {
	// Delete the room if it is empty.
}
