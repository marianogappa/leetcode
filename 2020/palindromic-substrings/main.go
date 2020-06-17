package main

import "fmt"

// Time: O(n^2)
// Space: O(1)
func countSubstrings(s string) int {
	var count int
	// For all possible string centers,
	// expand outwards while valid
	// palindrome (count++ for each).
	//
	// If _i_ is even, center is letter.
	// If _i_ is odd, center is between
	// letters.
	for i := 0; i < 2*len(s)-1; i++ {
		count += expansionCount(s, i)
	}
	return count
}

func expansionCount(s string, i int) int {
	var (
		count int
		l     = i / 2
		r     = i / 2
	)
	if i%2 == 1 {
		r++
	}
	for l >= 0 && r < len(s) && s[l] == s[r] {
		count++
		l--
		r++
	}
	return count
}

func main() {
	ts := []struct {
		input    string
		expected int
	}{
		{
			input:    "abc",
			expected: 3,
		},
		{
			input:    "aaa",
			expected: 6,
		},
		{
			input:    "aba",
			expected: 4,
		},
		{
			input:    "",
			expected: 0,
		},
		{
			input:    "a",
			expected: 1,
		},
		{
			input:    "aa",
			expected: 3,
		},
	}
	for _, tc := range ts {
		actual := countSubstrings(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
