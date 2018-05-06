package main

import "fmt"

type heap struct {
	ns []int
}

// Time: O(nlogn)
func findKthLargest(nums []int, k int) int {
	var h = &heap{}

	// Time: O(nlogn)
	for _, n := range nums {
		h.insert(n)
	}
	// Time: O(klogn)
	for i := 1; i < k; i++ {
		h.popMax()
	}
	return h.ns[0]
}

func (h *heap) siftDown() {
	if len(h.ns) <= 1 {
		return
	}
	var root = 0
	for 2*root+1 < len(h.ns) {
		var child = 2*root + 1
		if child+1 < len(h.ns) && h.ns[child+1] >= h.ns[child] {
			child++
		}
		if h.ns[child] <= root {
			break
		}
		h.ns[child], h.ns[root], root = h.ns[root], h.ns[child], child
	}
}

func (h *heap) popMax() int {
	var max = h.ns[0]
	h.ns[0] = h.ns[len(h.ns)-1]
	h.ns = h.ns[:len(h.ns)-1]
	h.siftDown()
	return max
}

func (h *heap) insert(n int) {
	h.ns = append(h.ns, n)
	h.siftUp()
}

func (h *heap) siftUp() {
	var (
		cur    = len(h.ns) - 1
		parent = h.parentOf(cur)
	)
	for cur != 0 && h.ns[parent] < h.ns[cur] {
		h.ns[parent], h.ns[cur] = h.ns[cur], h.ns[parent]
		cur, parent = parent, h.parentOf(parent)
	}
}

func (h *heap) parentOf(i int) int {
	if i%2 == 1 {
		return (i - 1) / 2
	}
	return (i - 2) / 2
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 9))
}
