package main

import (
	"fmt"
	"reflect"
)

var ns = [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

func inBounds(x, y int, board [][]int) bool {
	return x >= 0 && y >= 0 && x < len(board[0]) && y < len(board)
}

func liveAround(x, y int, board [][]int) int {
	var live = 0
	for _, n := range ns {
		var nx, ny = x + n[0], y + n[1]

		if inBounds(nx, ny, board) && board[ny][nx] == 1 {
			live++
		}
	}
	return live
}

func update(x, y int, board [][]int) {
	var l = liveAround(x, y, board)
	if inBounds(x+1, y, board) {
		update(x+1, y, board)
	} else if inBounds(0, y+1, board) {
		update(0, y+1, board)
	}
	switch {
	case board[y][x] == 1 && (l == 2 || l == 3):
	case board[y][x] == 0 && l == 3:
		board[y][x] = 1
	default:
		board[y][x] = 0
	}
}

func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	update(0, 0, board)
}

func printBoard(board [][]int) {
	for _, l := range board {
		fmt.Println(l)
	}
}

func main() {
	var ts = []struct {
		i [][]int
		e [][]int
	}{
		{
			i: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			e: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			i: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			e: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			i: [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{0, 0, 0},
			},
			e: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			i: [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{1, 0, 0},
			},
			e: [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{1, 1, 0},
			},
		},
		{
			i: [][]int{},
			e: [][]int{},
		},
		{
			i: [][]int{{0, 1}},
			e: [][]int{{0, 0}},
		},
	}
	for _, t := range ts {
		// fmt.Println("with initial")
		// printBoard(t.i)
		gameOfLife(t.i)
		if !reflect.DeepEqual(t.e, t.i) {
			fmt.Println("gameOfLife should have been")
			printBoard(t.e)
			fmt.Println("but was")
			printBoard(t.i)
		}
		fmt.Println()
	}
}
