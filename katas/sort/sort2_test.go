package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQSort(t *testing.T) {
	assert.Equal(t, []int{}, QSort([]int{}))
	assert.Equal(t, []int{1}, QSort([]int{1}))
	assert.Equal(t, []int{1, 2}, QSort([]int{1, 2}))
	assert.Equal(t, []int{1, 2}, QSort([]int{2, 1}))
	assert.Equal(t, []int{1, 2, 3}, QSort([]int{1, 2, 3}))
	assert.Equal(t, []int{1, 2, 3}, QSort([]int{2, 1, 3}))
	assert.Equal(t, []int{1, 2, 3}, QSort([]int{3, 2, 1}))
	assert.Equal(t, []int{1, 2, 3, 4}, QSort([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{1, 2, 3, 4}, QSort([]int{2, 1, 3, 4}))
	assert.Equal(t, []int{1, 2, 3, 4}, QSort([]int{4, 3, 2, 1}))
	assert.Equal(t,
		[]int{1, 2, 2, 3, 3, 4, 4, 5, 7, 8, 8, 9},
		QSort([]int{8, 9, 4, 2, 4, 3, 1, 8, 7, 5, 3, 2}))
}
