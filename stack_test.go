package tdd_kata

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

func (suite *StackTestSuite) SetupTest() {
	suite.stack = NewStack()
}

func (suite *StackTestSuite) TestCanCreateStack() {
	suite.True(suite.stack.IsEmpty())
}

func (suite *StackTestSuite) TestCamPush() {
	suite.stack.Push(1)
	suite.Equal(1, suite.stack.Size())
	suite.False(suite.stack.IsEmpty())
}

func (suite *StackTestSuite) TestAfterOnePushAndPop_stackShouldEmpty() {
	suite.stack.Push(1)
	suite.stack.Pop()
	suite.True(suite.stack.IsEmpty())
	suite.Equal(0, suite.stack.Size())

}

func (suite *StackTestSuite) TestAfterPushingTwoElement_sizeShouldBeTwo() {
	suite.stack.Push(1)
	suite.stack.Push(2)
	suite.Equal(2, suite.stack.Size())
}

func (suite *StackTestSuite) TestPushedX_returnsX() {
	suite.stack.Push(100)
	value1, _ := suite.stack.Pop()
	suite.Equal(100, value1)
	suite.stack.Push(200)
	value2, _ := suite.stack.Pop()
	suite.Equal(200, value2)
}

func (suite *StackTestSuite) TestPushX_Y_thenReturnY_X() {
	suite.stack.Push(100)
	suite.stack.Push(200)
	value1, _ := suite.stack.Pop()
	value2, _ := suite.stack.Pop()
	suite.Equal(200, value1)
	suite.Equal(100, value2)
}

func (suite *StackTestSuite) TestPoppingOnEmptyStack_returnError() {
	_, err := suite.stack.Pop()
	suite.EqualError(err, "underflow error")
}
