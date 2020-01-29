package main

import "fmt"

// Strategy is to start from both edges of the map and narrow
// that range until all interesting areas are covered, linearly
// but without losing any interesting areas.
//
// The trick to keep it linear is to move the "cursors" one at a
// time, and always moving the smallest first. Do the exercise
// manually on a whiteboard to check, but this way no interesting
// areas are lost. Otherwise, the time complexity is quadratic.
//
// Time: O(n)
// Space: O(1)
func maxArea(height []int) int {
	var (
		start = 0
		end   = len(height) - 1
		mx    int
	)
	for start < end {
		mx = max(mx, min(height[start], height[end])*(end-start))
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
}
