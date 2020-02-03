package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	var j = 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			if i > j {
				nums[j] = nums[i]
			}
			j++
		}
	}
	return j
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 1, 2, 2, 2, 3, 3, 3}))
	fmt.Println(removeDuplicates([]int{}))
	fmt.Println(removeDuplicates([]int{1}))
	fmt.Println(removeDuplicates([]int{-1, -1}))
	fmt.Println(removeDuplicates([]int{-1, 0}))
}
