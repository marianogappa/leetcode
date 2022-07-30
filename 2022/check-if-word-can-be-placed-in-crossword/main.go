package main

import "fmt"

// Time: O(m*n)
// Space: O(len(word)) for the recursive stack
//
// First trick is to realise that words can only start at top-left edges OR after a block.
// Second trick is to use deltas on functions, cause otherwise all code needs to be duplicated
// for each dimension.
func placeWordInCrossword(board [][]byte, word string) bool {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if !isTopLeftEdge(x, y, 1, 0) && !isCellAfterBlocked(x, y, 1, 0, board) {
				continue
			}
			if placeWordInPos(x, y, 1, 0, board, word) || placeWordInPos(x, y, 1, 0, board, reverse(word)) {
				return true
			}
		}
	}
	for x := 0; x < len(board[0]); x++ {
		for y := 0; y < len(board); y++ {
			if !isTopLeftEdge(x, y, 0, 1) && !isCellAfterBlocked(x, y, 0, 1, board) {
				continue
			}
			if placeWordInPos(x, y, 0, 1, board, word) || placeWordInPos(x, y, 0, 1, board, reverse(word)) {
				return true
			}
		}
	}
	return false
}

func reverse(word string) string {
	var (
		i = 0
		j = len(word) - 1
		w = []byte(word)
	)
	for i < j {
		w[i], w[j] = w[j], w[i]
		i++
		j--
	}
	return string(w)
}

func placeWordInPos(x, y, dx, dy int, board [][]byte, word string) bool {
	if len(word) == 0 {
		return !exists(x, y, board) || board[y][x] == '#'
	}
	return exists(x, y, board) && // Current cell exists
		exists(dx*(x+len(word)-1), dy*(y+len(word)-1), board) && // Dimension has enough space for word
		(board[y][x] == ' ' || board[y][x] == word[0]) && // Current cell is space or first letter of word
		placeWordInPos(x+dx, y+dy, dx, dy, board, word[1:]) // Recursive call for rest of the word
}

func exists(x, y int, board [][]byte) bool {
	return x >= 0 && y >= 0 && y < len(board) && x < len(board[0])
}

func isCellAfterBlocked(x, y, dx, dy int, board [][]byte) bool {
	prevX := x + dx*-1
	prevY := y + dy*-1
	return exists(prevX, prevY, board) && board[prevY][prevX] == '#' && board[y][x] != '#'
}

func isTopLeftEdge(x, y, dx, dy int) bool {
	return (dx == 1 && x == 0) || (dy == 1 && y == 0)
}

func main() {
	ts := []struct {
		board    [][]byte
		word     string
		expected bool
	}{
		{
			board: [][]byte{
				{'#', ' ', '#'},
				{' ', ' ', '#'},
				{'#', 'c', ' '},
			},
			word:     "abc",
			expected: true,
		},
		{
			board: [][]byte{
				{' ', '#', 'a'},
				{' ', '#', 'c'},
				{' ', '#', 'a'},
			},
			word:     "ac",
			expected: false,
		},
		{
			board: [][]byte{
				{'#', ' ', '#'},
				{' ', ' ', '#'},
				{'#', ' ', 'c'},
			},
			word:     "ca",
			expected: true,
		},
		{
			board: [][]byte{
				{'#', ' ', '#'},
				{'#', ' ', '#'},
				{'#', ' ', 'c'},
			},
			word:     "ca",
			expected: true,
		},
	}
	for _, tc := range ts {
		actual := placeWordInCrossword(tc.board, tc.word)
		if tc.expected != actual {
			fmt.Printf("For (%v, %v) expected %v but got %v\n", tc.board, tc.word, tc.expected, actual)
		}
	}
}
