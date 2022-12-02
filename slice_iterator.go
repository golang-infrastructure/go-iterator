package iterator

type SliceIterator[T any] struct {
	index int
	slice []T
}

var _ Iterator[any] = &SliceIterator[any]{}

func NewSliceIterator[T any](slice []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		index: 0,
		slice: slice,
	}
}

func (x *SliceIterator[T]) Next() bool {
	return x.index < len(x.slice)
}

func (x *SliceIterator[T]) Value() T {
	r := x.slice[x.index]
	x.index++
	return r
}
