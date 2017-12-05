package main

import "fmt"

func trap(height []int) int {
	return doTrap(height, 0, len(height)-1)
}

func doTrap(height []int, from, to int) int {
	if from < 0 || to >= len(height) {
		return 0
	}
	var max1, max1i, max2, max2i int
	for i := from; i <= to; i++ {
		if height[i] > max1 {
			max1, max1i, max2, max2i = height[i], i, max1, max1i
		} else if height[i] > max2 {
			max2, max2i = height[i], i
		}
	}
	if max1 == 0 || max2 == 0 {
		return 0
	}
	var (
		minV = min(max1, max2)
		minI = min(max1i, max2i)
		maxI = max(max1i, max2i)
	)
	return minV*(maxI-minI-1) - sum(height, minI+1, maxI-1) + doTrap(height, from, minI) + doTrap(height, maxI, to)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum(height []int, from, to int) int {
	acc := 0
	for i := from; i <= to; i++ {
		acc += height[i]
	}
	return acc
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
