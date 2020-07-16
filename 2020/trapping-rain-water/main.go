package main

import (
	"fmt"
	"math"
)

// Time: O(h)
// Space: O(1)
func trap(height []int) int {
	var (
		water = 0
		l     = 0
		r     = len(height) - 1
		maxL  = math.MinInt32
		maxR  = math.MinInt32
	)
	// With two pointers to the start and end of the
	// array, iterate until they meet, by always
	// advancing the one at smaller height, one pointer
	// at a time.
	for l < r {
		if height[l] < height[r] {
			// Keep a max found from both ends.
			// The trapped rain water cannot exceed the
			// height of the smaller max.
			if height[l] >= maxL {
				maxL = height[l]
			} else {
				// When the height decreases as we advance,
				// there must be trapped rain water.
				water += maxL - height[l]
			}
			l++
		} else {
			if height[r] >= maxR {
				maxR = height[r]
			} else {
				water += maxR - height[r]
			}
			r--
		}
	}

	return water
}

func main() {
	ts := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
	}
	for _, tc := range ts {
		actual := trap(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
