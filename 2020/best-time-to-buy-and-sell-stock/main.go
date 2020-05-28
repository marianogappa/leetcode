package main

import (
	"fmt"
	"math"
)

// Time: O(n)
// Space: O(1)
func maxProfit(prices []int) int {
	var (
		mn = math.MaxInt32
		mx = 0
	)
	for _, price := range prices {
		mx = max(mx, price-mn)
		mn = min(mn, price)
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
			input:    []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			input:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			input:    []int{2, 1, 2, 1, 0, 1, 2},
			expected: 2,
		},
	}
	for _, tc := range ts {
		actual := maxProfit(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
