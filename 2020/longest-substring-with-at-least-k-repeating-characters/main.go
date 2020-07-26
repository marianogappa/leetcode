package main

import "fmt"

// Time: O(n) going 26 times over the input string
// Space: O(n) keeping a map of character frequencies
func longestSubstring(s string, k int) int {
	var mx int

	// It is very tricky to understand why it's necessary to do this 26 times.
	// This is a 2-pointer sliding window solution, but there's no simple way
	// to decide when to move the left pointer.
	//
	// _distinctLetterTarget_ is the number of distinct/unique letters that we are gonna
	// consider for the window. This doesn't make sense because the exercise
	// doesn't say anything about it, but it so happens that it works out as a
	// rule to move the left pointer that doesn't leave any interesting results
	// out, while keeping the complexity linear.
	for distinctLetterTarget := 1; distinctLetterTarget <= 26; distinctLetterTarget++ {
		mx = max(mx, longestSubstringWithTarget(s, k, distinctLetterTarget))
	}
	return mx
}

func longestSubstringWithTarget(s string, k, distinctLetterTarget int) int {
	var (
		freqs                                      = map[byte]int{}
		left, distinctLetterCount, numAtLeastK, mx int
	)
	// Move the right pointer until we have more distinct letters than the target.
	for right := 0; right < len(s); right++ {
		if freqs[s[right]] == 0 {
			distinctLetterCount++
		}
		freqs[s[right]]++
		if freqs[s[right]] == k {
			numAtLeastK++
		}

		// We have more distinct letters than target! Move left.
		for distinctLetterCount > distinctLetterTarget {
			if freqs[s[left]] == 1 {
				distinctLetterCount--
			}
			if freqs[s[left]] == k {
				numAtLeastK--
			}
			freqs[s[left]]--
			left++
		}

		// If we reached our target of distinct letters, and they are all repeated
		// at least k times, then we might have our maximum window.
		if distinctLetterCount == distinctLetterTarget && numAtLeastK == distinctLetterCount {
			mx = max(mx, right-left+1)
		}
	}
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ts := []struct {
		s        string
		k        int
		expected int
	}{
		{
			s:        "aaabb",
			k:        3,
			expected: 3,
		},
		{
			s:        "ababbc",
			k:        2,
			expected: 5,
		},
	}
	for _, tc := range ts {
		actual := longestSubstring(tc.s, tc.k)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.s, tc.k, tc.expected, actual)
		}
	}
}
