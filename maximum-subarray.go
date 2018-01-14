package main

import "fmt"

func maxSubArray(nums []int) int {
	var maxSoFar, maxEndingHere = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(maxEndingHere+nums[i], nums[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) == 6)
	fmt.Println(maxSubArray([]int{-2}) == -2)
	fmt.Println(maxSubArray([]int{-2, -3, -1}) == -1)
	fmt.Println(maxSubArray([]int{-1, 0}) == 0)
	fmt.Println(maxSubArray([]int{31, -41, 59, 26, -53, 58, 97, -93, -23, 84}) == 187)
}
