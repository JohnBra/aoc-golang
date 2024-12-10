package datastructures

import "fmt"

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: []T{}}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var item T
	if len(s.items) == 0 {
		return item, false
	}
	item = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item, true
}

func (s *Stack[T]) Top() (T, bool) {
	var item T

	if len(s.items) == 0 {
		return item, false
	}

	item = s.items[len(s.items)-1]

	return item, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", s.items)
}
