package main

import "fmt"

// Time: O(n)
// Space: O(n) because the grid is reused (could be cloned)
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var count int
	for y := range grid {
		for x := range grid[y] {
			count += dfs(grid, x, y)
		}
	}
	return count
}

func dfs(grid [][]byte, x, y int) int {
	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[0]) {
		return 0
	}
	switch grid[y][x] {
	case '0':
		return 0
	default:
		grid[y][x] = '0'
		dfs(grid, x-1, y)
		dfs(grid, x, y-1)
		dfs(grid, x+1, y)
		dfs(grid, x, y+1)
		return 1
	}
}

func main() {
	ts := []struct {
		input    [][]byte
		expected int
	}{
		{
			input: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			input: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
	}
	for _, tc := range ts {
		actual := numIslands(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
