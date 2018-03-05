package main

import "fmt"

func multiply(a, b [][]int) [][]int {
	if len(a) == 0 || len(b) == 0 || len(a[0]) == 0 || len(b[0]) == 0 {
		return [][]int{}
	}
	// build sparse versions of a & b: []int{x, y, val}
	// key: a is sorted by y <- not really; doesn't matter
	var (
		sa = make([][]int, 0)
		sb = make([][]int, 0)
	)
	for y := range a {
		for x := range a[y] {
			if a[y][x] != 0 {
				sa = append(sa, []int{x, y, a[y][x]})
			}
		}
	}
	for x := range b[0] {
		for y := range b {
			if b[y][x] != 0 {
				sb = append(sb, []int{x, y, b[y][x]})
			}
		}
	}
	// initialize result matrix
	var res = make([][]int, len(a), len(a))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(b[0]))
	}
	// fill result matrix of cardinality y(a)*x(b)
	for _, a := range sa {
		for _, b := range sb {
			if a[0] == b[1] {
				res[a[1]][b[0]] += a[2] * b[2]
			}
		}
	}
	return res
}

func main() {
	// var res = multiply(
	// 	[][]int{
	// 		{1, -5},
	// 	},
	// 	[][]int{
	// 		{12},
	// 		{-1},
	// 	},
	// )
	// for _, r := range res {
	// 	fmt.Println(r)
	// }
	var res = multiply(
		[][]int{
			{1, 0, 0},
			{-1, 0, 3},
		},
		[][]int{
			{7, 0, 0},
			{0, 0, 0},
			{0, 0, 1},
		},
	)
	for _, r := range res {
		fmt.Println(r)
	}
}
