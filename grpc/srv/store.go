package srv

import (
	"errors"
	"sync"
)

var store = &InMemStorage{
	mutex: &sync.Mutex{},
	todos: make(map[int64]*TodoModel),
}

type (
	TodoModel struct {
		ID        int64  `json:"id"`
		UserID    int64  `json:"user_id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	InMemStorage struct {
		mutex *sync.Mutex
		todos map[int64]*TodoModel
	}
)

func (s *InMemStorage) Get(id int64) (*TodoModel, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if t, ok := s.todos[id]; ok {
		return t, nil
	}

	return nil, errors.New("not found")
}

func (s *InMemStorage) Set(t *TodoModel) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if t.Title == "" {
		return errors.New("title is required")
	}

	if _, ok := s.todos[t.ID]; !ok {
		t.ID = int64(len(s.todos) + 1)
	}

	s.todos[t.ID] = t
	return nil
}
