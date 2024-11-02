package collections

import "sync"

type Stack[T any] struct {
	items []T
	mu    sync.RWMutex
}

func (s *Stack[T]) Push(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.items == nil {
		s.items = []T{}
	}

	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	s.mu.Lock()
	defer s.mu.Unlock()

	length := len(s.items)
	if length == 0 {
		var zero T
		return zero
	}

	item := s.items[length-1]
	s.items = s.items[0 : length-1]
	return item
}

func (s *Stack[T]) Length() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.items)
}
