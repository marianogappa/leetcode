package main

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		r   = 1
		l   = 1
		res = -1 << 31
		j   = len(nums) - 1
	)
	for i := range nums {
		r *= nums[i]
		l *= nums[j]
		if l > res {
			res = l
		}
		if r > res {
			res = r
		}
		res = max(l, max(r, res))
		if r == 0 {
			r = 1
		}
		if l == 0 {
			l = 1
		}
		j--
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}) == 6)
	fmt.Println(maxProduct([]int{0, 2, 3, -2, 4}) == 6)
	fmt.Println(maxProduct([]int{1, 0, 2, 3, -2, 4}) == 6)
	fmt.Println(maxProduct([]int{1, 0, 2, 3, -2, 4, 0, 0}) == 6)
	fmt.Println(maxProduct([]int{1, 0, 1, 0, 1, 0, 2, 3, -2, 4, 0, 0}) == 6)
	fmt.Println(maxProduct([]int{-1, -7, 0, 1, 0, 1, 0, 2, 3, -2, 4, 0, 0}) == 7)
	fmt.Println(maxProduct([]int{}) == 0)
	fmt.Println(maxProduct([]int{0}) == 0)
	fmt.Println(maxProduct([]int{0, 0, 0, 0}) == 0)
	fmt.Println(maxProduct([]int{0, -1, 0, 0}) == 0)
	fmt.Println(maxProduct([]int{-1}) == -1)
}
