package main

import "fmt"

// Time: O(n) Space: O(n)
func SpiralCopy(inputMatrix [][]int) []int {
	var (
		res      = make([]int, len(inputMatrix)*len(inputMatrix[0]))
		deltaMap = map[int][]int{
			0: []int{1, 0},
			1: []int{0, 1},
			2: []int{-1, 0},
			3: []int{0, -1},
		}
		deltaLimits = map[int][]int{
			0: []int{len(inputMatrix[0]), -1},
			1: []int{-1, len(inputMatrix)},
			2: []int{-1, -1},
			3: []int{-1, -1},
		}
		di, resI int
		x, y     int
	)

	// Time: O(n) Space: O(n)
	for resI < len(res) {
		res[resI] = inputMatrix[y][x]
		resI++
		x += deltaMap[di][0]
		y += deltaMap[di][1]
		if x == deltaLimits[di][0] || y == deltaLimits[di][1] {
			x -= deltaMap[di][0]
			y -= deltaMap[di][1]
			deltaLimits[di][0] -= deltaMap[di][0]
			deltaLimits[di][1] -= deltaMap[di][1]
			di++
			if di == len(deltaMap) {
				di = 0
			}
			x += deltaMap[di][0]
			y += deltaMap[di][1]
		}
	}

	return res
}

func main() {
	fmt.Println(SpiralCopy(
		[][]int{
			[]int{1, 2, 3, 4, 5},
			[]int{6, 7, 8, 9, 10},
			[]int{11, 12, 13, 14, 15},
			[]int{16, 17, 18, 19, 20},
		},
	))
}
