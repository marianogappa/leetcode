package main

import (
	"fmt"
	"math"
)

// Time: O(n)
// Space: O(1)
//
// Only a small variation from Kadane's algorithm:
// https://en.wikipedia.org/wiki/Maximum_subarray_problem
//
// If you keep a rolling minimum_price & maximum_profit greedily, you
// cannot escape finding the max profit.
func maxProfit(prices []int) int {
	var (
		minPrice  = math.MaxInt32
		maxProfit = 0
	)
	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxProfit = max(maxProfit, price-minPrice)
	}
	return maxProfit
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
	}
	for _, tc := range ts {
		actual := 1
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
