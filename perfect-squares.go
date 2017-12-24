package main

import "fmt"

func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	var (
		ns = make([]int, 0)
		ni = 1
	)
	for {
		if ni*ni > n {
			break
		}
		ns = append(ns, ni*ni)
		ni++
	}
	return solve(n, ns, 0, 1<<32-1) // assumes solution exists
}

func solve(n int, ns []int, part, min int) int {
	if part >= min {
		return min
	}
	if n == 0 {
		if part < min {
			min = part
		}
		return min
	}
	for i := len(ns) - 1; i >= 0; i-- {
		if n-ns[i] < 0 {
			continue
		}
		var lmin = solve(n-ns[i], ns[:i+1], part+1, min)
		if lmin < min {
			min = lmin
		}
	}
	return min
}

func main() {
	var ts = []struct {
		i int
		e int
	}{
		{i: 1, e: 1},
		{i: 3, e: 3},
		{i: 4, e: 1},
		{i: 9, e: 1},
		{i: 16, e: 1},
		{i: 12, e: 3},
		{i: 13, e: 2},
		{i: 998001, e: 1},
		{i: 998002, e: 2},
	}
	for _, t := range ts {
		var a = numSquares(t.i)
		if t.e != a {
			fmt.Printf("numSquares(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}
