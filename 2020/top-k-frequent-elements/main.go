package main

import (
	"fmt"
	"reflect"
	"sort"
)

type fn struct {
	f int
	n int
}

// This exercise is trivial if you understand heaps.
// But the key thing to remember here is that heapify
// is O(n), even though it looks O(nlogn). Otherwise,
// there's no way to solve this better than O(nlogn),
// at which point sorting is simpler.
//
// Time: O(klogn)
// Space: O(n)
func topKFrequent(nums []int, k int) []int {
	// 1: get the frequencies in a map
	// Time: O(n)
	// Space: O(n)
	freqs := map[int]int{}
	for _, n := range nums {
		freqs[n]++
	}

	// 2: make a slice of "freq+num"
	// Time: O(n)
	// Space: O(n)
	fns := make([]fn, len(freqs))
	i := 0
	for n, f := range freqs {
		fns[i] = fn{f, n}
		i++
	}

	// 3: heapify slice
	// Time: O(n)
	// Space: O(1)
	heapify(fns)

	// 4: pop max k times
	// Time: O(klogn)
	// Space: O(k)
	topk := make([]int, k)
	for i := 0; i < k; i++ {
		topk[i] = popMax(&fns)
	}

	return topk
}

// Time: O(n)
// Space: O(1)
func heapify(fns []fn) {
	for i := (len(fns) / 2) - 1; i >= 0; i-- {
		siftDown(fns, i)
	}
}

// Time: O(logn)
// Space: O(1)
func siftDown(fns []fn, i int) {
	for root := i; root*2+1 < len(fns); { // while root is a parent
		var child = root*2 + 1                                   // child = left child of root
		if child+1 < len(fns) && fns[child].f < fns[child+1].f { // child = max sibling
			child++
		}
		if fns[root].f < fns[child].f { // if root is unordered to child
			fns[root], fns[child] = fns[child], fns[root] // swap them
			root = child                                  // continue algorithm from child
		} else {
			return // if root and child are ordered, we're done
		}
	}
}

// Time: O(logn)
// Space: O(1)
func popMax(fns *[]fn) int {
	max := (*fns)[0].n
	(*fns)[0], (*fns)[len(*fns)-1] = (*fns)[len(*fns)-1], (*fns)[0]
	(*fns) = (*fns)[:len(*fns)-1]
	siftDown(*fns, 0)
	return max
}

func main() {
	ts := []struct {
		input    []int
		k        int
		expected []int
	}{
		{
			input:    []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			input:    []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			input:    []int{4, 1, -1, 2, -1, 2, 3},
			k:        2,
			expected: []int{-1, 2},
		},
	}
	for _, tc := range ts {
		actual := topKFrequent(tc.input, tc.k)
		sort.Ints(actual)
		sort.Ints(tc.expected)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
