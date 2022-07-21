package main

import "fmt"

// Time: O(m*n)
// Space: O(m*n) because the board is used as state
//
// This is the same exercise as Number of Islands. Go through the board and count++ when you find a battleship. But
// before you continue looking, sink the battleship by going right and down recursively until you run out of battleship
// indicators. The board acts as memo, but if this is prohibited then you can just have a visited [][]byte clone.
func countBattleships(board [][]byte) int {
	count := 0
	for y := range board {
		for x := range board[y] {
			if board[y][x] == '.' {
				continue
			}
			count++
			sinkBattleship(x, y, board)
		}
	}
	return count
}

func sinkBattleship(x, y int, board [][]byte) {
	if y < 0 || x < 0 || y >= len(board) || x >= len(board[0]) || board[y][x] == '.' {
		return
	}
	board[y][x] = '.'
	sinkBattleship(x+1, y, board)
	sinkBattleship(x, y+1, board)
}

func main() {
	ts := []struct {
		input    [][]byte
		expected int
	}{
		{
			input:    [][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}},
			expected: 2,
		},
		{
			input:    [][]byte{{'.'}},
			expected: 0,
		},
	}
	for _, tc := range ts {
		actual := countBattleships(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
