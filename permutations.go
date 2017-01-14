package main

import (
	"fmt"
)

func main() {
	fmt.Printf("permutations(%v) = %v\n", []int{}, permutations([]int{}, []int{}, [][]int{}))
	fmt.Printf("permutations(%v) = %v\n", []int{1}, permutations([]int{1}, []int{}, [][]int{}))
	fmt.Printf("permutations(%v) = %v\n", []int{1, 2}, permutations([]int{1, 2}, []int{}, [][]int{}))
	fmt.Printf("permutations(%v) = %v\n", []int{1, 2, 3}, permutations([]int{1, 2, 3}, []int{}, [][]int{}))
}

func permutations(ns []int, is []int, ps [][]int) [][]int {
	if len(ns) == 0 {
		return append(ps, copyAppend(is))
	}

	for i := range ns {
		ps = permutations(copyAppend(ns[:i], ns[i+1:]...), copyAppend(is, ns[i]), ps)
	}

	return ps
}

func copyAppend(a []int, b ...int) []int {
	c := make([]int, len(a)+len(b))
	copy(c, a)
	copy(c[len(a):], b)
	return c
}
