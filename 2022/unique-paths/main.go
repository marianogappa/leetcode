package main

import "fmt"

// Time: O(m*n) because memoisations make sure each cell's count contribution is calculated only once
// Space: O(m*n) because the memo array has m*n size
//
// This is a very straightforward DP with memo.
//
// The number of paths from (0,0) to the solution is the sum of the solutions of starting from (0, 1) plus the ones
// starting from (1, 0). In turn, those are the sums of the solutions of their right and bottom cells, and so on.
// At the end, there's only one solution for the entire last row & last column, which is going stright down or right.
//
// By configuring a memo array with one entry per cell, and making sure the last cell equals 1, one can confidently
// use the recursive approach described above.
func uniquePaths(m int, n int) int {
	memo := make([]int, m*n)

	// It would actually be more efficient to set the entire last columns and rows to 1.
	memo[len(memo)-1] = 1

	return doUniquePaths(0, 0, m, n, memo)
}

func doUniquePaths(x, y, maxX, maxY int, memo []int) int {
	// Small trick to condense (x, y) coordinates into one dimensional array.
	idx := x*maxY + y

	// Use memo if available. Zero value means "not set yet", and this is ok because all cells have at least 1 solution.
	if memo[idx] != 0 {
		return memo[idx]
	}

	// Count possible solutions for the current cell:
	count := 0

	// Sum the solutions of going to the right, if it's possible to go there.
	if x < maxX-1 {
		count += doUniquePaths(x+1, y, maxX, maxY, memo)
	}

	// Sum the solutions of going to the bottom, if it's possible to go there.
	if y < maxY-1 {
		count += doUniquePaths(x, y+1, maxX, maxY, memo)
	}

	// Update the memo and return the count
	memo[idx] = count
	return memo[idx]
}

func main() {
	ts := []struct {
		m, n     int
		expected int
	}{
		{
			m:        1,
			n:        1,
			expected: 1,
		},
		{
			m:        1,
			n:        2,
			expected: 1,
		},
		{
			m:        2,
			n:        1,
			expected: 1,
		},
		{
			m:        2,
			n:        2,
			expected: 2,
		},
		{
			m:        3,
			n:        2,
			expected: 3,
		},
		{
			m:        3,
			n:        7,
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
