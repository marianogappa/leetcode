package main

import "fmt"

func rotate(nums []int, k int) {
	k %= len(nums)
	if k == 0 {
		return
	}
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	var ns = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(ns, 3)
	fmt.Println(ns)
}
