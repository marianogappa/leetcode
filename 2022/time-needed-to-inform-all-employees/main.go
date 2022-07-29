package main

import "fmt"

// Time: O(n)
// Space: O(n)
//
// Very straightforward exercise. Linear time/space step to turn the manager slice into a tree, and then traverse the
// tree to find the largest branch edge sum (also linear time, and linear space in the worst case of a very top-down
// hierarchy).
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	tree := make([][]int, n)
	for child, parent := range manager {
		if child == headID {
			continue
		}
		tree[parent] = append(tree[parent], child)
	}

	return informTime[headID] + maxMinutes(headID, tree, informTime)
}

func maxMinutes(fromID int, tree [][]int, informTime []int) int {
	mx := 0
	for _, childID := range tree[fromID] {
		mx = max(mx, informTime[childID]+maxMinutes(childID, tree, informTime))
	}
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ts := []struct {
		n          int
		headID     int
		manager    []int
		informTime []int
		expected   int
	}{
		{
			n:          1,
			headID:     0,
			manager:    []int{-1},
			informTime: []int{0},
			expected:   0,
		},
		{
			n:          6,
			headID:     2,
			manager:    []int{2, 2, -1, 2, 2, 2},
			informTime: []int{0, 0, 1, 0, 0, 0},
			expected:   1,
		},
	}
	for _, tc := range ts {
		actual := numOfMinutes(tc.n, tc.headID, tc.manager, tc.informTime)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v,%v,%v) expected %v but got %v\n", tc.n, tc.headID, tc.manager, tc.informTime, tc.expected, actual)
		}
	}
}
