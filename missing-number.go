package main

import "fmt"

func missingNumber(nums []int) int {
	var sum = 0
	for _, n := range nums {
		sum += n
	}
	return (len(nums)*(len(nums)+1))/2 - sum
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}) == 2)
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}) == 8)
	fmt.Println(missingNumber([]int{0}) == 1)
	fmt.Println(missingNumber([]int{1}) == 0)
	fmt.Println(missingNumber([]int{0, 1}) == 2)
}
