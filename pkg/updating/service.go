package updating

// Service functionality for updating Todo items
type Service interface {
	UpdateTodo(Todo) (Todo, error)
}

// Repository for storing data items
type Repository interface {
	UpdateTodo(Todo) (Todo, error)
}

type service struct {
	r Repository
}

// NewService returns a new Updating Service wired with the Repo
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateTodo(t Todo) (Todo, error) {
	return s.r.UpdateTodo(t)
}
