package memory

import (
	"fmt"

	"github.com/logitick/todo-api/pkg/adding"
	"github.com/logitick/todo-api/pkg/listing"
	"github.com/logitick/todo-api/pkg/updating"
)

type Storage struct {
	todo []Todo
}

func (s *Storage) GetTodos() []listing.Todo {
	todos := make([]listing.Todo, len(s.todo))
	for i, t := range s.todo {
		todos[i] = listing.Todo{
			ID:          t.ID,
			Title:       t.Title,
			CompletedAt: t.CompletedAt,
		}
	}
	return todos // ♪ todos hermanos ♪
}

func (s *Storage) GetTodo(i int) (listing.Todo, error) {
	var td listing.Todo
	for _, t := range s.todo {
		if t.ID == i {
			td = listing.Todo{
				ID:          t.ID,
				Title:       t.Title,
				CompletedAt: t.CompletedAt,
			}
			return td, nil
		}
	}
	return td, fmt.Errorf("Cannot find todo with id %v", i)
}

func (s *Storage) AddTodo(t adding.Todo) (int, error) {
	id := len(s.todo) + 1
	tt := Todo{
		ID:    id,
		Title: t.Title,
	}
	s.todo = append(s.todo, tt)
	return id, nil
}

func (s *Storage) UpdateTodo(t updating.Todo) (updating.Todo, error) {
	for i, st := range s.todo {
		if t.ID == st.ID {
			st.CompletedAt = t.CompletedAt
			st.Title = t.Title
			s.todo[i] = st
			return t, nil
		}
	}
	return t, fmt.Errorf("Not found: %v", t)
}
