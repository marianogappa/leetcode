package main

import "fmt"

// Time: O(n)
// Space: O(1)
//
// We always buy+sell when we will make a profit, and not when we will lose. Note that, in terms of calculating
// profit, buying at first timestamp and selling at last is exactly the same as buying and selling at every pair.
// So this should be as trivial as computing an array of differences between subsequent numbers, & ∑ the positive ones.
//
// Strategy looks like it works well on the examples:
//
// [7, 1, 5, 3, 6, 4] becomes [-6, 4, -2, 3, -2]: 4+3 = 7
// [1, 2, 3, 4, 5] becomes [1, 1, 1, 1]: 1+1+1+1 = 4
// [7, 6, 4, 3, 1] becomes [-1,-2,-1,-2]: 0
func maxProfit(prices []int) int {
	profit := 0
	// Note that len(prices) is guaranteed to have at least one entry by constraint.
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			profit += diff
		}
	}
	return profit
}

func main() {
	ts := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			input:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}
	for _, tc := range ts {
		actual := maxProfit(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
