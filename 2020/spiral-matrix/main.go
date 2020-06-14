package main

import (
	"fmt"
	"reflect"
)

// Time: O(n)
// Space: O(n)
//
// Keys here:
// 1. We need exactly m*n iterations. Don't worry
//    about clauses to stop.
// 2. Intuition is to come up with a Space O(1)
//    solution, but it gets too hairy. Just keep
//    a visited matrix.
// 3. Off by one errors when rotating: go over the
//    limit, backtrack, rotate and move again.
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	// Visited matrix to know when to rotate.
	visited := make([][]bool, len(matrix))
	for i := range matrix {
		visited[i] = make([]bool, len(matrix[i]))
	}

	// Clockwise directions.
	dir := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var dirI, x, y int

	result := []int{}
	// This is exactly how many iterations we need.
	// Brilliant idea!
	for i := 0; i < len(matrix)*len(matrix[0]); i++ {
		if x < 0 || x >= len(matrix[0]) || y < 0 || y >= len(matrix) || visited[y][x] {
			// Backtrack
			x -= dir[dirI][0]
			y -= dir[dirI][1]
			// Rotate clockwise
			dirI = (dirI + 1) % 4
			// Move once
			x += dir[dirI][0]
			y += dir[dirI][1]

		}
		result = append(result, matrix[y][x])
		visited[y][x] = true
		x += dir[dirI][0]
		y += dir[dirI][1]
	}

	return result
}

func main() {
	ts := []struct {
		input    [][]int
		expected []int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			input:    [][]int{},
			expected: []int{},
		},
		{
			input:    [][]int{{}},
			expected: []int{},
		},
		{
			input:    [][]int{{1}},
			expected: []int{1},
		},
		{
			input:    [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			input:    [][]int{{1}, {2}, {3}},
			expected: []int{1, 2, 3},
		},
	}
	for _, tc := range ts {
		actual := spiralOrder(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
