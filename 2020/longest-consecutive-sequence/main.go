package main

import "fmt"

// Time: O(n)
// Space: O(n)
func longestConsecutive(nums []int) int {
	// Build a set of _nums_ for constant access.
	set := make(map[int]bool, len(nums))
	for _, num := range nums {
		set[num] = true
	}

	var longest int
	for i, num := range nums {
		// We only want to explore numbers that start sequences.
		if set[num-1] {
			continue
		}

		// Check sequentially if current number plus 1 exists.
		var length int
		for length = 1; set[nums[i]+length]; length++ {
		}

		// Update _longest_.
		if length > longest {
			longest = length
		}
	}

	return longest
}

func main() {
	ts := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
	}
	for _, tc := range ts {
		actual := longestConsecutive(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
