package main

import "fmt"

// Time: O(n^2*4^(len(word)))
// Space: O(1) or O(n^2) if the board counts
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	for y := range board {
		for x := range board[y] {
			if doExist(x, y, word, board) {
				return true
			}
		}
	}
	return false
}

func doExist(x, y int, word string, board [][]byte) bool {
	if len(word) == 0 {
		return true
	}
	if x < 0 || y < 0 || x >= len(board[0]) || y >= len(board) || board[y][x] != word[0] {
		return false
	}

	// Remove the letter so that we don't use it again
	temp := board[y][x]
	board[y][x] = ' '

	for _, dir := range [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
		if doExist(x+dir[0], y+dir[1], word[1:], board) {
			return true
		}
	}

	// Backtrack
	board[y][x] = temp
	return false
}

func main() {
	ts := []struct {
		board    [][]byte
		word     string
		expected bool
	}{
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCCED",
			expected: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "SEE",
			expected: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCB",
			expected: false,
		},
	}
	for _, tc := range ts {
		actual := exist(tc.board, tc.word)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.board, tc.word, tc.expected, actual)
		}
	}
}
