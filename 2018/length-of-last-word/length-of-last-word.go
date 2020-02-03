package main

import "fmt"

// Time: O(n)
// Space: O(1)
// In hindsight, it'd be much more efficient to look for first space from the right side after strings.TrimSpace()!
func lengthOfLastWord(s string) int {
	length := 0
	lastLength := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			if length > 0 {
				lastLength, length = length, 0
			}
		default:
			length++
		}
	}
	if length > 0 {
		return length
	}
	return lastLength
}

func main() {
	ts := []struct {
		input    string
		expected int
	}{
		{
			input:    "Hello World",
			expected: 5,
		},
		{
			input:    "",
			expected: 0,
		},
		{
			input:    "   ",
			expected: 0,
		},
		{
			input:    " a  ",
			expected: 1,
		},
		{
			input:    " a  aa",
			expected: 2,
		},
		{
			input:    " a a aa",
			expected: 2,
		},
		{
			input:    " a aa a",
			expected: 1,
		},
		{
			input:    " a aa a ",
			expected: 1,
		},
		{
			input:    "aaaa",
			expected: 4,
		},
	}
	for _, tc := range ts {
		actual := lengthOfLastWord(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
