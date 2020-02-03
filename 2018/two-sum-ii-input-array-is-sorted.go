package main

import (
	"fmt"
	"reflect"
)

// I used binary search: for each number from the left, find (target-ns[i]) using binary search
// worst case O(nlogn)
// It's much better to do 2 pointers: start at beginning and end. If sum is > target: high--, if <: low++
// Worst case O(n)

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers) && numbers[i] <= target; i++ {
		if j, ok := bs(i+1, len(numbers)-1, target-numbers[i], numbers); ok {
			return []int{i + 1, j + 1}
		}
	}
	return []int{}
}

func bs(l, r, target int, ns []int) (int, bool) {
	if l > r {
		return 0, false
	}
	var m = (l + r) / 2
	switch {
	case ns[m] > target:
		return bs(l, m-1, target, ns)
	case ns[m] < target:
		return bs(m+1, r, target, ns)
	default:
		return m, true
	}
}

func main() {
	var ts = []struct {
		numbers  []int
		target   int
		expected []int
	}{
		{
			numbers:  []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			numbers:  []int{-1, 0},
			target:   -1,
			expected: []int{1, 2},
		},
	}
	for _, t := range ts {
		var a = twoSum(t.numbers, t.target)
		if !reflect.DeepEqual(t.expected, a) {
			fmt.Printf("Wanted %v but got %v\n", t.expected, a)
		}
	}
}
