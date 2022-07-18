package main

import "fmt"

// Time: O(m*n)
// Space: O(m*n) because we mutate the grid and depend on that state
//
// Notice that the value at the bottom-right depends on no other value, the ones on top and left of it depend only on it,
// and so on the chain goes.
//
// This means that, as long as we traverse the matrix in a particular way, every cell can be calculated only once based
// on previous calculations of values on the bottom and the right.
//
// The way to traverse, thus, is to put a pivot on bottom-right, calculate sums to the top and to the left of the pivot,
// and then repeat moving the pivot diagonally towards top-left. It doesn't matter if the matrix is not square; just
// stop when any of the axis exceeded index zero.
func minPathSum(grid [][]int) int {
	pivotY := len(grid) - 1
	pivotX := len(grid[0]) - 1
	for pivotX >= 0 && pivotY >= 0 {
		updateFromPivot(pivotX, pivotY, grid)
		pivotX--
		pivotY--
	}

	return grid[0][0]
}

func updateFromPivot(pivotX, pivotY int, grid [][]int) {
	for y := pivotY; y >= 0; y-- {
		grid[y][pivotX] += minOfMaybeInts(at(pivotX, y+1, grid), at(pivotX+1, y, grid))
	}
	for x := pivotX - 1; x >= 0; x-- {
		grid[pivotY][x] += minOfMaybeInts(at(x+1, pivotY, grid), at(x, pivotY+1, grid))
	}
}

func at(x, y int, grid [][]int) *int {
	if !exists(x, y, grid) {
		return nil
	}
	return &grid[y][x]
}

func minOfMaybeInts(num2, num3 *int) int {
	if num2 == nil && num3 == nil {
		return 0
	}
	if num2 == nil {
		return (*num3)
	}
	if num3 == nil {
		return (*num2)
	}
	return min(*num2, *num3)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func exists(x, y int, grid [][]int) bool {
	return y < len(grid) && x < len(grid[y])
}

func main() {
	ts := []struct {
		input    [][]int
		expected int
	}{
		{
			input: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1}},
			expected: 7,
		},
		{
			input:    [][]int{{1, 2, 3}, {4, 5, 6}},
			expected: 12,
		},
	}
	for _, tc := range ts {
		actual := minPathSum(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
