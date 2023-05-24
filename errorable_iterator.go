package iterator

// ErrorableIterator 表示一个迭代器，迭代器的元素类型是T，这个迭代器在遍历的时候允许发生错误
type ErrorableIterator[T any] interface {
	Iterator[T]

	// NextE 将迭代器的指针往后移动一个，同时返回指针当前指向的位置是否有元素
	NextE() (bool, error)

	// NextE 返回指针指向的元素，如果没有元素的话返回对应类型的零值
	ValueE() (T, error)
}
