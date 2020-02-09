package main

import "fmt"

// This is basic DP with memo.
//
// The possible paths at a given position is equal to the
// sum of possible paths from the neighbouring down and
// right cells, and that number is the same for every
// branch that leads there, so we can store it in a map.
//
// Thus every position's possible paths will eventually be
// on the map, and we don't need to calculate them more
// than once. This is why is linear both in space and time.
//
// Optimisation: when we reach the right or bottom corner,
// there's only one possible solution from there, so we
// don't need to continue: the result is 1 from there.
//
// Time: O(m*n) because of the memoisation
// Space: O(m*n) because every grid cell will be on the map
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	return doUniquePaths(m-1, n-1, map[xy]int{})
}

type xy struct{ x, y int }

func doUniquePaths(m int, n int, memo map[xy]int) int {
	if m == 0 && n == 0 {
		return 1
	}
	if _, ok := memo[xy{m, n}]; !ok {
		if m > 0 && n > 0 {
			memo[xy{m, n}] = doUniquePaths(m-1, n, memo) + doUniquePaths(m, n-1, memo)
		} else {
			memo[xy{m, n}] = 1
		}
	}
	return memo[xy{m, n}]
}

func main() {
	ts := []struct {
		m, n     int
		expected int
	}{
		{
			m:        3,
			n:        2,
			expected: 3,
		},
		{
			m:        7,
			n:        3,
			expected: 28,
		},
	}
	for _, tc := range ts {
		actual := uniquePaths(tc.m, tc.n)
		if tc.expected != actual {
			fmt.Printf("For (%v, %v) expected %v but got %v\n", tc.m, tc.n, tc.expected, actual)
		}
	}
}
