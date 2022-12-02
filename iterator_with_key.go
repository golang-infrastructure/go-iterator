package iterator

type IteratorWithKey[T any] interface {
	Iterator[T]
	Key() T
	Begin()
	First() bool
	NextTo(func(key T, value T) bool) bool
}

type ReverseIteratorWithKey[T any] interface {
	IteratorWithKey[T]
	Prev() bool
	End()
	Last() bool
	PrevTo(func(key interface{}, value interface{}) bool) bool
}
