package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var (
		bits     int   = 64
		i64      int64 = 1 << 32
		finalNum       = 0
	)
	if int(i64) == 0 {
		bits = 32
	}

	for i := 0; i < bits; i++ {
		c := 0
		for _, num := range nums {
			if ((num >> uint(i)) & 1) == 1 {
				c++
			}
		}
		if c%3 != 0 {
			finalNum |= 1 << uint(i)
		}
	}
	return finalNum
}

func main() {
	fmt.Println(singleNumber([]int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}))
}
