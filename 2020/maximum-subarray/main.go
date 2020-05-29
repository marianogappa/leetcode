package main

import (
	"fmt"
	"math"
)

// Time: O(n)
// Space: O(1)
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		mx            = math.MinInt32
		cumulativeSum int
	)
	for _, num := range nums {
		// Calculate cumulative sum, unless
		// current number is greater than cum.
		// sum, in which case start over with
		// the current number.
		cumulativeSum += num
		if cumulativeSum < num {
			cumulativeSum = num
		}
		// On every iteration, update max
		mx = max(mx, cumulativeSum)
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
		input    []int
		expected int
	}{
		{
			input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			input:    []int{-1},
			expected: -1,
		},
	}
	for _, tc := range ts {
		actual := maxSubArray(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
