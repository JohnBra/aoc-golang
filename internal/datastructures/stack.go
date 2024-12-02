package datastructures

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Clear() {
	s.items = nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}
