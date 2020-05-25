package main

import "fmt"

// Time: O(n)
// Space: O(n)
func characterReplacement(s string, k int) int {
	var (
		start, end, countTopLetter, maxWindow int
		letterFreqs                           = map[byte]int{}
	)

	// Sliding window algorithm
	for end < len(s) {
		// Maintain frequency of letters in the window
		letterFreqs[s[end]]++
		// Keep track of the max frequency
		countTopLetter = max(countTopLetter, letterFreqs[s[end]])

		// Here's the key: window size is "countTopLetter + k, so
		// if by ++ing _end_ our window is larger than that, our
		// window is now invalid and we need to ++ _start_.
		for end-start+1 > countTopLetter+k {
			// As we slide _start_, update frequency of letters
			letterFreqs[s[start]]--
			start++
		}

		// The above loop made sure our window is valid again.
		// We can now keep track of the maxWindow and continue
		// to ++ _end_.
		maxWindow = max(maxWindow, end-start+1)
		end++
	}

	return maxWindow
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ts := []struct {
		inputS   string
		inputK   int
		expected int
	}{
		{
			inputS:   "ABAB",
			inputK:   2,
			expected: 4,
		},
		{
			inputS:   "AABABBA",
			inputK:   1,
			expected: 4,
		},
	}
	for _, tc := range ts {
		actual := characterReplacement(tc.inputS, tc.inputK)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.inputS, tc.inputK, tc.expected, actual)
		}
	}
}
