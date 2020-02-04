package main

import "fmt"

// We only need to keep track of the maximum reachable step, because all
// the ones before by definition can be reached. As we walk towardd the
// max reachable, some steps may move max reachable further. If max
// reachable moves up to or exceeding the last step, then result is true.
//
// Time: O(n) because you only go through the nums once
// Space: O(1)
func canJump(nums []int) bool {
	var maxReached int
	for i := 0; i <= maxReached && maxReached < len(nums)-1; i++ {
		maxReached = max(maxReached, i+nums[i])
	}
	return maxReached >= len(nums)-1
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
