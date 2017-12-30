package main

import "fmt"

func numWays(n int, k int) int {
	if k == 0 || n == 0 {
		return 0
	}
	if n == 1 {
		return k
	}
	var (
		diff = k * (k - 1)
		same = k
	)
	for i := 3; i <= n; i++ {
		diff, same = (diff+same)*(k-1), diff
	}
	return diff + same
}

func main() {
	var ts = []struct {
		n, k, e int
	}{
		{n: 0, k: 0, e: 0},
		{n: 0, k: 1, e: 0},
		{n: 1, k: 0, e: 0},
		{n: 1, k: 1, e: 1},
		{n: 2, k: 1, e: 1},
		{n: 2, k: 2, e: 4},
		{n: 50, k: 1, e: 0},
		// {n: 10, k: 2, e: 20},
		{n: 4, k: 2, e: 10},
	}
	for _, t := range ts {
		var a = numWays(t.n, t.k)
		if t.e != a {
			fmt.Printf("numWays(%v, %v) should have been %v but was %v\n", t.n, t.k, t.e, a)
		}
	}
}
