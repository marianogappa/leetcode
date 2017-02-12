package main

import "fmt"

func translate(input [][]int, n int) [][]string {
	output := [][]string{}
	for i := range input {
		o := []string{}
		ii := 0
		j := 0
		for y := 0; y < n; y++ {
			oo := ""
			for x := 0; x < n; x++ {
				if ii < len(input[i]) && j == input[i][ii] {
					oo += "Q"
					ii++
				} else {
					oo += "."
				}
				j++
			}
			o = append(o, oo)
		}
		output = append(output, o)
	}
	return output
}

func place(remaining, n, ly int, grid []int, grids *[][]int, burnedX, burnedDE, burnedDO []bool) [][]int {
	if remaining == 0 {
		*grids = append(*grids, grid)
		return nil
	}

	for y := ly + 1; y < n; y++ {
		for x := 0; x < n; x++ {
			if burnedX[x] {
				continue
			}
			var de, do int
			oy := n - y
			if x <= y {
				de = y - x
			} else {
				de = n + x - y
			}
			if x <= oy {
				do = oy - x
			} else {
				do = n + x - oy
			}
			if burnedDE[de] || burnedDO[do] {
				continue
			}

			newBurnedX, newBurnedDE, newBurnedDO := make([]bool, n), make([]bool, n+n), make([]bool, n+n)
			copy(newBurnedX, burnedX)
			copy(newBurnedDE, burnedDE)
			copy(newBurnedDO, burnedDO)
			newBurnedX[x], newBurnedDE[de], newBurnedDO[do] = true, true, true

			r := make([]int, len(grid)+1)
			copy(r, grid)
			r[len(grid)] = y*n + x

			place(remaining-1, n, y, r, grids, newBurnedX, newBurnedDE, newBurnedDO)
		}
	}
	return *grids
}

func solveNQueens(n int) [][]string {
	return translate(place(n, n, -1, []int{}, &[][]int{}, make([]bool, n), make([]bool, n+n), make([]bool, n+n)), n)
}

func main() {
	n := 9

	for _, grid := range solveNQueens(n) {
		for i := range grid {
			fmt.Printf(grid[i] + "\n")
		}
		fmt.Println()
	}
}
