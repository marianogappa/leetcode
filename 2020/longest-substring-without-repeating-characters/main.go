package main

import "fmt"

// Time: O(n)
// Space: O(n)
func lengthOfLongestSubstring(s string) int {
	var (
		start, maxCount int
		letters         = make(map[byte]struct{}, len(s))
	)
	// Advance until we hit a duplicate letter.
	// Keep track of the letters found.
	for end := 0; end < len(s); end++ {
		// If a duplicate letter is found...
		if _, ok := letters[s[end]]; ok {
			// Advance _start_ until we remove the original
			// instance of the duplicate letter.
			for ; s[start] != s[end]; start++ {
				// The letters we pass on the way should no
				// longer be on the list of letters found.
				delete(letters, s[start])
			}
			start++ // Because _start_ still in a duplicate letter
		}
		letters[s[end]] = struct{}{}
		maxCount = max(maxCount, end-start+1)
	}
	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ts := []struct {
		input    string
		expected int
	}{
		{
			input:    "abcabcbb",
			expected: 3,
		},
		{
			input:    "bbbbb",
			expected: 1,
		},
		{
			input:    "pwwkew",
			expected: 3,
		},
	}
	for _, tc := range ts {
		actual := lengthOfLongestSubstring(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
