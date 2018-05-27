package main

import "fmt"

func maxSubArray(nums []int) int {
	var (
		acc  = 0
		maxS = -1 << 31
	)
	for _, n := range nums {
		acc = max(acc+n, n)
		maxS = max(maxS, acc)
	}
	return maxS
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) == 6)
}
