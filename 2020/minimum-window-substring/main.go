package main

import (
	"fmt"
	"math"
)

// Time: O(n)
// Space: O(n)
func minWindow(s string, t string) string {
	// Calculate the requirements for a valid window:
	// which letters and how many of them.
	var (
		letterFreqs = map[byte]int{}
		counter     = len(t)
	)
	for i := 0; i < len(t); i++ {
		letterFreqs[t[i]]++
	}

	// Sliding window algorithm: outer loop increments
	// start, inner loop increments end.
	var (
		minStart   int
		minEnd     = math.MaxInt32
		start, end int
	)
	// Start sliding window by incrementing _end_ until window is valid.
	for end < len(s) {
		// Insight: only relevant letters can be positive.
		if letterFreqs[s[end]] > 0 {
			counter--
		}
		// Insight: irrelevant letters become negative.
		letterFreqs[s[end]]--
		end++

		// While window remains valid, start++.
		for counter == 0 {
			// On every valid window iteration we may found the
			// smallest window, so check every time.
			if end-start < minEnd-minStart {
				minStart = start
				minEnd = end
			}
			// Removing letter at _start_ means incrementing it in the map.
			letterFreqs[s[start]]++

			// Because irrelevant letters are all negative when reaching
			// them from _end_, they can only get up to zero on this second
			// pass. Therefore, this is asking: is the letter at _start_
			// within _t_?
			if letterFreqs[s[start]] > 0 {
				// Exit loop after incrementing _start_, given the window
				// is no longer valid.
				counter++
			}
			start++
		}
	}

	// If window was never found, _minEnd_ never leaves MaxInt32.
	if minEnd == math.MaxInt32 {
		return ""
	}

	return s[minStart:minEnd]
}

func main() {
	ts := []struct {
		inputS   string
		inputT   string
		expected string
	}{
		{
			inputS:   "ADOBECODEBANC",
			inputT:   "ABC",
			expected: "BANC",
		},
	}
	for _, tc := range ts {
		actual := 1
		if tc.expected != minWindow(tc.inputS, tc.inputT) {
			fmt.Printf("For (%v, %v) expected %v but got %v\n", tc.inputS, tc.inputT, tc.expected, actual)
		}
	}
}
