package main

import "fmt"

func qsort(a []int) {
	if len(a) < 2 {
		return
	}
	left, right := 0, len(a)-1
	pivotIndex := len(a) / 2
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	for i := range a {
		if a[i] > a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	qsort(a[:left])
	qsort(a[left+1:])
	return
}

func findKthLargest(nums []int, k int) int {
	qsort(nums)
	return nums[k-1]
}

func main() {
	var ts = []struct {
		nums []int
		k, e int
	}{
		{nums: []int{3, 2, 1, 5, 6, 4}, k: 2, e: 5},
		{nums: []int{3, 2, 1, 5, 6, 4}, k: 3, e: 4},
		{nums: []int{3, 2, 1, 5, 6, 4}, k: 4, e: 3},
		{nums: []int{3, 2, 1, 5, 6, 4}, k: 5, e: 2},
		{nums: []int{3, 3, 3, 3, 3, 3, 3, 3, 3}, k: 8, e: 3},
		{nums: []int{5, 2, 4, 1, 3, 6, 0}, k: 2, e: 5},
	}
	for _, t := range ts {
		var a = findKthLargest(t.nums, t.k)
		if t.e != a {
			fmt.Printf("findKthLargest(%v, %v) should have been %v but was %v\n", t.nums, t.k, t.e, a)
		}
	}
}
