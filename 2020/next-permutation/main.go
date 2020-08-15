package main

import (
	"fmt"

	"reflect"
)

// Time: O(n)
// Space: O(1)
func nextPermutation(nums []int) {
	var i, j int
	// 1. Find first decreasing number from the right
	for i = len(nums) - 2; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}
	// 2. If number was found, find immediately larger and swap them
	if i >= 0 {
		for j = len(nums) - 1; j >= 0 && nums[j] <= nums[i]; j-- {
		}
		fmt.Println("for", nums, "i is", nums[i], "and j is", nums[j])
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 3. Reverse everything after first decreasing number (or whole array)
	reverse(nums, i+1)
}

func reverse(nums []int, start int) {
	l := start
	r := len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func main() {
	ts := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			input:    []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{1, 1, 5},
			expected: []int{1, 5, 1},
		},
		{
			input:    []int{1, 5, 1},
			expected: []int{5, 1, 1},
		},
		{
			input:    []int{8, 7, 4, 5, 4, 3, 2, 1},
			expected: []int{8, 7, 5, 1, 2, 3, 4, 4},
		},
	}
	for _, tc := range ts {
		nextPermutation(tc.input)
		if !reflect.DeepEqual(tc.expected, tc.input) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, tc.input)
		}
	}
}
