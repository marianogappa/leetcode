package main

import (
	"fmt"
)

// Time: O(n*m) where n is len(commands) and m is len(obstacles)
// Space: O(1)

// Another alternative is to make a map of maps for obstacles, which would make the algorithm much faster
// if there are many obstacles. Considering max movement is 9, it's probably a much better decision.

// Made mistakes:
// (1) Incorrect assumption that going north was dy = -1 (cartesian vs monitor)
// (2) Made silly mistake in rotation directions
// (3) In collision detection assumed that x1 < x2 whereas that depends on dx (for both dimensions)
// (4) Didn't consider the edge case that the robot starts at the position of an obstacle
func robotSim(commands []int, obstacles [][]int) int {
	maxD, x, y, dx, dy := 0, 0, 0, 0, 1
	for _, command := range commands {
		x, y, dx, dy = move(command, x, y, dx, dy, obstacles)
		d := euclidianDistanceSquared(0, 0, x, y)
		if d > maxD {
			maxD = d
		}
	}
	return maxD
}

func euclidianDistanceSquared(x1, y1, x2, y2 int) int {
	return (x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)
}

func move(command, x, y, dx, dy int, obstacles [][]int) (int, int, int, int) {
	switch command {
	case -2:
		dx, dy = rotateCounterClockwise(dx, dy)
	case -1:
		dx, dy = rotateClockwise(dx, dy)
	default:
		x2, y2 := x+command*dx, y+command*dy
		x, y = moveWithCollisions(x, y, x2, y2, dx, dy, obstacles)
	}
	return x, y, dx, dy
}

func rotateCounterClockwise(dx, dy int) (int, int) {
	if dx == 0 && dy == -1 {
		return 1, 0
	} else if dx == 1 && dy == 0 {
		return 0, 1
	} else if dx == 0 && dy == 1 {
		return -1, 0
	}
	return 0, -1
}

func rotateClockwise(dx, dy int) (int, int) {
	if dx == 0 && dy == -1 {
		return -1, 0
	} else if dx == -1 && dy == 0 {
		return 0, 1
	} else if dx == 0 && dy == 1 {
		return 1, 0
	}
	return 0, -1
}

func moveWithCollisions(x1, y1, x2, y2, dx, dy int, obstacles [][]int) (int, int) {
	minEuclidian, minX, minY := euclidianDistanceSquared(x1, y1, x2, y2), x2, y2
	for _, obstacle := range obstacles {
		if x1 == obstacle[0] && y1 == obstacle[1] {
			continue // special case for when the robot starts at an obstacle
		}
		if collides(x1, y1, x2, y2, obstacle[0], obstacle[1]) {
			cx, cy := obstacle[0]+dx*-1, obstacle[1]+dy*-1
			potEuclidian := euclidianDistanceSquared(x1, y1, cx, cy)
			if potEuclidian < minEuclidian {
				minEuclidian, minX, minY = potEuclidian, cx, cy
			}
		}
	}
	return minX, minY
}

func collides(x1, y1, x2, y2, ox, oy int) bool {
	return ((ox >= x1 && ox <= x2) || (ox <= x1 && ox >= x2)) && ((oy >= y1 && oy <= y2) || (oy <= y1 && oy >= y2))
}

func main() {
	ts := []struct {
		commands  []int
		obstacles [][]int
		expected  int
	}{
		{
			commands:  []int{7, -2, -2, 7, 5},
			obstacles: [][]int{{-3, 2}, {-2, 1}, {0, 1}, {-2, 4}, {-1, 0}, {-2, -3}, {0, -3}, {4, 4}, {-3, 3}, {2, 2}},
			expected:  4,
		},
		{
			commands:  []int{-2, -1, -2, 3, 7},
			obstacles: [][]int{{1, -3}, {2, -3}, {4, 0}, {-2, 5}, {-5, 2}, {0, 0}, {4, -4}, {-2, -5}, {-1, -2}, {0, 2}},
			expected:  100,
		},
		{
			commands:  []int{4, -1, 3},
			obstacles: [][]int{},
			expected:  25,
		},
		{
			commands:  []int{4, -1, 4, -2, 4},
			obstacles: [][]int{{2, 4}},
			expected:  65,
		},
		{
			commands:  []int{-2, -1, 8, 9, 6},
			obstacles: [][]int{{-1, 3}, {0, 1}, {-1, 5}, {-2, -4}, {5, 4}, {-2, -3}, {5, -1}, {1, -1}, {5, 5}, {5, 2}},
			expected:  0,
		},
	}
	for _, tc := range ts {
		actual := robotSim(tc.commands, tc.obstacles)
		if tc.expected != actual {
			fmt.Printf("For (%v, %v) expected %v but got %v\n", tc.commands, tc.obstacles, tc.expected, actual)
		}
	}
}
