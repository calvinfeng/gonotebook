package cli

import (
	"encoding/json"
)

func NewTodo(id int) *Todo {
	return &Todo{
		ID:        id,
		Networker: NewHTTPNetworker("https://jsonplaceholder.typicode.com/todos"),
	}
}

type Todo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`

	Networker
}

func (t *Todo) Load() error {
	data, err := t.Get(t.ID)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, t)
}
