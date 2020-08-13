package main

import "fmt"

// Time: O((9!)^9) => O(1)
// Space: 81*3 => O(1)
func solveSudoku(board [][]byte) {
	ys := make([][]byte, 9)
	xs := make([][]byte, 9)
	bs := make([][]byte, 9)
	for i := range ys {
		ys[i] = make([]byte, 10)
		xs[i] = make([]byte, 10)
		bs[i] = make([]byte, 10)
	}

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if board[y][x] == '.' {
				continue
			}
			block := y/3*3 + x/3
			ys[y][board[y][x]-'0'] = 1
			xs[x][board[y][x]-'0'] = 1
			bs[block][board[y][x]-'0'] = 1
		}
	}

	doSolveSudoku(board, ys, xs, bs)
}

func doSolveSudoku(board [][]byte, ys, xs, bs [][]byte) bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if board[y][x] != '.' {
				continue
			}
			block := y/3*3 + x/3
			for n := 1; n <= 9; n++ {
				if ys[y][n]+xs[x][n]+bs[block][n] > 0 {
					continue
				}
				board[y][x] = '0' + byte(n)
				ys[y][n] = 1
				xs[x][n] = 1
				bs[block][n] = 1
				if doSolveSudoku(board, ys, xs, bs) {
					return true
				}
				board[y][x] = '.'
				ys[y][n] = 0
				xs[x][n] = 0
				bs[block][n] = 0
			}
			return false
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'.', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '.', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '.', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '.', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '.', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '.', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '.', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '.', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '.'},
	}
	solveSudoku(board)
	for y := range board {
		for x := range board[y] {
			fmt.Print(string(board[y][x]), " ")
		}
		fmt.Println()
	}
}

/*
	{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
	{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
	{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
	{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
	{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
	{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
	{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
	{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
	{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
*/
