package main

import "fmt"

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	var (
		ns = make([][]int, n)
		c  = 1
	)
	for i := 3; i < n; i += 2 {
		if len(ns[i]) == 0 {
			c++
			var next = i + 2*i
			if next < n {
				ns[next] = append(ns[next], i)
			}
			continue
		}
		for _, p := range ns[i] {
			var next = i + 2*p
			if next < n {
				ns[next] = append(ns[next], p)
			}
		}
	}
	return c
}

func main() {
	var ts = []struct {
		i int
		e int
	}{
		{0, 0},
		{2, 0},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 3},
		{7, 3},
		{8, 4},
		{9, 4},
		{100, 25},
		{499979, 41537},
	}
	for _, t := range ts {
		var a = countPrimes(t.i)
		if t.e != a {
			fmt.Printf("countPrimes(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}
