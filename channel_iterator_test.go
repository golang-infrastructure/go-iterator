package iterator

import (
	"testing"
)

func TestNewChannelIterator(t *testing.T) {

	channel := make(chan int, 3)
	channel <- 1
	channel <- 2
	channel <- 3
	close(channel)

	iterator := NewChannelIterator(channel)
	slice := ToSlice[int](iterator)
	t.Log(slice)

}
