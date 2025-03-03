package iterator

import "fmt"

func ExampleIterator() {
	c := &NumberCollection{
		numbers: []int{1, 2, 3, 4, 5},
	}
	it := c.Iterator()
	for val, ok := it.Next(); ok; val, ok = it.Next() {
		fmt.Println(val)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}
