package main

import (
	"fmt"
)

// Time: O(s + o) loop over order and s
// Space: O(o + s) worst case: every char in s is in extra
//
// Just use counting sort! But remember that there are extra characters.
func customSortString(order string, s string) string {
	// Map of character to idx in order, to have constant lookups.
	byteToOrder := map[byte]int{}
	for i := 0; i < len(order); i++ {
		byteToOrder[order[i]] = i
	}

	// Iterate over s, and count occurrence of every character in order.
	// Also store characters not present in order somewhere.
	counts := make([]int, len(order))
	extra := []byte{}
	for i := 0; i < len(s); i++ {
		if _, ok := byteToOrder[s[i]]; !ok {
			extra = append(extra, s[i])
			continue
		}
		counts[byteToOrder[s[i]]]++
	}

	// Construct the result string by appending each character in "order"
	// as many times as "counts[i]" says we found them.
	result := []byte{}
	for i := 0; i < len(order); i++ {
		for j := 0; j < counts[i]; j++ {
			result = append(result, order[i])
		}
	}

	// Don't forget to add the extra characters not in order.
	result = append(result, extra...)
	return string(result)
}

func main() {
	ts := []struct {
		order    string
		s        string
		expected string
	}{
		{
			order:    "cba",
			s:        "abcd",
			expected: "cbad",
		},
		{
			order:    "cbafg",
			s:        "abcd",
			expected: "cbad",
		},
		{
			order:    "exv",
			s:        "xwvee",
			expected: "eexvw",
		},
	}
	for _, tc := range ts {
		actual := customSortString(tc.order, tc.s)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.order, tc.s, tc.expected, actual)
		}
	}
}
