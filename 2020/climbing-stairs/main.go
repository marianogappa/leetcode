package main

import "fmt"

// Time: O(n)
// Space: O(n)
func climbStairs(n int) int {
	return doClimbStairs(n, map[int]int{0: 0, 1: 1, 2: 2})
}

func doClimbStairs(n int, memo map[int]int) int {
	if _, ok := memo[n]; !ok {
		memo[n] = doClimbStairs(n-1, memo) + doClimbStairs(n-2, memo)
	}
	return memo[n]
}

func main() {
	ts := []struct {
		input    int
		expected int
	}{
		{
			input:    2,
			expected: 2,
		},
		{
			input:    3,
			expected: 3,
		},
		{
			input:    4,
			expected: 5,
		},
	}
	for _, tc := range ts {
		actual := climbStairs(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
