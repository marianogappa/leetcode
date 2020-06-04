package main

import "fmt"

// Time: O(n) due to memoization
// Space: O(n) memo being size of nums at most
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// The added difficulty with respect to House Robber is
	// just that if you robbed nums[0] you can't rob
	// nums[len(nums)-1], because they are adjacent.
	//
	// That difficulty crucially affects the memo!
	//
	// The solution is this line, which acts as a special
	// case for when the first house is robbed or not, and
	// takes the larger (and builds separate memos for each).
	return max(nums[0]+doRob(nums[:len(nums)-1], 2, map[int]int{}), doRob(nums[1:], 0, map[int]int{}))
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
			input:    []int{2, 3, 2},
			expected: 3,
		},
		{
			input:    []int{},
			expected: 0,
		},
		{
			input:    []int{1},
			expected: 1,
		},
		{
			input:    []int{1, 3, 1, 3, 100},
			expected: 103,
		},
		{
			input:    []int{94, 40, 49, 65, 21, 21, 106, 80, 92, 81, 679, 4, 61, 6, 237, 12, 72, 74, 29, 95, 265, 35, 47, 1, 61, 397, 52, 72, 37, 51, 1, 81, 45, 435, 7, 36, 57, 86, 81, 72},
			expected: 2926,
		},
	}
	for _, tc := range ts {
		actual := rob(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
