package iterator

// ToSlice 把迭代器转为切片
func ToSlice[T any](iterator Iterator[T]) []T {
	// 先转为切片
	slice := make([]T, 0)
	for iterator.Next() {
		slice = append(slice, iterator.Value())
	}
	return slice
}

// ToChannel 把迭代器转为channel
func ToChannel[T any](iterator Iterator[T]) chan T {
	slice := ToSlice(iterator)
	// 再转为channel
	channel := make(chan T, len(slice))
	for _, value := range slice {
		channel <- value
	}
	return channel
}
