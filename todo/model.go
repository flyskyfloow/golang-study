package todo

import "sync"

// Todo represents a single todo item.
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Store is an in-memory thread-safe store for Todo items.
type Store struct {
	mu      sync.RWMutex
	todos   []Todo
	counter int
}

// NewStore creates and returns an empty Store.
func NewStore() *Store {
	return &Store{}
}

// Add inserts a new Todo and returns it with the assigned ID.
func (s *Store) Add(title string) Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counter++
	t := Todo{ID: s.counter, Title: title, Completed: false}
	s.todos = append(s.todos, t)
	return t
}

// List returns a copy of all stored Todo items.
func (s *Store) List() []Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]Todo, len(s.todos))
	copy(result, s.todos)
	return result
}

// Update modifies the title and completed status of the Todo with the given id.
// It returns the updated Todo and true on success, or the zero value and false when not found.
func (s *Store) Update(id int, title string, completed bool) (Todo, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, t := range s.todos {
		if t.ID == id {
			s.todos[i].Title = title
			s.todos[i].Completed = completed
			return s.todos[i], true
		}
	}
	return Todo{}, false
}

// Delete removes the Todo with the given id.
// It returns true on success and false when not found.
func (s *Store) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, t := range s.todos {
		if t.ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return true
		}
	}
	return false
}
