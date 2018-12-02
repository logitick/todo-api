package adding

// Service for adding Todo items
type Service interface {
	// AddTodo returns the identifier for the added item
	AddTodo(Todo) (int, error)
}

// Repository for storing data items
type Repository interface {
	// AddTodo saves the Todo item to the datastore and returns the item
	AddTodo(Todo) (int, error)
}

type service struct {
	r Repository
}

// NewService returns a new instance of the default adder Service
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddTodo(t Todo) (int, error) {
	return s.r.AddTodo(t)
}
