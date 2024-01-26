package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	is := IsSquare(43)
	assert.Equal(t, false, is)
	is2 := IsSquare(4)
	assert.Equal(t, true, is2)
	res := Filter([]int{2, 7, 4, 49, 32, 100}, IsSquare)
	assert.EqualValues(t, []int{4, 49, 100}, res)
	ab := Abs(-5)
	assert.Equal(t, 5, ab)
	res2 := Map([]int{2, -8, 89, 5, -100, 0}, Abs)
	assert.EqualValues(t, []int{2, 8, 89, 5, 100, 0}, res2)
	res3 := Filter([]int{5, 42, 131, 14224, -9009, 9009}, IsPalindrome)
	assert.EqualValues(t, []int{5, 131, -9009, 9009}, res3)
}
