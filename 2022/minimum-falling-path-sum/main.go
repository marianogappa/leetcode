package main

import (
	"fmt"
	"math"
)

// Time: O(n)
// Space: O(n)
//
// Think about the matrix falling sum as a tree:
//
// 2 -> +6 -> +7
//         -> +8
//   -> +5 -> +7
//         -> +8
//         -> +9
//   -> +4 -> +8
//         -> +9
//
// We can solve it with recursion but it requires memoization from the bottom up.
// The intuition is that there are a lot of partially repeated calculation branches,
// unless we start from the bottom of the tree:
//
// If we only had one row, the minimum falling path is the minimum number in the array.
//
// If we had two, the minimum falling path is achieved by going through the first row and
// replacing each cell's value with itself plus the minimum of the up to three contiguous
// cells below, and finally finding the minimum number in that replacement.
//
// If we had three, we do the same summing on the middle row, and then on the first row, and
// then get the minimum number in the first row.
//
// And so on.
func minFallingPathSum(matrix [][]int) int {
	for y := len(matrix) - 2; y >= 0; y-- {
		for x := range matrix[y] {
			minOfNextRow := matrix[y+1][x]
			if x > 0 {
				minOfNextRow = min(minOfNextRow, matrix[y+1][x-1])
			}
			if x < len(matrix[y])-1 {
				minOfNextRow = min(minOfNextRow, matrix[y+1][x+1])
			}
			matrix[y][x] += minOfNextRow
		}
	}
	return minOfIntArray(matrix[0])
}

func minOfIntArray(nums []int) int {
	mn := math.MaxInt
	for _, num := range nums {
		mn = min(mn, num)
	}
	return mn
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	ts := []struct {
		input    [][]int
		expected int
	}{
		{
			input:    [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}},
			expected: 13,
		},
		{
			input:    [][]int{{-19, 57}, {-40, -5}},
			expected: -59,
		},
	}
	for _, tc := range ts {
		actual := minFallingPathSum(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
