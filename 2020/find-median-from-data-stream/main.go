package main

import (
	"fmt"
)

type MedianFinder struct {
	low  *heap
	high *heap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{&heap{false, nil}, &heap{true, nil}}
}

// Time: O(log n)
// Space: O(n)
func (this *MedianFinder) AddNum(num int) {
	this.low.push(num)
	this.high.push(this.low.pop())
	if len(this.low.arr) < len(this.high.arr) {
		this.low.push(this.high.pop())
	}
}

// Time: O(1)
// Space: O(1)
func (this *MedianFinder) FindMedian() float64 {
	if len(this.low.arr) == 0 {
		return 0
	}
	if len(this.low.arr) > len(this.high.arr) {
		return float64(this.low.peek())
	}
	return float64((this.low.peek() + this.high.peek())) / 2.0
}

type heap struct {
	isMinHeap bool
	arr       []int
}

func (h *heap) peek() int {
	return h.arr[0]
}

func (h *heap) pop() int {
	head := h.arr[0]
	h.arr[0], h.arr[len(h.arr)-1] = h.arr[len(h.arr)-1], h.arr[0]
	h.arr = h.arr[:len(h.arr)-1]
	h.siftDown()
	return head
}

func (h *heap) push(n int) {
	h.arr = append(h.arr, n)
	h.siftUp()
}

func (h *heap) siftDown() {
	if len(h.arr) <= 1 {
		return
	}
	current := 0
	for h.childOf(current) < len(h.arr) {
		maxChild := h.maxChildOf(current)
		if (h.isMinHeap && h.arr[maxChild] >= h.arr[current]) || (!h.isMinHeap && h.arr[maxChild] <= h.arr[current]) {
			return
		}
		h.arr[maxChild], h.arr[current] = h.arr[current], h.arr[maxChild]
		current = maxChild
	}
}

func (h *heap) siftUp() {
	current := len(h.arr) - 1
	parent := h.parentOf(current)
	for current > 0 && ((h.isMinHeap && h.arr[current] < h.arr[parent]) || (!h.isMinHeap && h.arr[current] > h.arr[parent])) {
		h.arr[current], h.arr[parent] = h.arr[parent], h.arr[current]
		current = parent
		parent = h.parentOf(parent)
	}
}

func (h *heap) parentOf(i int) int {
	if i%2 == 1 {
		return (i - 1) / 2
	}
	return (i - 2) / 2
}

func (h *heap) childOf(i int) int {
	return 2*i + 1
}

func (h *heap) maxChildOf(i int) int {
	child := h.childOf(i)
	if h.isMinHeap && (child+1 < len(h.arr) && h.arr[child+1] < h.arr[child]) {
		child++
	}
	if !h.isMinHeap && (child+1 < len(h.arr) && h.arr[child+1] > h.arr[child]) {
		child++
	}
	return child
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())

	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// rand.Seed(time.Now().UnixNano())
	// rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	// fmt.Println("before heapify:", a)

	// h := &heap{true, nil}
	// for _, n := range a {
	// 	h.push(n)
	// }
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(h.pop())
	// }
}
