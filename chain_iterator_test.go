package iterator

import (
	"testing"
)

func TestNewChainIterator(t *testing.T) {

	iteratorA := NewRangeIterator(0, 5)
	iteratorB := NewRangeIterator(5, 10)

	iterator := NewChainIterator[int](iteratorA, iteratorB)
	slice := ToSlice[int](iterator)
	t.Log(slice)

}
