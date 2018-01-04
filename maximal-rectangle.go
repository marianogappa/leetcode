package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var (
		maxArea = 0
		mx      = len(matrix[0])
		my      = len(matrix)
	)
	for y := range matrix {
		if maxArea >= (my-y)*mx {
			break
		}
		for x := range matrix[0] {
			if maxArea >= (mx-x)*(my-y) {
				break
			}
			if matrix[y][x] == '1' {
				var a = traverse(x, y, x, y, mx-1, my-1, matrix)
				if a > maxArea {
					maxArea = a
				}
			}
		}
	}
	return maxArea
}

func traverse(x0, y0, px, py, mx, my int, matrix [][]byte) int {
	var (
		toX = px
		toY = py
	)
	for i := px + 1; i <= mx; i++ {
		if matrix[py][i] != '1' {
			break
		}
		toX++
	}
	for i := py + 1; i <= my; i++ {
		if matrix[i][px] != '1' {
			break
		}
		toY++
	}
	if toX == px || toY == py || matrix[py+1][px+1] == '0' {
		return max((toX-x0+1)*(py-y0+1), (px-x0+1)*(toY-y0+1))
	}

	return max(max((px-x0+1)*(toY-y0+1), (toX-x0+1)*(py-y0+1)), traverse(x0, y0, px+1, py+1, toX, toY, matrix))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var ts = []struct {
		i [][]byte
		e int
	}{
		{
			i: [][]byte{
				{},
			},
			e: 0,
		},
		{
			i: [][]byte{
				{'1'},
			},
			e: 1,
		},
		{
			i: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			e: 0,
		},
		{
			i: [][]byte{
				{'0', '0', '0'},
				{'0', '1', '1'},
				{'0', '1', '1'},
			},
			e: 4,
		},
		{
			i: [][]byte{
				{'0', '0', '0'},
				{'0', '1', '0'},
				{'0', '0', '0'},
			},
			e: 1,
		},
		{
			i: [][]byte{
				{'1', '0', '0'},
				{'0', '1', '0'},
				{'0', '0', '0'},
			},
			e: 1,
		},
		{
			i: [][]byte{
				{'0', '0', '0'},
				{'0', '1', '1'},
				{'0', '0', '0'},
			},
			e: 2,
		},
		{
			i: [][]byte{
				{'0', '0', '0'},
				{'0', '1', '0'},
				{'0', '1', '0'},
			},
			e: 2,
		},
		{
			i: [][]byte{
				{'0', '0', '0'},
				{'0', '1', '1'},
				{'0', '1', '0'},
			},
			e: 2,
		},
		{
			i: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			e: 9,
		},
		{
			i: [][]byte{
				{'0', '1', '0'},
				{'1', '1', '1'},
				{'0', '1', '0'},
			},
			e: 3,
		},
		{
			i: [][]byte{
				{'0', '1', '0'},
				{'1', '1', '1'},
			},
			e: 3,
		},
		{
			i: [][]byte{
				{'0', '0'},
				{'0', '1'},
				{'0', '1'},
			},
			e: 2,
		},
		{
			i: [][]byte{
				{'1', '1'},
				{'1', '1'},
				{'1', '1'},
			},
			e: 6,
		},
		{
			i: [][]byte{
				{'1', '1', '1', '1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1', '1', '1', '0'},
				{'1', '1', '1', '1', '1', '1', '1', '0'},
				{'1', '1', '1', '1', '1', '0', '0', '0'},
				{'0', '1', '1', '1', '1', '0', '0', '0'},
			},
			e: 21,
		},
	}
	for _, t := range ts {
		var a = maximalRectangle(t.i)
		if t.e != a {
			fmt.Printf("maximalRectangle(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}
