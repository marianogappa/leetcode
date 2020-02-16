package main

import "fmt"

// This is a tricky number base system, because it looks like
// a base-27 but there's no zero, so it's base-26. Remember
// that A = 1. So the rightmost digit is (26 ^ 0) * digit,
// the second (26 ^ 1) * digit and so on, but digit is never
// zero.
//
// Time: O(n) where n is len(s)
// Space: O(1)
func titleToNumber(s string) int {
	var (
		total int
		pow   = 1
	)
	for i := len(s) - 1; i >= 0; i-- {
		total += pow * (int(s[i]-'A') + 1)
		pow *= 26
	}
	return total
}

func main() {
	ts := []struct {
		input    string
		expected int
	}{
		{
			input:    "A",
			expected: 1,
		},
		{
			input:    "B",
			expected: 2,
		},
		{
			input:    "AA",
			expected: 27,
		},
		{
			input:    "AB",
			expected: 28,
		},
		{
			input:    "ZY",
			expected: 701,
		},
		{
			input:    "ZZ",
			expected: 702,
		},
		{
			input:    "AAA",
			expected: 703,
		},
	}
	for _, tc := range ts {
		actual := titleToNumber(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
