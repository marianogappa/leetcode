package main

import (
	"fmt"
)

// Time: O(log n)
// Space: O(n)
//
// The trick here is that there's no way to tell what n is gonna be on Next(n) so one cannot assume iterator is gonna
// go one by one. Which means binary search is the best approach, with an evolving startOffset as the iterator is
// consumed.
type RLEIterator struct {
	startOffset, i int
	ranges         []rangedValue
}

type rangedValue struct {
	start, end, value int
}

func Constructor(encoding []int) RLEIterator {
	iter := RLEIterator{i: -1}
	valueCount := 0
	for i := 0; i < len(encoding); i += 2 {
		if encoding[i] == 0 {
			continue
		}
		iter.ranges = append(iter.ranges, rangedValue{start: valueCount, end: valueCount + encoding[i] - 1, value: encoding[i+1]})
		valueCount += encoding[i]
	}
	return iter
}

func (this *RLEIterator) Next(n int) int {
	this.i += n
	if this.i > this.ranges[len(this.ranges)-1].end {
		return -1
	}
	var value int
	value, this.startOffset = binarySearch(this.i, this.ranges, this.startOffset, len(this.ranges)-1)
	return value
}

func binarySearch(idx int, ranges []rangedValue, min, max int) (int, int) {
	mid := (min + max) / 2
	if ranges[mid].start > idx {
		return binarySearch(idx, ranges, min, mid-1)
	}
	if ranges[mid].end < idx {
		return binarySearch(idx, ranges, mid+1, max)
	}
	return ranges[mid].value, mid
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(encoding);
 * param_1 := obj.Next(n);
 */

func main() {
	encoding := []int{3, 8, 0, 9, 2, 5}
	obj := Constructor(encoding)
	pObj := &obj
	fmt.Println(pObj.Next(2))
	fmt.Println(pObj.Next(1))
	fmt.Println(pObj.Next(1))
	fmt.Println(pObj.Next(2))
}
