package main

import "fmt"

// Time: O(log n)
// Space: O(1)
func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	// Finding the integer square root is
	// finding a number that squared equals
	// the closest integer square root.
	//
	// That number naively is between 1 and x.

	// We can further refine that between 1
	// and x/2.
	upper := x / 2
	lower := 1

	// We should do binary search to avoid
	// linear time complexity.
	for lower <= upper {
		mid := (upper + lower) / 2
		pow := mid * mid
		switch {
		case pow == x:
			return mid
		case pow > x:
			upper = mid - 1
		case pow < x:
			lower = mid + 1
		}
	}
	// If the loop is no longer true, it means
	// upper < lower. We want the smaller of the
	// two.
	return upper
}

func main() {
	ts := []struct {
		input    int
		expected int
	}{
		{
			input:    1,
			expected: 1,
		},
		{
			input:    0,
			expected: 0,
		},
		{
			input:    2,
			expected: 1,
		},
		{
			input:    4,
			expected: 2,
		},
		{
			input:    8,
			expected: 2,
		},
		{
			input:    9,
			expected: 3,
		},
		{
			input:    25,
			expected: 5,
		},
	}
	for _, tc := range ts {
		actual := mySqrt(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
