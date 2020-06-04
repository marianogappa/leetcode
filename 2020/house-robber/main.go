package main

import "fmt"

// Time: O(n) due to memoization
// Space: O(n) memo being size of nums at most
func rob(nums []int) int {
	return doRob(nums, 0, map[int]int{})
}

func doRob(nums []int, i int, memo map[int]int) int {
	if i >= len(nums) {
		return 0
	}
	if v, ok := memo[i]; ok {
		return v
	}
	memo[i] = max(nums[i]+doRob(nums, i+2, memo), doRob(nums, i+1, memo))
	return memo[i]
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
			input:    []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			input:    []int{2, 7, 9, 3, 1},
			expected: 12,
		},
		{
			input:    []int{},
			expected: 0,
		},
		{
			input:    []int{1},
			expected: 1,
		},
	}
	for _, tc := range ts {
		actual := rob(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
