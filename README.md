# Golang Iterator

# 一、简介

Golang迭代器模式的定义与实现，用于在一些适合迭代器的场景使用。

# 二、安装

```bash
go get -u github.com/golang-infrastructure/go-iterator
```

# 三、Example

```go
slice := []int{1, 2, 3, 4, 5}
iterator := NewSliceIterator(slice)
toSlice := ToSlice[int](iterator)
assert.Equal(t, slice, toSlice)
```

# 四、API

## Iterator[T any]

```go
// Iterator 表示一个迭代器，迭代器的元素类型是T
type Iterator[T any] interface {

	// Next 将迭代器的指针往后移动一个，同时返回指针当前指向的位置是否有元素
	Next() bool

	// Value 返回指针指向的元素，如果没有元素的话返回对应类型的零值
	Value() T
}
```

## IteratorWithIndex[T any]

## IteratorWithKey[T any]

## ChainIterator[T any] 

用来把多个迭代器组合为一个：

```go
	iteratorA := NewRangeIterator(0, 5)
	iteratorB := NewRangeIterator(5, 10)

	iterator := NewChainIterator[int](iteratorA, iteratorB)
	slice := ToSlice[int](iterator)
	t.Log(slice)
```

## ChannelIterator[T any]

用来把channel封装为迭代器：

```go
	channel := make(chan int, 3)
	channel <- 1
	channel <- 2
	channel <- 3
	close(channel)

	iterator := NewChannelIterator(channel)
	slice := ToSlice[int](iterator)
	t.Log(slice)
```

## CycleIterator[T any]

用来无限循环返回某个序列： 

```go
	iterator := NewCycleIterator(1, 2, 3)
	for i := 0; i < 100 && iterator.Next(); i++ {
		t.Log(iterator.Value())
	}
```

## MapIterator[K comparable, V any]

用来迭代Map：

```go
	m := make(map[string]int, 0)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	iterator := NewMapIterator(m)
	slice := ToSlice[*MapEntry[string, int]](iterator)
	t.Log(slice)
```

## RangeIterator

用来返回一个范围：

```go
	begin := -100
	end := 100
	iterator := NewRangeIterator(begin, end)
	slice := ToSlice[int](iterator)

	exceptedSlice := make([]int, 0)
	for begin < end {
		exceptedSlice = append(exceptedSlice, begin)
		begin++
	}

	assert.Equal(t, exceptedSlice, slice)
```

## SliceIterator[T any]

用来把切片包装为迭代器： 

```go
	slice := []int{1, 2, 3, 4, 5}
	iterator := NewSliceIterator(slice)
	toSlice := ToSlice[int](iterator)
	assert.Equal(t, slice, toSlice)
```
