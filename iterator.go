package iterator

import "reflect"

// ---------------------------------------------------------------------------------------------------------------------

// Iterator 表示一个迭代器，迭代器的元素类型是T
type Iterator[T any] interface {

	// Next 将迭代器的指针往后移动一个，同时返回指针当前指向的位置是否有元素
	Next() bool

	// Value 返回指针指向的元素，如果没有元素的话返回对应类型的零值
	Value() T
}

// ---------------------------------------------------------------------------------------------------------------------

// IsIterator 判断给定的struct是否实现了迭代器
func IsIterator[T any](structObject any) bool {
	reflectValue := reflect.ValueOf(structObject)
	// 拥有Next方法，并且签名符合
	if !reflectValue.CanInterface() {
		return false
	}
	_, ok := reflectValue.Interface().(Iterator[T])
	return ok
}

// CastToIterator 转换为迭代器
func CastToIterator[T any](structObject any) (Iterator[T], error) {
	reflectValue := reflect.ValueOf(structObject)
	// 拥有Next方法，并且签名符合
	if !reflectValue.CanInterface() {
		return nil, ErrIsNotIterator
	}
	iterator, ok := reflectValue.Interface().(Iterator[T])
	if !ok {
		return nil, ErrIsNotIterator
	}
	return iterator, nil
}

// DoIterator 遍历一个迭代器
func DoIterator[T any](iterator Iterator[T], iteratorFunc func(value T) bool) {
	for iterator.Next() && iteratorFunc(iterator.Value()) {
	}
}

// ---------------------------------------------------------------------------------------------------------------------
