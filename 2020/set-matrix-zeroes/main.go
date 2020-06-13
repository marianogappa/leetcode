package main

import (
	"fmt"
	"reflect"
)

// Time: O(n)
// Space: O(1) assuming using matrix doesn't count
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	var shouldZeroFirstColumn bool

	// Go through the matrix, and when a 0 is found,
	// mark the beginning of the row and column to
	// 0.
	for y := range matrix {
		// There's no way to mark both the first
		// row and the first column, so we choose
		// to make the (0,0) marker mean row, and
		// we add a special boolean for the
		// column.
		if matrix[y][0] == 0 {
			shouldZeroFirstColumn = true
		}
		for x := 1; x < len(matrix[0]); x++ {
			if matrix[y][x] == 0 {
				matrix[0][x] = 0
				matrix[y][0] = 0
			}
		}
	}

	// Zero out the marked cells (except on the initial
	// rows and columns)
	for y := 1; y < len(matrix); y++ {
		for x := 1; x < len(matrix[0]); x++ {
			if matrix[0][x] == 0 || matrix[y][0] == 0 {
				matrix[y][x] = 0
			}
		}
	}

	// Zero out the first row, if the (0,0) marker is
	// set.
	if matrix[0][0] == 0 {
		for x := range matrix[0] {
			matrix[0][x] = 0
		}
	}

	// Zero out the first row, if the special marker is
	// set.
	if shouldZeroFirstColumn {
		for y := range matrix {
			matrix[y][0] = 0
		}
	}
}

func main() {
	ts := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			input: [][]int{
				{1, 1, 1},
				{0, 1, 2},
			},
			expected: [][]int{
				{0, 1, 1},
				{0, 0, 0},
			},
		},
		{
			input: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}
	for _, tc := range ts {
		setZeroes(tc.input)
		if !reflect.DeepEqual(tc.input, tc.expected) {
			fmt.Printf("Expected %v but got %v\n", tc.expected, tc.input)
		}
	}
}
