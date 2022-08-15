package tdd_kata

type stack struct {
	size int
}

func NewStack() *stack {
	return &stack{}
}

func (s *stack) IsEmpty() bool {
	return s.size == 0
}

func (s *stack) Push(value int) {
	s.size++
}

func (s *stack) Size() int {
	return s.size
}
