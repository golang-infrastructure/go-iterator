package iterator

import (
	"testing"
)

func TestNewCycleIterator(t *testing.T) {

	iterator := NewCycleIterator(1, 2, 3)
	for i := 0; i < 100 && iterator.Next(); i++ {
		t.Log(iterator.Value())
	}

}
