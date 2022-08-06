package main

import (
	"fmt"
)

// Time: O(m*n) where m == len(points), n == len(points[0])
// Space: O(n), because we need to store one row with aggregates
//
// The more intuitive cubic solution times out.
//
// First intuition: knowing the optimal choices in one row only requires knowing that row and the next one. So, try to
// find a linear solution to each two rows. Then, final solution will be linear to cells in points.
//
// Using two subsequent rows, "prevRow" & "curRow", the optimal "prevRow[i]" for every "curRow[i]" can be calculated
// linearly by greedily tracking a "maximum" pivot in "prevRow", doing one pass from the left and one from the right.
func maxPoints(points [][]int) int64 {
	resultRow := points[0]

	// Take rows by twos: (0th, 1th) => result; (result, 2th) => result; (result, 3th) => result
	for i := 1; i < len(points); i++ {
		prevRow := cloneRow(resultRow)
		populateMaxLeftToRightOntoResultRow(prevRow, points[i], resultRow)
		populateMaxRightToLeftOntoResultRow(prevRow, points[i], resultRow)
	}

	return int64(maxOfSlice(resultRow))
}

func populateMaxLeftToRightOntoResultRow(prevRow, curRow, resultRow []int) {
	var j int
	// i starts at the left, and always points at the last known max cell in prevRow, like a pivot.
	for i := 0; i < len(prevRow); i = j {
		// This next loop's purpose is to populate maxes using the pivot, until a new pivot is discovered.
		// j starts at pivot, and moves towards the right until it finds a prevRow cell higher than pivot.
		for j = i; j < len(prevRow) && prevRow[j]-(j-j) <= prevRow[i]-(j-i); j++ {
			// This line will run once per cell, always with the right pivot.
			resultRow[j] = curRow[j] + prevRow[i] - (j - i)
		}
	}
}

func populateMaxRightToLeftOntoResultRow(prevRow, curRow, resultRow []int) {
	var j int
	// i starts at the right, and always points at the last known max cell in prevRow, like a pivot.
	for i := len(prevRow) - 1; i >= 0; i = j {
		// This next loop's purpose is to populate maxes using the pivot, until a new pivot is discovered.
		// j starts at pivot, and moves towards the left until it finds a prevRow cell higher than pivot.
		for j = i; j >= 0 && prevRow[j]-(j-j) <= prevRow[i]-(i-j); j-- {
			// This line will run once per cell, always with the right pivot.
			// But in this case, resultRow is already populated, so only override it with larger values.
			resultRow[j] = max(resultRow[j], curRow[j]+prevRow[i]-(i-j))
		}
	}
}

func cloneRow(row []int) []int {
	cloned := make([]int, len(row))
	copy(cloned, row)
	return cloned
}

func maxOfSlice(slice []int) (result int) {
	for _, val := range slice {
		result = max(result, val)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ts := []struct {
		input    [][]int
		expected int64
	}{
		{
			input:    [][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}},
			expected: 9,
		},
		{
			input:    [][]int{{1, 5}, {2, 3}, {4, 2}},
			expected: 11,
		},
		{
			input:    [][]int{{1, 5}, {2, 3}, {4, 2}},
			expected: 11,
		},
	}
	for _, tc := range ts {
		actual := maxPoints(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
