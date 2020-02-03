package main

import "fmt"

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || (len(matrix) == 1 && len(matrix[0]) <= 1) {
		return
	}
	var col0, row0 bool
	for i := range matrix[0] {
		if matrix[0][i] == 0 { // special case: does first row have a 0
			row0 = true
		}
	}
	for i := range matrix {
		if matrix[i][0] == 0 { // special case: does first column have a 0
			col0 = true
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 { // flag encountered zero on first row & first column
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < len(matrix[0]); i++ { // zeroing rows
		if matrix[0][i] == 0 {
			for j := 1; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ { // zeroing columns
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[i]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	if col0 { // zeroing first column
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
	if row0 { //zeroing first row
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
	}
}

func printM(m [][]int) {
	for _, r := range m {
		fmt.Println(r)
	}
	fmt.Println()
}

func main() {
	var ts = []struct {
		i [][]int
		e [][]int
	}{
		{
			i: [][]int{
				{1, 2, 3},
				{4, 0, 6},
				{7, 8, 9},
			},
			e: [][]int{
				{1, 0, 3},
				{0, 0, 0},
				{7, 0, 9},
			},
		},
		{
			i: [][]int{
				{0, 2, 3},
				{4, 5, 6},
				{7, 8, 0},
			},
			e: [][]int{
				{0, 0, 0},
				{0, 5, 0},
				{0, 0, 0},
			},
		},
		{
			i: [][]int{
				{1, 0, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			e: [][]int{
				{0, 0, 0},
				{4, 0, 6},
				{7, 0, 9},
			},
		},
		{
			i: [][]int{},
			e: [][]int{},
		},
		{
			i: [][]int{{1}},
			e: [][]int{{1}},
		},
		{
			i: [][]int{{0}},
			e: [][]int{{0}},
		},

		{
			i: [][]int{
				{0, 0, 0, 5},
				{4, 3, 1, 4},
				{0, 1, 1, 4},
				{1, 2, 1, 3},
				{0, 0, 1, 1},
			},
			e: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 4},
				{0, 0, 0, 0},
				{0, 0, 0, 3},
				{0, 0, 0, 0},
			},
		},
	}
	for _, t := range ts {
		setZeroes(t.i)
		fmt.Println("algorithm did this")
		printM(t.i)
		fmt.Println("should look like")
		printM(t.e)
	}
}
