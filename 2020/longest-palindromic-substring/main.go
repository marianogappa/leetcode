package main

import "fmt"

// Time: O(n^2) expand n times over a string of length n
// Space: O(1)
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var maxL, maxR int
	// Strategy: expand from every possible
	// palindrome center. There are 2*n-1
	// centers in a string.
	//
	// For even numbers, center is a letter.
	// For odd numbers, center is between letters.
	//
	// Pick the longest.
	for i := 0; i < 2*len(s)-1; i++ {
		// Expand over current center and get
		// palindromic range.
		//
		// Careful! Early exit yields l > r!
		l, r := expand(s, i)

		// Keep track of largest palindromic range.
		if l <= r && maxR-maxL < r-l {
			maxL, maxR = l, r
		}
	}
	return s[maxL : maxR+1]
}

func expand(s string, i int) (int, int) {
	var (
		l = i / 2
		r = i / 2
	)
	if i%2 == 1 {
		r++
	}
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l + 1, r - 1
}

func main() {
	ts := []struct {
		input    string
		expected string
	}{
		{
			input:    "babad",
			expected: "bab",
		},
		{
			input:    "cbbd",
			expected: "bb",
		},
		{
			input:    "aaa",
			expected: "aaa",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "a",
			expected: "a",
		},
	}
	for _, tc := range ts {
		actual := longestPalindrome(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
