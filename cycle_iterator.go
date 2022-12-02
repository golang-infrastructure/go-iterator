package iterator

// CycleIterator 把一个给定的序列无限循环
type CycleIterator[T any] struct {
	index int
	slice []T
}

var _ Iterator[any] = &CycleIterator[any]{}

func NewCycleIterator[T any](slice ...T) *CycleIterator[T] {
	return &CycleIterator[T]{
		index: 0,
		slice: slice,
	}
}

func (x *CycleIterator[T]) Next() bool {
	// 序列循环迭代器没有穷尽
	return true
}

func (x *CycleIterator[T]) Value() T {
	r := x.slice[x.index]
	x.index = (x.index + 1) % len(x.slice)
	return r
}
