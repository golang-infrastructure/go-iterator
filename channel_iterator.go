package iterator

// ChannelIterator 用于把一个channel封装为iterator
type ChannelIterator[T any] struct {
	channel <-chan T
	hasNext bool
	value   T
}

var _ Iterator[any] = &ChannelIterator[any]{}

func NewChannelIterator[T any](channel <-chan T) *ChannelIterator[T] {
	return &ChannelIterator[T]{
		channel: channel,
	}
}

func (x *ChannelIterator[T]) Next() bool {
	// 尝试从channel中读取值
	x.value, x.hasNext = <-x.channel
	return x.hasNext
}

func (x *ChannelIterator[T]) Value() T {
	return x.value
}
