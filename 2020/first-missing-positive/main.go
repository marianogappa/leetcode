package main

import "fmt"

// Time: O(n)
// Space: O(1)
func firstMissingPositive(nums []int) int {
	// Remove negatives and zeroes.
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[i] = len(nums) + 1
		}
	}

	// Mark indices of used nums.
	for _, num := range nums {
		absNum := abs(num)
		if absNum > len(nums) {
			continue
		}
		absIndex := absNum - 1

		// Marking it by making it negative
		// Could be marking it twice, so
		// have to do -abs().
		nums[absIndex] = -abs(nums[absIndex])
	}

	// Pick first unmark index
	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	ts := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 0},
			expected: 3,
		},
		{
			input:    []int{3, 4, -1, 1},
			expected: 2,
		},
		{
			input:    []int{7, 8, 9, 11, 12},
			expected: 1,
		},
		{
			input:    []int{},
			expected: 1,
		},
		{
			input:    []int{1},
			expected: 2,
		},
		{
			input:    []int{2},
			expected: 1,
		},
	}
	for _, tc := range ts {
		actual := firstMissingPositive(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
