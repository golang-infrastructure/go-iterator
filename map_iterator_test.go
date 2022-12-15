package iterator

import (
	"testing"
)

func TestNewMapEntry(t *testing.T) {
	m := make(map[string]int, 0)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	iterator := NewMapIterator(m)
	slice := ToSlice[*MapEntry[string, int]](iterator)
	t.Log(slice)
}
