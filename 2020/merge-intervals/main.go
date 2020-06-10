package main

import (
	"fmt"
	"reflect"
	"sort"
)

// Time: O()
// Space: O()
func merge(intervals [][]int) [][]int {
	// Make sure intervals are sorted
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	newIntervals := [][]int{}
	for i := 0; i < len(intervals); i++ {
		// If this is the first interval, or if the current one doesn't overlap with last, add it
		if len(newIntervals) == 0 || !isOverlap(newIntervals[len(newIntervals)-1], intervals[i]) {
			newIntervals = append(newIntervals, intervals[i])
			continue
		}
		// Otherwise merge current interval with latest!
		newIntervals[len(newIntervals)-1] = mergeIntervals(newIntervals[len(newIntervals)-1], intervals[i])
	}
	return newIntervals
}

func isOverlap(i1, i2 []int) bool {
	if i2[0] < i1[0] {
		i1, i2 = i2, i1
	}
	return i1[0] <= i2[1] && i2[0] <= i1[1]
}

func mergeIntervals(i1, i2 []int) []int {
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
		input    [][]int
		expected [][]int
	}{
		{
			input:    [][]int{{1, 4}, {4, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			input:    [][]int{{4, 5}, {1, 4}},
			expected: [][]int{{1, 5}},
		},
		{
			input:    [][]int{{4, 5}, {4, 5}},
			expected: [][]int{{4, 5}},
		},
		{
			input:    [][]int{{5, 5}, {5, 5}},
			expected: [][]int{{5, 5}},
		},
		{
			input:    [][]int{{4, 5}},
			expected: [][]int{{4, 5}},
		},
		{
			input:    [][]int{},
			expected: [][]int{},
		},
	}
	for _, tc := range ts {
		actual := merge(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
