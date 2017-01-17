package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("permutations(%v) = %v\n", []int{}, permutations([]int{}, []int{}, [][]int{}))
	fmt.Printf("permutations(%v) = %v\n", []int{1}, permutations([]int{1}, []int{}, [][]int{}))
	fmt.Printf("permutations(%v) = %v\n", []int{1, 1}, permutations([]int{1, 1}, []int{}, [][]int{}))
	fmt.Printf("permutations(%v) = %v\n", []int{1, 1, 3, 1}, permutations([]int{1, 1, 3, 1}, []int{}, [][]int{}))
	fmt.Printf("permutations(%v) = %v\n", []int{1, 1, 1, 1}, permutations([]int{1, 1, 1, 1}, []int{}, [][]int{}))
	fmt.Printf("permutations(%v) = %v\n", []int{1, 2, 2, 1}, permutations([]int{1, 2, 2, 1}, []int{}, [][]int{}))
}

func permutations(ns []int, is []int, ps [][]int) [][]int {
	sort.Ints(ns)
	return doPerms(ns, is, ps)
}

func doPerms(ns []int, is []int, ps [][]int) [][]int {
	if len(ns) == 0 {
		return append(ps, copyAppend(is))
	}

	for i := range ns {
		if i+1 < len(ns) && ns[i+1] == ns[i] {
			continue
		}
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
