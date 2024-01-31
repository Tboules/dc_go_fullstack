package database

type Store struct {
	count int
}

func New() *Store {
	return &Store{
		count: 0,
	}
}

func (s *Store) Increment() int {
	s.count = s.count + 1
	return s.count
}

func (s *Store) CurrentCount() int {
	return s.count
}
