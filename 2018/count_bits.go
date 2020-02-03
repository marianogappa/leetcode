package main

import "fmt"

func countBits(num int) []int {
	r := make([]int, num+1)
	for i := 1; i <= num; i++ {
		r[i] = r[i&(i-1)] + 1
	}
	return r
}

func main() {
	fmt.Println(countBits(20))
}
