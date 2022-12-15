package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRangeIterator(t *testing.T) {

	begin := -100
	end := 100
	iterator := NewRangeIterator(begin, end)
	slice := ToSlice[int](iterator)

	exceptedSlice := make([]int, 0)
	for begin < end {
		exceptedSlice = append(exceptedSlice, begin)
		begin++
	}

	assert.Equal(t, exceptedSlice, slice)

}
