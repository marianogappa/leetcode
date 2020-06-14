package main

import (
	"fmt"
	"reflect"
)

// Time: O(n)
// Space: O(1)
func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	// Draw an imaginary diagonal
	// from (0,0) to (len(x)-1,len(y)-1)
	// and flip the values
	for y := range matrix {
		for x := range matrix[y] {
			if y == x {
				break
			}
			matrix[y][x], matrix[x][y] = matrix[x][y], matrix[y][x]
		}
	}

	// Flip horizontally
	for y := range matrix {
		for x := 0; x < len(matrix[y])/2; x++ {
			matrix[y][x], matrix[y][len(matrix[y])-1-x] = matrix[y][len(matrix[y])-1-x], matrix[y][x]
		}
	}
}

func main() {
	ts := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input:    [][]int{},
			expected: [][]int{},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			input: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}
	for _, tc := range ts {
		rotate(tc.input)
		if !reflect.DeepEqual(tc.input, tc.expected) {
			fmt.Printf("Expected %v but got %v\n", tc.expected, tc.input)
		}
	}
}
