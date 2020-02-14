package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Strategy is to dfs from all edge cells that have a 'O', marking them
// with a lowercase 'x'. After that's done, all cells with an 'x' are
// the ones that must remain as 'O', and all cells with an 'O' should
// become 'X'. By using the 'x' as memo, the recursion is linear.
//
// Time: O(n) because we use the board as memo to not do any square twice
// Space: O(n) because the board is recycled for state
func solve(board [][]byte) {
	// For all edges, save adjacents to O recursively
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if y != 0 && y != len(board)-1 && x != 0 && x != len(board[y])-1 {
				continue
			}
			dfsSave(x, y, board)
		}
	}

	// Replace all saved cells with O and all unsaved cells with X
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			switch board[y][x] {
			case 'x':
				board[y][x] = 'O'
			default:
				board[y][x] = 'X'
			}
		}
	}
}

type xy struct {
	x, y int
}

func dfsSave(x, y int, board [][]byte) {
	if board[y][x] != 'O' {
		return
	}
	board[y][x] = 'x'
	for _, c := range []xy{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}} {
		if isValid(c.x, c.y, board) {
			dfsSave(c.x, c.y, board)
		}
	}
}

func isValid(x, y int, board [][]byte) bool {
	return x >= 1 && y >= 1 && y <= len(board)-2 && x <= len(board[y])-2
}

func cloneInput(board [][]byte) [][]byte {
	c := make([][]byte, len(board))
	for i := range board {
		c[i] = make([]byte, len(board[i]))
		for j := range board[i] {
			c[i][j] = board[i][j]
		}
	}
	return c
}

func printBoard(board [][]byte) string {
	var sb strings.Builder
	for i := range board {
		sb.WriteString(string(board[i]))
		sb.WriteString("\n")
	}
	sb.WriteString("\n\n")
	return sb.String()
}

func main() {
	ts := []struct {
		input    [][]byte
		expected [][]byte
	}{
		// {
		// 	input: [][]byte{
		// 		{'X', 'X', 'X', 'X'},
		// 		{'X', 'O', 'O', 'X'},
		// 		{'X', 'X', 'O', 'X'},
		// 		{'X', 'O', 'X', 'X'},
		// 	},
		// 	expected: [][]byte{
		// 		{'X', 'X', 'X', 'X'},
		// 		{'X', 'X', 'X', 'X'},
		// 		{'X', 'X', 'X', 'X'},
		// 		{'X', 'O', 'X', 'X'},
		// 	},
		// },
		{
			input: [][]byte{
				{'X', 'O', 'X', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'X'},
			},
			expected: [][]byte{
				{'X', 'O', 'X', 'O', 'X', 'O'},
				{'O', 'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'X'},
			},
		},
	}
	for _, tc := range ts {
		input := cloneInput(tc.input)
		solve(tc.input)
		if !reflect.DeepEqual(tc.expected, tc.input) {
			fmt.Printf("For \n%vexpected \n%vbut got \n%v\n", printBoard(input), printBoard(tc.expected), printBoard(tc.input))
		}
	}
}
