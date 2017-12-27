package main

import "fmt"

func productExceptSelf(nums []int) []int {
	var (
		n     = len(nums)
		o     = make([]int, n)
		left  = 1
		right = 1
	)
	for i := 0; i < n; i++ {
		o[i] = 1
	}
	for i := 0; i < n; i++ {
		o[i] *= left
		left *= nums[i]
		o[n-1-i] *= right
		right *= nums[n-1-i]
	}
	return o
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
