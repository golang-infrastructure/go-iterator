package iterator

func FromSlice[T any](slice []T) Iterator[T] {
	return NewSliceIterator(slice)
}

func FromChannel[T any](channel <-chan T) Iterator[T] {
	return NewChannelIterator(channel)
}
