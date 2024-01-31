package database

type Todo struct {
	Title       string
	Description string
	Completed   bool
	ID          int
}

type Store struct {
	count int
	todos []Todo
}

func New() *Store {
	localTodos := []Todo{
		{
			ID:          1,
			Title:       "Testing 1",
			Description: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
			Completed:   false,
		},
		{
			ID:          2,
			Title:       "Testing 2",
			Description: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum",
			Completed:   false,
		},
	}

	return &Store{
		count: 0,
		todos: localTodos,
	}
}

func (s *Store) Increment() int {
	s.count = s.count + 1
	return s.count
}

func (s *Store) CurrentCount() int {
	return s.count
}
