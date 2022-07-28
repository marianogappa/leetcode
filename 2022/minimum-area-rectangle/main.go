package main

import (
	"fmt"
	"math"
)

// Time: O(quadratic to points) as we try all eligible pairs
// Space: O(linear to points), as we need to save all points in a map
//
// Careful with overcomplicating the solution! You have to find the trick and it should be simple.
//
// It's obvious that at some point in the algorithm we're gonna have to know if a point exists in constant time,
// so we might as well put the points in a map!
//
// Trick: pick all pairs of points (n^2) that don't share any coordinates. These are potential edges to calculate
// area with. They can only form rectangles if there are two other points that share a coordinate of one and one of
// the other. That's it! Calculate areas and keep a min.
func minAreaRect(points [][]int) int {
	// Put all points in a map.
	pointMap := map[xy]struct{}{}
	for _, point := range points {
		pointMap[xy{x: point[0], y: point[1]}] = struct{}{}
	}

	minArea := math.MaxInt
	// For all pais of points
	for xy1 := range pointMap {
		for xy2 := range pointMap {
			// If they don't share coordinates
			if xy1.x == xy2.x || xy1.y == xy2.y {
				continue
			}
			// And there are two other points that share one coordinate of one, and one coordinate of the other
			if _, ok1 := pointMap[xy{x: xy1.x, y: xy2.y}]; !ok1 {
				continue
			}
			if _, ok2 := pointMap[xy{x: xy2.x, y: xy1.y}]; !ok2 {
				continue
			}
			// Then we have a rectangle! Calculate area and keep a min.
			area := (max(xy1.x, xy2.x) - min(xy1.x, xy2.x)) * (max(xy1.y, xy2.y) - min(xy1.y, xy2.y))
			minArea = min(minArea, area)
		}
	}

	if minArea == math.MaxInt {
		return 0
	}
	return minArea
}

type xy struct {
	x, y int
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
	ts := []struct {
		input    [][]int
		expected int
	}{
		{
			input:    [][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}},
			expected: 4,
		},
		{
			input:    [][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {4, 1}, {4, 3}},
			expected: 2,
		},
		{
			input:    [][]int{{0, 1}, {1, 3}, {3, 3}, {4, 4}, {1, 4}, {2, 3}, {1, 0}, {3, 4}},
			expected: 2,
		},
	}
	for _, tc := range ts {
		actual := minAreaRect(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
