package iterator

// RangeIterator 范围递增的迭代器
type RangeIterator struct {
	begin, end int
	current    int
}

var _ Iterator[int] = &RangeIterator{}

// NewRangeIterator [begin, end)的左闭右开区间
func NewRangeIterator(begin, end int) *RangeIterator {
	return &RangeIterator{
		begin:   begin,
		end:     end,
		current: begin - 1,
	}
}

func (x *RangeIterator) Next() bool {
	// 不包含end
	if x.current+1 < x.end {
		x.current++
		return true
	}
	return false
}

func (x *RangeIterator) Value() int {
	r := x.current
	return r
}
