package main

import (
	"fmt"
	"strings"
)

// Time: O(n) where n is len(str)
// Space: O(n)
func wordPattern(pattern string, str string) bool {
	if len(pattern) > 0 && len(str) == 0 {
		return false
	}
	words := strings.Split(strings.TrimSpace(str), " ")
	if len(words) != len(pattern) {
		return false
	}
	ps := map[byte]string{}
	for i, s := range words {
		if ms, ok := ps[pattern[i]]; ok && s != ms {
			return false
		}
		ps[pattern[i]] = s
	}
	uniqueWords := map[string]struct{}{}
	for _, w := range ps {
		uniqueWords[w] = struct{}{}
	}
	return len(uniqueWords) == len(ps)
}

func main() {
	ts := []struct {
		pattern  string
		str      string
		expected bool
	}{
		{
			pattern:  "abba",
			str:      "dog cat cat dog",
			expected: true,
		},
		{
			pattern:  "abba",
			str:      "dog cat cat fish",
			expected: false,
		},
		{
			pattern:  "aaaa",
			str:      "dog cat cat dog",
			expected: false,
		},
		{
			pattern:  "abba",
			str:      "dog dog dog dog",
			expected: false,
		},
		{
			pattern:  "",
			str:      "",
			expected: false,
		},
		{
			pattern:  "",
			str:      "a",
			expected: false,
		},
		{
			pattern:  "a",
			str:      "",
			expected: false,
		},
	}
	for _, tc := range ts {
		actual := wordPattern(tc.pattern, tc.str)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.pattern, tc.str, tc.expected, actual)
		}
	}
}
