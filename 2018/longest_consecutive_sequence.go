package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	var (
		m         = make(map[int]struct{})
		maxConsec = 0
	)
	for _, n := range nums {
		m[n] = struct{}{}
	}
	for n := range m {
		var (
			c       = 1
			smaller = n - 1
			bigger  = n + 1
		)
		for {
			if _, ok := m[smaller]; !ok {
				break
			}
			delete(m, smaller)
			smaller--
			c++
		}
		for {
			if _, ok := m[bigger]; !ok {
				break
			}
			delete(m, bigger)
			bigger++
			c++
		}

		if c > maxConsec {
			maxConsec = c
		}
	}
	return maxConsec
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{2147483646, -2147483647, 0, 2, 2147483644, -2147483645, 2147483645}))
}
