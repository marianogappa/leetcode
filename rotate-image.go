package main

import "fmt"

func rotate(matrix [][]int) {
	for i := range matrix {
		reverse(matrix[i])
	}
	for y := 0; y < len(matrix)-1; y++ {
		diagonalReverse(matrix, 0, y)
	}
	for x := 1; x < len(matrix)-1; x++ {
		diagonalReverse(matrix, x, 0)
	}
}

func reverse(ns []int) {
	var (
		i = 0
		j = len(ns) - 1
	)
	for i < j {
		ns[i], ns[j] = ns[j], ns[i]
		i++
		j--
	}
}

func diagonalReverse(m [][]int, x1, y1 int) {
	var (
		x2, y2 int
	)
	if x1 > y1 {
		x2, y2 = len(m)-1, len(m)-1-(x1-y1)
	} else {
		x2, y2 = len(m)-1-(y1-x1), len(m)-1
	}
	for x1 < x2 && y1 < y2 {
		m[y1][x1], m[y2][x2] = m[y2][x2], m[y1][x1]
		x1++
		y1++
		x2--
		y2--
	}
}

func printM(m [][]int) {
	for _, r := range m {
		fmt.Println(r)
	}
	fmt.Println()
}

func main() {
	var m = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(m)
	printM(m)

	var m2 = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(m2)
	printM(m2)

	var m3 = [][]int{}
	rotate(m3)
	printM(m3)

	var m4 = [][]int{{1}}
	rotate(m4)
	printM(m4)
}
