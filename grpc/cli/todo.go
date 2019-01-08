package cli

import (
	"encoding/json"
)

func NewTodo(id int64) *Todo {
	return &Todo{
		ID: id,
	}
}

type Todo struct {
	ID        int64  `json:"id"`
	UserID    int    `json:"user_id"`
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

func (t *Todo) Save() error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	data, err = t.Set(t.ID, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, t)
}
