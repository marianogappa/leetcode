package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	count := 0
	for y := range grid {
		lastX := 0
		for x := range grid[y] {
			if grid[y][x] == 1 {
				count += 4
				if lastX == 1 {
					count -= 2
				}
			}
			lastX = grid[y][x]
		}
	}
	for x := range grid[0] {
		lastY := 0
		for y := range grid {
			if grid[y][x] == 1 && lastY == 1 {
				count -= 2
			}
			lastY = grid[y][x]
		}
	}
	return count
}

func main() {
	fmt.Println(islandPerimeter(
		[][]int{
			[]int{0, 1, 0, 0},
			[]int{1, 1, 1, 0},
			[]int{0, 1, 0, 0},
			[]int{1, 1, 0, 0},
		},
	))
}
