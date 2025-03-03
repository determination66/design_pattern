package iterator

// 迭代器接口
type Iterator interface {
	Next() (int, bool) // 返回元素和是否有下一个1
}

type NumberIterator struct {
	numbers []int
	index   int
}

func (it *NumberIterator) Next() (int, bool) {
	if it.index < len(it.numbers) {
		val := it.numbers[it.index]
		it.index++
		return val, true
	}
	return 0, false
}

type NumberCollection struct {
	numbers []int
}

func (c *NumberCollection) Iterator() Iterator {
	return &NumberIterator{numbers: c.numbers, index: 0}
}
