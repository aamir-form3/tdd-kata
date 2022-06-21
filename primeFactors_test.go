package tdd_kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrimeFactor(t *testing.T) {
	assert.Empty(t, FactorOf(1))
	assert.Equal(t, FactorOf(2), []int{2})
	assert.Equal(t, FactorOf(3), []int{3})
	assert.Equal(t, FactorOf(4), []int{2, 2})
	assert.Equal(t, FactorOf(5), []int{5})
	assert.Equal(t, FactorOf(6), []int{2, 3})
	assert.Equal(t, FactorOf(7), []int{7})
	assert.Equal(t, FactorOf(8), []int{2, 2, 2})
	assert.Equal(t, FactorOf(9), []int{3, 3})
	assert.Equal(t, FactorOf(2*2*3*3*7*11*11*13), []int{2, 2, 3, 3, 7, 11, 11, 13})
}
