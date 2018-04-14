package main

import "fmt"

// Time: O(n*logn) Space: O(1)
func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	// Time: O(n*logn) Space: O(1)
	qsort(nums)

	// Time: O(n) Space: O(1)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}

	return 0 // otherwise it doesn't compile
}

// Time: O(n*logn) Space: O(1)
func qsort(ns []int) {
	if len(ns) <= 1 {
		return
	}
	var (
		l = 0
		r = len(ns) - 1
		p = len(ns) / 2
	)
	ns[r], ns[p] = ns[p], ns[r]
	for i := 0; i < r; i++ {
		if ns[i] < ns[r] {
			ns[i], ns[l] = ns[l], ns[i]
			l++
		}
	}
	ns[l], ns[r] = ns[r], ns[l]

	qsort(ns[:l])
	qsort(ns[l+1:])
}

func main() {
	fmt.Println(findDuplicate([]int{1, 1, 2, 3}) == 1)
	fmt.Println(findDuplicate([]int{1, 2, 2, 3}) == 2)
	fmt.Println(findDuplicate([]int{1, 2, 3, 3}) == 3)
	fmt.Println(findDuplicate([]int{1, 2, 3, 3, 3}) == 3)
	fmt.Println(findDuplicate([]int{3, 2, 1, 2}) == 2)
}
