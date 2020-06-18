package main

import "fmt"

// Time: O(n)
// Space: O(1)
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	var (
		l = 0
		r = len(s) - 1
	)
	for l < r {
		// Advance _l_ until it reaches the next letter
		sl := s[l]
		if s[l] >= 'A' && s[l] <= 'Z' {
			sl = s[l] - 'A' + 'a'
		}
		if !((sl >= 'a' && sl <= 'z') || (sl >= '0' && sl <= '9')) {
			l++
			continue
		}

		// Advance _r_ until it reaches the next letter
		sr := s[r]
		if s[r] >= 'A' && s[r] <= 'Z' {
			sr = s[r] - 'A' + 'a'
		}
		if !((sr >= 'a' && sr <= 'z') || (sr >= '0' && sr <= '9')) {
			r--
			continue
		}

		// Bail if letters don't match and advance
		if sl != sr {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	ts := []struct {
		input    string
		expected bool
	}{
		{
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			input:    "race a car",
			expected: false,
		},
		{
			input:    "",
			expected: true,
		},
		{
			input:    " ",
			expected: true,
		},
		{
			input:    "a",
			expected: true,
		},
		{
			input:    "A",
			expected: true,
		},
		{
			input:    "0P",
			expected: false,
		},
	}
	for _, tc := range ts {
		actual := isPalindrome(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
