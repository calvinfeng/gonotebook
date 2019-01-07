package main

import (
	"encoding/json"
	"fmt"
)

func NewTodo(id int) *Todo {
	return &Todo{
		ID:        id,
		Networker: NewHTTPNetworker(),
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
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d/", t.ID)
	data, err := t.Get(url)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, t)
}
