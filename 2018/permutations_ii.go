package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("permuteUnique(%v) = %v\n", []int{}, permuteUnique([]int{}))
	fmt.Printf("permuteUnique(%v) = %v\n", []int{1}, permuteUnique([]int{1}))
	fmt.Printf("permuteUnique(%v) = %v\n", []int{1, 1}, permuteUnique([]int{1, 1}))
	fmt.Printf("permuteUnique(%v) = %v\n", []int{1, 1, 3, 1}, permuteUnique([]int{1, 1, 3, 1}))
	fmt.Printf("permuteUnique(%v) = %v\n", []int{1, 1, 1, 1}, permuteUnique([]int{1, 1, 1, 1}))
	fmt.Printf("permuteUnique(%v) = %v\n", []int{1, 2, 2, 1}, permuteUnique([]int{1, 2, 2, 1}))
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return doPerms(nums, []int{}, [][]int{})
}

func doPerms(ns []int, is []int, ps [][]int) [][]int {
	if len(ns) == 0 {
		return append(ps, copyAppend(is))
	}

	for i := range ns {
		if i+1 < len(ns) && ns[i+1] == ns[i] {
			continue
		}
		ps = doPerms(copyAppend(ns[:i], ns[i+1:]...), copyAppend(is, ns[i]), ps)
	}

	return ps
}

func copyAppend(a []int, b ...int) []int {
	c := make([]int, len(a)+len(b))
	copy(c, a)
	copy(c[len(a):], b)
	return c
}
