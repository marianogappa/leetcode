package main

import (
	"fmt"

	"container/heap"
	"reflect"
)

type elem struct {
	val   int
	point []int
}
type minHeap []elem

func (m minHeap) Less(i, j int) bool { return m[i].val < m[j].val }
func (m minHeap) Swap(i, j int) {
	m[i].val, m[j].val = m[j].val, m[i].val
	m[i].point, m[j].point = m[j].point, m[i].point
}
func (m minHeap) Len() int            { return len(m) }
func (m *minHeap) Push(i interface{}) { *m = append(*m, i.(elem)) }
func (m *minHeap) Pop() interface{} {
	tail := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return tail
}

// Time: O(p + k*logp)
// Space: O(p)
func kClosest(points [][]int, K int) [][]int {
	// Build a heap with the points, using the squared distance as val
	h := &minHeap{}
	for _, point := range points {
		heap.Push(h, elem{point: []int{point[0], point[1]}, val: point[0]*point[0] + point[1]*point[1]})
	}

	// Pull k min points out
	results := [][]int{}
	for i := 0; i < K; i++ {
		e := heap.Pop(h).(elem)
		results = append(results, e.point)
	}
	return results
}

func main() {
	ts := []struct {
		points   [][]int
		k        int
		expected [][]int
	}{
		{
			points:   [][]int{{1, 3}, {-2, 2}},
			k:        1,
			expected: [][]int{{-2, 2}},
		},
		{
			points:   [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:        2,
			expected: [][]int{{3, 3}, {-2, 4}},
		},
	}
	for _, tc := range ts {
		actual := kClosest(tc.points, tc.k)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.points, tc.k, tc.expected, actual)
		}
	}
}
