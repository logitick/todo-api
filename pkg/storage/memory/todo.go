package memory

import "time"

// The Todo item to be listed
type Todo struct {
	ID          int
	Title       string
	CompletedAt time.Time
}
