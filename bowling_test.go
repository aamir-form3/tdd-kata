package tdd_kata

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BowlingTestSuite struct {
	suite.Suite
	g *game
}

func (suite *BowlingTestSuite) rollSpare() {
	suite.g.Roll(5)
	suite.g.Roll(5)
}

func (suite *BowlingTestSuite) rollMany(pins int, times int) {
	for i := 0; i < times; i++ {
		suite.g.Roll(pins)
	}
}

func TestBowlingSuite(t *testing.T) {
	suite.Run(t, new(BowlingTestSuite))
}

func (suite *BowlingTestSuite) SetupTest() {
	suite.g = NewGame()
}

func (suite *BowlingTestSuite) TestCanRoll() {
	suite.g.Roll(0)
	assert.Equal(suite.T(), 0, suite.g.Score())
}

func (suite *BowlingTestSuite) TestGutterGame() {
	suite.rollMany(0, 20)
	assert.Equal(suite.T(), 0, suite.g.Score())
}

func (suite *BowlingTestSuite) TestAllOneGame() {
	suite.rollMany(1, 20)
	assert.Equal(suite.T(), 20, suite.g.Score())
}

func (suite *BowlingTestSuite) TestOneSpare() {
	suite.rollSpare()
	suite.g.Roll(3)
	suite.rollMany(0, 17)
	assert.Equal(suite.T(), 16, suite.g.Score())
}

func (suite *BowlingTestSuite) TestOneStrike() {
	suite.g.Roll(10)
	suite.g.Roll(3)
	suite.g.Roll(4)
	suite.rollMany(0, 16)
	assert.Equal(suite.T(), 24, suite.g.Score())
}

func (suite *BowlingTestSuite) TestPerfectGame() {
	suite.rollMany(10, 12)
	assert.Equal(suite.T(), 300, suite.g.Score())
}
