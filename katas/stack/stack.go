package stack

import "errors"

type stack struct {
	size    int
	element [10]int
}

func NewStack() *stack {
	return &stack{}
}

func (s *stack) IsEmpty() bool {
	return s.size == 0
}

func (s *stack) Push(element int) {
	s.element[s.size] = element
	s.size++
}

func (s *stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("underflow error")
	}
	s.size--
	return s.element[s.size], nil
}

func (s *stack) Size() int {
	return s.size
}
