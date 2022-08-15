package tdd_kata

import (
	"github.com/stretchr/testify/assert"
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

func (s *StackTestSuite) TestCanCreat() {
	assert.True(s.T(), s.stack.IsEmpty())
}

func (s *StackTestSuite) TestPushOneValue_stackIsNotEmpty() {
	s.stack.Push(1)
	assert.False(s.T(), s.stack.IsEmpty())
	assert.Equal(s.T(), 1, s.stack.Size())

}

func (s *StackTestSuite) TestPushTwoValue_sizeShouldBeTow() {
	s.stack.Push(100)
	s.stack.Push(200)
	assert.Equal(s.T(), 2, s.stack.Size())
}
