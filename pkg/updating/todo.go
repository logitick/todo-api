package updating

import "time"

// The Todo item to be listed
type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	CompletedAt time.Time `json:"completedAt"`
}
