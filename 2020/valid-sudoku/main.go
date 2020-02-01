package main

import "fmt"

// Not a lot of magic on this one. Keep a set of used numbers for every
// row, every column and every 3x3 grid. If a number is already on any
// of the sets, return false.
//
// Time: O(n)
// Space: O(n)
func isValidSudoku(board [][]byte) bool {
	var (
		ys [9]map[byte]struct{}
		xs [9]map[byte]struct{}
		ss [3][3]map[byte]struct{}
	)
	for i := 0; i < 9; i++ {
		ys[i] = map[byte]struct{}{}
		xs[i] = map[byte]struct{}{}
	}
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			ss[y][x] = map[byte]struct{}{}
		}
	}
	for y := range board {
		for x := range board[0] {
			if board[y][x] == '.' {
				continue
			}
			sy := y / 3
			sx := x / 3
			if _, ok := ys[y][board[y][x]]; ok {
				return false
			}
			if _, ok := xs[x][board[y][x]]; ok {
				return false
			}
			if _, ok := ss[sy][sx][board[y][x]]; ok {
				return false
			}
			ys[y][board[y][x]] = struct{}{}
			xs[x][board[y][x]] = struct{}{}
			ss[sy][sx][board[y][x]] = struct{}{}
		}
	}
	return true
}

func main() {
	ts := []struct {
		input    [][]byte
		expected bool
	}{
		{
			input: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expected: true,
		},
		{
			input: [][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expected: false,
		},
	}
	for _, tc := range ts {
		actual := isValidSudoku(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
