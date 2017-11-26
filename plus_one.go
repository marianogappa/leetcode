package main

import "fmt"

func plusOne(digits []int) []int {
	cur := len(digits) - 1
	for cur >= 0 && digits[cur]+1 >= 10 {
		digits[cur] = 0
		cur--
	}
	if cur == -1 {
		digits = append([]int{1}, digits...)
	} else {
		digits[cur]++
	}
	return digits
}

func main() {
	fmt.Println("12345", plusOne([]int{1, 2, 3, 4, 5}))
	fmt.Println("12349", plusOne([]int{1, 2, 3, 4, 9}))
	fmt.Println("19999", plusOne([]int{1, 9, 9, 9, 9}))
	fmt.Println("89999", plusOne([]int{8, 9, 9, 9, 9}))
	fmt.Println("99999", plusOne([]int{9, 9, 9, 9, 9}))
}
