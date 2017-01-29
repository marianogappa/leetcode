package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	ic, x, lx, ly := 0, 0, len(grid), len(grid[0])
	for x < lx {
		y := 0
		for y < ly {
			if grid[x][y] == 49 {
				burn(x, y, &grid)
				ic++
				y++
			}
			y++
		}
		x++
	}

	return ic
}

func burn(x, y int, grid *[][]byte) {
	if x < 0 || y < 0 || x >= len(*grid) || y >= len((*grid)[0]) {
		return
	}
	if (*grid)[x][y] == 49 {
		(*grid)[x][y] = 50
		burn(x+1, y, grid)
		burn(x-1, y, grid)
		burn(x, y+1, grid)
		burn(x, y-1, grid)
	}
}

func main() {
	fmt.Println("6 ==", numIslands([][]byte{
		[]byte("110001"),
		[]byte("000001"),
		[]byte("100100"),
		[]byte("000000"),
		[]byte("011001"),
	}))
	fmt.Println("4 ==", numIslands([][]byte{
		[]byte("010101"),
		[]byte("110111"),
		[]byte("000000"),
		[]byte("010001"),
		[]byte("011011"),
	}))
	fmt.Println("5 ==", numIslands([][]byte{
		[]byte("100000"),
		[]byte("010000"),
		[]byte("001000"),
		[]byte("000100"),
		[]byte("000010"),
	}))
	fmt.Println("1 ==", numIslands([][]byte{
		[]byte("11110"),
		[]byte("11010"),
		[]byte("11000"),
		[]byte("00000"),
	}))
	fmt.Println("3 ==", numIslands([][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}))
}
