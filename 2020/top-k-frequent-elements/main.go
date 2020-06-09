package main

import (
	"fmt"
	"reflect"
)

// Time:  O(k*log(n))
// Space: O(n)
func topKFrequent(nums []int, k int) []int {
	// Calculate frequency of each number
	// Time: O(n)
	// Space: O(n)
	freqs := map[int]int{}
	for _, num := range nums {
		freqs[num]++
	}

	// Put previous calculation into a slice
	// Time: O(n)
	// Space: O(n)
	numFreqs := make([]numFreq, len(freqs))
	i := 0
	for num, freq := range freqs {
		numFreqs[i] = numFreq{num, freq}
		i++
	}

	// Time: O(n)
	// Space: O(1)
	heapify(numFreqs)

	h := &maxHeap{numFreqs}

	topK := make([]int, k)
	// Time: O(k*log(n))
	// Space: O(1)
	for i := 0; i < k; i++ {
		// Pop max while keeping the heap property
		topK[i] = h.pop()
	}
	return topK
}

type numFreq struct {
	num, freq int
}

func heapify(arr []numFreq) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		siftDown(arr, i)
	}
}

type maxHeap struct {
	arr []numFreq
}

func (h *maxHeap) pop() int {
	head := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	siftDown(h.arr, 0)
	return head.num
}

func siftDown(arr []numFreq, i int) {
	if len(arr) <= 1 {
		return
	}
	current := i // Starting from the root...

	// While current node has a child...
	for childOf(current) < len(arr) {
		maxChild := maxChildOf(arr, current)

		// If the max child is not > current, done!
		if arr[maxChild].freq <= arr[current].freq {
			return
		}

		// Otherwise swap current and maxChild, and continue from maxChild
		arr[maxChild], arr[current] = arr[current], arr[maxChild]
		current = maxChild
	}
}

func childOf(i int) int {
	return 2*i + 1
}

func maxChildOf(arr []numFreq, i int) int {
	child := childOf(i)
	if child+1 < len(arr) && arr[child+1].freq > arr[child].freq {
		child++
	}
	return child
}

func main() {
	// a := []numFreq{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}, {10, 10}}
	// rand.Seed(time.Now().UnixNano())
	// rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	// fmt.Println("before heapify:", a)
	// heapify(a)
	// fmt.Println("after heapify:", a)
	// h := &maxHeap{a}
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(h.pop())
	// }

	ts := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			nums:     []int{5, 1, -1, -8, -7, 8, -5, 0, 1, 10, 8, 0, -4, 3, -1, -1, 4, -5, 4, -3, 0, 2, 2, 2, 4, -2, -4, 8, -7, -7, 2, -8, 0, -8, 10, 8, -8, -2, -9, 4, -7, 6, 6, -1, 4, 2, 8, -3, 5, -9, -3, 6, -8, -5, 5, 10, 2, -5, -1, -5, 1, -3, 7, 0, 8, -2, -3, -1, -5, 4, 7, -9, 0, 2, 10, 4, 4, -4, -1, -1, 6, -8, -9, -1, 9, -9, 3, 5, 1, 6, -1, -2, 4, 2, 4, -6, 4, 4, 5, -5},
			k:        7,
			expected: []int{4, -1, 2, -5, -8, 0, 8},
		},
	}
	for _, tc := range ts {
		actual := topKFrequent(tc.nums, tc.k)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.nums, tc.k, tc.expected, actual)
		}
	}
}
