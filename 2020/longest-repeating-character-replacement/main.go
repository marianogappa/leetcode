package main

import "fmt"

// Time: O(n)
// Space: O(n)
func characterReplacement(s string, k int) int {
	var (
		start, end, countTopLetter, maxWindow int
		letterFreqs                           = map[byte]int{}
	)

	for end < len(s) {
		letterFreqs[s[end]]++
		countTopLetter = max(countTopLetter, letterFreqs[s[end]])
		for end-start+1-countTopLetter > k {
			letterFreqs[s[start]]--
			start++
		}

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
