package main

import "fmt"

type xy struct {
	x, y int
}

// Time: O(m*n)
// Space: O(m*n)
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	// -1 is to make it zero-based
	return doUniquePaths(m-1, n-1, map[xy]int{})
}

func doUniquePaths(m int, n int, memo map[xy]int) int {
	if m < 0 || n < 0 {
		return 0
	}
	if m == 0 && n == 0 {
		return 1
	}
	pos := xy{m, n}
	if v, ok := memo[pos]; ok {
		return v
	}
	memo[pos] = doUniquePaths(m-1, n, memo) + doUniquePaths(m, n-1, memo)
	return memo[pos]
}

func main() {
	ts := []struct {
		m        int
		n        int
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
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.m, tc.n, tc.expected, actual)
		}
	}
}
