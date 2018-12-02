package memory

import (
	"fmt"

	"github.com/logitick/todo-api/pkg/adding"
	"github.com/logitick/todo-api/pkg/listing"
)

type Storage struct {
	todo []Todo
}

func (s *Storage) GetTodos() []listing.Todo {
	// for i := 0; i < 100; i++ {
	// 	s.todo = append(s.todo, Todo{
	// 		ID:          i,
	// 		Title:       fake.Sentence(),
	// 		CompletedAt: time.Now(),
	// 	})
	// }
	todos := make([]listing.Todo, len(s.todo))
	for _, t := range s.todo {
		todos = append(
			todos,
			listing.Todo{
				ID:          t.ID,
				Title:       t.Title,
				CompletedAt: t.CompletedAt,
			},
		)
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
