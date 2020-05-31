package main

import "fmt"

// Time: O(n)
// Space: O(1)
func maxArea(height []int) int {
	var (
		l  = 0
		r  = len(height) - 1
		mx = 0
	)
	for l < r {
		mx = max(mx, min(height[l], height[r])*(r-l))
		// This is the key: as long as we move one
		// pointer at a time, and we move the one with
		// the lower height, we won't lose any
		// interesting areas.
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return mx
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

func main() {
	ts := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
	}
	for _, tc := range ts {
		actual := maxArea(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
