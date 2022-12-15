package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSliceIterator(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	iterator := NewSliceIterator(slice)
	toSlice := ToSlice[int](iterator)
	assert.Equal(t, slice, toSlice)
}
