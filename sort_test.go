package tdd_kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	assert.Equal(t, []int{}, Sort([]int{}))
	assert.Equal(t, []int{1}, Sort([]int{1}))
	assert.Equal(t, []int{1, 2}, Sort([]int{1, 2}))
	assert.Equal(t, []int{1, 2}, Sort([]int{2, 1}))
	assert.Equal(t, []int{1, 2, 3}, Sort([]int{1, 2, 3}))
	assert.Equal(t, []int{1, 2, 3}, Sort([]int{2, 1, 3}))
	assert.Equal(t, []int{1, 2, 3}, Sort([]int{3, 1, 2}))
	assert.Equal(t, []int{1, 2, 3}, Sort([]int{3, 2, 1}))

}
