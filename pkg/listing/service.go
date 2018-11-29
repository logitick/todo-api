package listing

// Service interface contains operations
// needed for listing Todo items
type Service interface {
	GetTodos() []Todo
	GetTodo(int) (Todo, error)
}

type Repository interface {
	GetTodos() []Todo
	GetTodo(int) (Todo, error)
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetTodos() []Todo {
	return s.r.GetTodos()
}

func (s *service) GetTodo(id int) (Todo, error) {
	return s.r.GetTodo(id)
}
