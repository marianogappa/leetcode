package main

import "fmt"

func majorityElement(nums []int) int {
	maj := len(nums) / 2
	if len(nums)%2 == 1 {
		maj = len(nums)/2 + 1
	}
	m := make(map[int]int, len(nums))
	for _, n := range nums {
		m[n] += 1
		if m[n] == maj {
			return n
		}
	}
	return -1
}

func main() {
	fmt.Println(majorityElement([]int{1, 2, 2, 2, 3}))
}
