package iterator

// SliceIterator 用于把切片封装为迭代器
type SliceIterator[T any] struct {
	// 当前遍历到哪个下标了
	index int
	// 正在被遍历的切片
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
