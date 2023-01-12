package iterator

// Iterable 用于表示是可迭代的，拥有一个Iterator方法，调用时返回一个迭代器
type Iterable[T any] interface {

	// Iterator 返回当前对象的迭代器
	Iterator() Iterator[T]
}
