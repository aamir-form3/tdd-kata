package stack

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type StackTestSuite struct {
	suite.Suite
	stack *stack
}

func TestStackSuite(t *testing.T) {
	suite.Run(t, new(StackTestSuite))
}

func (s *StackTestSuite) SetupTest() {
	s.stack = NewStack()
}

func (s *StackTestSuite) TestCanCreateStack() {
	s.True(s.stack.IsEmpty())
}

func (s *StackTestSuite) TestCamPush() {
	s.stack.Push(1)
	s.Equal(1, s.stack.Size())
	s.False(s.stack.IsEmpty())
}

func (s *StackTestSuite) TestAfterOnePushAndPop_stackShouldEmpty() {
	s.stack.Push(1)
	s.stack.Pop()
	s.True(s.stack.IsEmpty())
	s.Equal(0, s.stack.Size())

}

func (s *StackTestSuite) TestAfterPushingTwoElement_sizeShouldBeTwo() {
	s.stack.Push(1)
	s.stack.Push(2)
	s.Equal(2, s.stack.Size())
}

func (s *StackTestSuite) TestPushedX_returnsX() {
	s.stack.Push(100)
	value1, _ := s.stack.Pop()
	s.Equal(100, value1)
	s.stack.Push(200)
	value2, _ := s.stack.Pop()
	s.Equal(200, value2)
}

func (s *StackTestSuite) TestPushX_Y_thenReturnY_X() {
	s.stack.Push(100)
	s.stack.Push(200)
	value1, _ := s.stack.Pop()
	value2, _ := s.stack.Pop()
	s.Equal(200, value1)
	s.Equal(100, value2)
}

func (s *StackTestSuite) TestPoppingOnEmptyStack_returnError() {
	_, err := s.stack.Pop()
	s.EqualError(err, "underflow error")
}
