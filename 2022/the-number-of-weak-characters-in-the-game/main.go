package main

import (
	"fmt"
	"sort"
)

// Time: O(nlogn)
// Space: O(1)
//
// You have to be a genius to figure this out, but if you sort by attack descending and secondly by defense ascending,
// then you can go through the sorted properties keeping a maxDefense, and when the current defense is weaker than the
// max, you're currently at a weaker character.
//
// {10, 4}, {4, 3}, {1, 5}
// 01 02 03 04 05 06 07 08 09 10
//       D  A
//          A                  D
// A           D
func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		return properties[i][0] > properties[j][0] || (properties[i][0] == properties[j][0] && properties[i][1] <= properties[j][1])
	})

	var (
		maxDefense int
		weakCount  int
	)
	for _, property := range properties {
		if property[1] < maxDefense {
			weakCount++
		}
		maxDefense = max(maxDefense, property[1])
	}
	return weakCount
}

func max(a, b int) int {
	if a > b {
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
			input:    [][]int{{5, 5}, {6, 3}, {3, 6}},
			expected: 0,
		},
		{
			input:    [][]int{{2, 2}, {3, 3}},
			expected: 1,
		},
		{
			input:    [][]int{{1, 5}, {10, 4}, {4, 3}},
			expected: 1,
		},
	}
	for _, tc := range ts {
		actual := numberOfWeakCharacters(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
