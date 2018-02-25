package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	var (
		ls, rs, ms = make([]int, len(nums), len(nums)), make([]int, len(nums), len(nums)), make([]int, len(nums)-k+1, len(nums)-k+1)
		lm, rm     = -1 << 31, -1 << 31
	)
	for i := 0; i < len(nums); i++ {
		if i%k == 0 {
			lm = -1 << 31
		}
		if nums[i] > lm {
			lm = nums[i]
		}
		ls[i] = lm
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i%k == k-1 {
			rm = -1 << 31
		}
		if nums[i] > rm {
			rm = nums[i]
		}
		rs[i] = rm
	}
	var j = 0
	for i := 0; i+k <= len(nums); i++ {
		if ls[i+k-1] > rs[i] {
			ms[j] = ls[i+k-1]
		} else {
			ms[j] = rs[i]
		}
		j++
	}
	return ms
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
