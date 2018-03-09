package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	res := make([]int, len(nums))
	for _, n := range nums {
		res[n-1] = -1
	}
	j := 0
	for i, v := range res {
		if v != -1 {
			res[j] = i + 1
			j++
		}
	}
	return res[:j]
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
