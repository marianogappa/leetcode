package main

import (
	"fmt"
	"reflect"
)

type status int

const (
	unmerged status = iota
	merging
	merged
)

// Time: O(n)
// Space: O(n)
func insert(intervals [][]int, newInterval []int) [][]int {
	newIntervals := [][]int{}
	var status = unmerged
	for i := 0; i < len(intervals); i++ {
		// If there's overlap with current interval, merge until there isn't.
		if status != merged && unsortedIsOverlap(intervals[i], newInterval) {
			newInterval = merge(intervals[i], newInterval)
			status = merging
			continue
		}
		// If there's no longer overlap, add merged interval to final list.
		// Also if _newInterval_ had no overlap and is to the left of _intervals[i][0]_!
		if status == merging || (status == unmerged && newInterval[1] < intervals[i][0]) {
			newIntervals = append(newIntervals, newInterval)
			status = merged
		}
		// Always add current interval (unless we're merging it).
		newIntervals = append(newIntervals, intervals[i])
	}
	// At the end, if we never added _newInterval_ in, do it now!
	if status != merged {
		newIntervals = append(newIntervals, newInterval)
	}
	return newIntervals
}

func unsortedIsOverlap(i1, i2 []int) bool {
	if i1[0] <= i2[0] {
		return isOverlap(i1, i2)
	}
	return isOverlap(i2, i1)
}

func isOverlap(i1, i2 []int) bool {
	return i1[0] <= i2[1] && i2[0] <= i1[1]
}

func merge(i1, i2 []int) []int {
	return []int{min(i1[0], i2[0]), max(i1[1], i2[1])}
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
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			expected:    [][]int{{1, 5}, {6, 9}},
		},
		{
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			expected:    [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			intervals:   [][]int{},
			newInterval: []int{5, 7},
			expected:    [][]int{{5, 7}},
		},
		{
			intervals:   [][]int{{1, 5}},
			newInterval: []int{6, 8},
			expected:    [][]int{{1, 5}, {6, 8}},
		},
		{
			intervals:   [][]int{{1, 5}},
			newInterval: []int{0, 0},
			expected:    [][]int{{0, 0}, {1, 5}},
		},
	}
	for _, tc := range ts {
		actual := insert(tc.intervals, tc.newInterval)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.intervals, tc.newInterval, tc.expected, actual)
		}
	}
}
