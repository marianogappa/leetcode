package main

import "fmt"

// This solution is poor.
// First optimisation: go right to left; store temps at the same
// time as checking warmers; only need latest pos per temp slot.
// There's also a more complicated stack solution: https://leetcode.com/articles/daily-temperatures/
func dailyTemperatures(temperatures []int) []int {
	ts := make([][]int, 71)
	for i, t := range temperatures {
		ts[t-30] = append(ts[t-30], i)
	}
	res := make([]int, 0)
	for i, t := range temperatures {
		if t == 100 || i == len(temperatures)-1 {
			res = append(res, 0)
			continue
		}
		pos := 1<<32 - 1
		for next := t - 30 + 1; next < 71; next++ {
			for j := 0; j < len(ts[next]); j++ {
				if ts[next][j] < pos && ts[next][j] > i {
					pos = ts[next][j]
				}
			}
		}
		if pos == 1<<32-1 {
			res = append(res, 0)
			continue
		}
		res = append(res, pos-i)
	}
	return res
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
