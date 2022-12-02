package iterator

// RangeIterator 范围递增的迭代器
type RangeIterator struct {
	begin, end int
	current    int
}

var _ Iterator[int] = &RangeIterator{}

func NewRangeIterator(begin, end int) *RangeIterator {
	return &RangeIterator{
		begin:   begin,
		end:     end,
		current: begin,
	}
}

func (x *RangeIterator) Next() bool {
	return x.current <= x.end
}

func (x *RangeIterator) Value() int {
	r := x.current
	x.current += 1
	return r
}
