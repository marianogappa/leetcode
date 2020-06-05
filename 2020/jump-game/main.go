package main

import "fmt"

// Time: O(n)
// Space: O(1)
//
// We should be able to jump to any location up to the
// current maximum reachable, so the key is we only need
// to store the _maxReachable_.
//
// As we go through the array, we update _maxReachable_,
// but we must break out of the loop if we "run out of
// fuel".
func canJump(nums []int) bool {
	var maxReachable int
	for i := 0; i < len(nums) && i <= maxReachable; i++ {
		maxReachable = max(maxReachable, i+nums[i])
	}
	return maxReachable >= len(nums)-1
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
		expected bool
	}{
		{
			input:    []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			input:    []int{3, 2, 1, 0, 4},
			expected: false,
		},
	}
	for _, tc := range ts {
		actual := canJump(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
