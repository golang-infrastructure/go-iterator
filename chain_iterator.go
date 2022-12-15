package iterator

// ChainIterator 把多个迭代器组合为一个迭代器，
type ChainIterator[T any] struct {

	// 当前消费到第几个迭代器了
	index int

	// 组合的所有的迭代器
	iterators []Iterator[T]
}

var _ Iterator[any] = &ChainIterator[any]{}

func NewChainIterator[T any](iterators ...Iterator[T]) *ChainIterator[T] {
	return &ChainIterator[T]{
		index:     0,
		iterators: iterators,
	}
}

func (x *ChainIterator[T]) Next() bool {

	// 看下指针当前是否已经指向空了
	if x.index >= len(x.iterators) {
		return false
	}

	// 看下当前指向的迭代器是否还有下一个
	hasNext := x.iterators[x.index].Next()
	if hasNext {
		return hasNext
	}

	// 没有的话看下是否还有下一个迭代器，有的话就把指针移过去消费下一个
	if x.index < len(x.iterators) {
		x.index++
		return x.Next()
	}

	// 所有的迭代器都被消费完了
	return false
}

func (x *ChainIterator[T]) Value() T {
	return x.iterators[x.index].Value()
}
