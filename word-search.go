package main

import "fmt"

func exist(board [][]byte, word string) bool {
	if word == "" {
		return true
	}
	for y := range board {
		for x := range board[0] {
			if doExist(board, word, x, y) {
				return true
			}
		}
	}
	return false
}

func doExist(board [][]byte, word string, x, y int) bool {
	if word == "" {
		return true
	}
	if x < 0 || y < 0 || x >= len(board[0]) || y >= len(board) || board[y][x] != word[0] {
		return false
	}
	var c = board[y][x]
	board[y][x] = 0
	if !(doExist(board, word[1:], x-1, y) ||
		doExist(board, word[1:], x+1, y) ||
		doExist(board, word[1:], x, y-1) ||
		doExist(board, word[1:], x, y+1)) {
		board[y][x] = c
		return false
	}
	return true
}

func main() {
	var ts = []struct {
		board [][]byte
		word  string
		e     bool
	}{
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCCED",
			e:    true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "SEE",
			e:    true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCB",
			e:    false,
		},
		{
			board: [][]byte{
				{'A', 'C'},
				{'B', 'D'},
			},
			word: "ABCD",
			e:    false,
		},
		{
			board: [][]byte{
				{'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A'},
			},
			word: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			e:    false,
		},
	}
	for _, t := range ts {
		var a = exist(t.board, t.word)
		if t.e != a {
			fmt.Println("Incorrect result for %v", t.word)
		} else {
			fmt.Println(t.word, "was ok")
		}
	}
}
