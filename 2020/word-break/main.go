package main

import (
	"fmt"
)

// Time: O(s^3) two for-loops plus a string comparison of length s in the worst case
// Space: O(w*w'+s) hashmap of words and bool array of size len(s)
func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		m[word] = true
	}

	// dp[i] means "is there a valid segmentation of _s_
	// up to the ith index?"
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// If there's a valid segmentation up to _j_
			// and s[j:i] is a valid word, then there's a
			// valid segmentation up to _i_.
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// Time: O(s*w*max(w',s))
// Space: O(s)
func anotherWordBreak(s string, wordDict []string) bool {
	// dp[i] means "is there a valid segmentation of _s_
	// up to the ith index?"
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			// If there's a word that ends at index _i_ of _s_, and starts
			// at an already valid index, then _dp[i]_ is also valid.
			if len(word) <= i && dp[i-len(word)] && s[i-len(word):i] == word {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func main() {
	ts := []struct {
		s        string
		wordDict []string
		expected bool
	}{
		{
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true,
		},
		{
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true,
		},
		{
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
		{
			s:        "aaaaaaa",
			wordDict: []string{"aaaa", "aa"},
			expected: false,
		},
	}
	for _, tc := range ts {
		actual := anotherWordBreak(tc.s, tc.wordDict)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.s, tc.wordDict, tc.expected, actual)
		}
	}
}
