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

func (s *BowlingTestSuite) rollSpare() {
	s.g.Roll(5)
	s.g.Roll(5)
}

func (s *BowlingTestSuite) rollMany(pins int, times int) {
	for i := 0; i < times; i++ {
		s.g.Roll(pins)
	}
}

func TestBowlingSuite(t *testing.T) {
	suite.Run(t, new(BowlingTestSuite))
}

func (s *BowlingTestSuite) SetupTest() {
	s.g = NewGame()
}

func (s *BowlingTestSuite) TestCanRoll() {
	s.g.Roll(0)
	assert.Equal(s.T(), 0, s.g.Score())
}

func (s *BowlingTestSuite) TestGutterGame() {
	s.rollMany(0, 20)
	assert.Equal(s.T(), 0, s.g.Score())
}

func (s *BowlingTestSuite) TestAllOneGame() {
	s.rollMany(1, 20)
	assert.Equal(s.T(), 20, s.g.Score())
}

func (s *BowlingTestSuite) TestOneSpare() {
	s.rollSpare()
	s.g.Roll(3)
	s.rollMany(0, 17)
	assert.Equal(s.T(), 16, s.g.Score())
}

func (s *BowlingTestSuite) TestOneStrike() {
	s.g.Roll(10)
	s.g.Roll(3)
	s.g.Roll(4)
	s.rollMany(0, 16)
	assert.Equal(s.T(), 24, s.g.Score())
}

func (s *BowlingTestSuite) TestPerfectGame() {
	s.rollMany(10, 12)
	assert.Equal(s.T(), 300, s.g.Score())
}
