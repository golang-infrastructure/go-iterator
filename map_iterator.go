package iterator

// MapIterator 用于遍历一个Map的迭代器
type MapIterator[K comparable, V any] struct {
	index      int
	entrySlice []*MapEntry[K, V]
}

var _ Iterator[*MapEntry[string, int]] = &MapIterator[string, int]{}

func NewMapIterator[K comparable, V any](m map[K]V) *MapIterator[K, V] {
	entrySlice := make([]*MapEntry[K, V], 0)
	for k, v := range m {
		entrySlice = append(entrySlice, &MapEntry[K, V]{Key: k, Value: v})
	}
	return &MapIterator[K, V]{
		index:      -1,
		entrySlice: entrySlice,
	}
}

func (x *MapIterator[K, V]) Next() bool {
	if x.index+1 < len(x.entrySlice) {
		x.index++
		return true
	} else {
		return false
	}
}

func (x *MapIterator[K, V]) Value() *MapEntry[K, V] {
	return x.entrySlice[x.index]
}

// ------------------------------------------------ MapEntry -----------------------------------------------------------

type MapEntry[K, V any] struct {
	Key   K
	Value V
}

func (x *MapEntry[K, V]) GetKey() K {
	return x.Key
}

func (x *MapEntry[K, V]) GetValue() V {
	return x.Value
}

// ------------------------------------------------ ---------------------------------------------------------------------
