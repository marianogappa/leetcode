package main

import "fmt"

func singleNumber(nums []int) int {
	var a = 0
	for _, n := range nums {
		a ^= n
	}
	return a
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 3, 1, 2}))
}
