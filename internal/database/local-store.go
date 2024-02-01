package database

import (
	"fmt"
	"math/rand"
)

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

func (s *Store) GetTodos() []Todo {
	fmt.Println(s.todos)
	return s.todos
}

func (s *Store) AddTodo(t Todo) []Todo {
	id := rand.Intn(10000001)

	s.todos = append(s.todos, Todo{
		ID:          id,
		Title:       t.Title,
		Description: t.Title,
		Completed:   t.Completed,
	})

	return s.todos
}

func (s *Store) DeleteTodo(ID int) []Todo {
	indexToDelete := -1

	for i, todo := range s.todos {
		if todo.ID == ID {
			indexToDelete = i
		}
	}

	if indexToDelete != -1 {
		s.todos = append(s.todos[:indexToDelete], s.todos[indexToDelete+1:]...)
	}

	fmt.Println(s.todos)
	return s.todos
}
