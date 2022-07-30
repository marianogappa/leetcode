package main

import "fmt"

// Time: O(x+y) where y == len(matrix), x == len(matrix[0])
// Space: O(1)
//
// Brute force: check if every single cell == target with 2 for loops
//              (linear time, constant space)
//
// First optimisation: for every matrix[i], run binary search for target
//                     (constant space, O(len(matrix)*log(len(matrix[0]))) time)
//
// Magical solution you need to be a genius to figure out: if you start a pivot at top-right,
// go left if pivot > target and go down if pivot < target. Either you find it, or you exceed
// the matrix's bounds.
//
// Also works from bottom-left.
func searchMatrix(matrix [][]int, target int) bool {
	// Start at top-right
	x, y := len(matrix[0])-1, 0
	// While in bounds...
	for y >= 0 && y < len(matrix) && x >= 0 && x < len(matrix[y]) {
		pivot := matrix[y][x]
		if pivot == target {
			return true
		}
		if pivot > target {
			x--
		} else { // pivot < target
			y++
		}
	}
	return false
}

func main() {
	ts := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target:   5,
			expected: true,
		},
		{
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target:   20,
			expected: false,
		},
		{
			matrix: [][]int{
				{-5},
			},
			target:   -2,
			expected: false,
		},
		{
			matrix: [][]int{
				{-5},
			},
			target:   -10,
			expected: false,
		},
		{
			matrix: [][]int{
				{1, 1},
			},
			target:   2,
			expected: false,
		},
		{
			matrix: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			target:   15,
			expected: true,
		},
	}
	for _, tc := range ts {
		actual := searchMatrix(tc.matrix, tc.target)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.matrix, tc.target, tc.expected, actual)
		}
	}
}
