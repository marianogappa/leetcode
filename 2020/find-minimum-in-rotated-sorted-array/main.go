package main

import "fmt"

// Time: O(logn)
// Space: O(1)
func findMin(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return min(nums[0], nums[1])
	default:
		if nums[0] < nums[len(nums)-1] {
			return nums[0]
		}
		mid := len(nums) / 2
		if len(nums)%2 == 0 {
			mid--
		}
		if nums[0] < nums[mid] {
			return findMin(nums[mid+1:])
		}
		// Note that it must be true that
		// nums[0] > nums[mid], because
		// exercise clarifies no dupes.
		return findMin(nums[0 : mid+1])
	}
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
			input:    []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			input:    []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			input:    []int{100, 0, 1, 2, 3, 4},
			expected: 0,
		},
	}
	for _, tc := range ts {
		actual := findMin(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
