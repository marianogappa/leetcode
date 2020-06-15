package main

import (
	"fmt"
	"reflect"
)

// Time: O(n)
// Space: O(n)
func isAnagram(s string, t string) bool {
	ms := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		ms[s[i]]++
	}
	mt := make(map[byte]int, len(s))
	for i := 0; i < len(t); i++ {
		mt[t[i]]++
	}
	return reflect.DeepEqual(ms, mt)
}

func main() {
	ts := []struct {
		s, t     string
		expected bool
	}{
		{
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			s:        "rat",
			t:        "car",
			expected: false,
		},
		{
			s:        "",
			t:        "",
			expected: true,
		},
		{
			s:        "a",
			t:        "a",
			expected: true,
		},
		{
			s:        "a",
			t:        "b",
			expected: false,
		},
	}
	for _, tc := range ts {
		actual := isAnagram(tc.s, tc.t)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.s, tc.t, tc.expected, actual)
		}
	}
}
