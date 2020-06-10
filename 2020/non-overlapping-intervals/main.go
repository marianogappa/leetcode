package main

import (
	"fmt"
	"sort"
)

// Time: O()
// Space: O()
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	var (
		count = 1
		end   = intervals[0][1]
	)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			count++
			end = intervals[i][1]
		}
	}
	return len(intervals) - count
}

func main() {
	ts := []struct {
		input    [][]int
		expected int
	}{
		{
			input:    [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			expected: 1,
		},
		{
			input:    [][]int{{1, 2}, {1, 2}, {1, 2}},
			expected: 2,
		},
		{
			input:    [][]int{{1, 2}, {2, 3}},
			expected: 0,
		},
	}
	for _, tc := range ts {
		actual := eraseOverlapIntervals(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
