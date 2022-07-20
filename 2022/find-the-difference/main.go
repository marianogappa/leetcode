package main

import "fmt"

// Time: O(len(s)+len(t))
// Space: O(26*2) = O(1)
//
// Make two arrays with the frequencies of the 26 letters, and compare them.
func findTheDifference(s string, t string) byte {
	var sCounts, tCounts [26]int
	for _, b := range s {
		sCounts[int(b-'a')]++
	}
	for _, b := range t {
		tCounts[int(b-'a')]++
	}
	for i := 0; i < 26; i++ {
		if sCounts[i] != tCounts[i] {
			return byte(int('a') + i)
		}
	}
	// N.B. Unreachable as long as counts are different
	return 'a'
}

func main() {
	ts := []struct {
		s, t     string
		expected byte
	}{
		{
			s:        "abcd",
			t:        "abcde",
			expected: 'e',
		},
		{
			s:        "",
			t:        "y",
			expected: 'y',
		},
	}
	for _, tc := range ts {
		actual := findTheDifference(tc.s, tc.t)
		if tc.expected != actual {
			fmt.Printf("For (%v, %v) expected %v but got %v\n", tc.s, tc.t, tc.expected, actual)
		}
	}
}
