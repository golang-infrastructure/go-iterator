package iterator

type IteratorWithIndex[T any] interface {
	Iterator[T]
	Index() int
	Begin()
	First() bool
	NextTo(func(index int, value T) bool) bool
}

type ReverseIteratorWithIndex[T any] interface {
	IteratorWithIndex[T]
	Prev() bool
	End()
	Last() bool
	PrevTo(func(index int, value interface{}) bool) bool
}
