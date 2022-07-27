package main

import "fmt"

// Time: O(s*w) where s => len(s), w => len(words)
// Space: O(w*x) where w => len(words), x => len(largest word)
//
// It's important that finding which words advance for each character in "s" is a constant time operation. So we keep a
// "nextLetter" map up to date. Whenever we advance, either we match words or we don't. If we do, then new "next
// letters" exist.
func numMatchingSubseq(s string, words []string) int {
	// Initialise a "nextLetter" map that answers which words start with that letter
	nextLetter := map[byte][]string{}
	for _, word := range words {
		nextLetter[word[0]] = append(nextLetter[word[0]], word[1:])
	}

	numMatching := 0
	// For each character in s...
	for i := 0; i < len(s); i++ {
		c := s[i]
		// Find which words start (or continue) with that letter in constant time
		matchedWords := nextLetter[c]
		// Since we're gonna process them, remove them from the map
		delete(nextLetter, c)
		// For each word that starts (or continues) with the letter...
		for _, word := range matchedWords {
			// If the word is empty, that means we finished matching it! Increase the resulting total.
			if len(word) == 0 {
				numMatching++
				continue
			}
			// Otherwise, slide the next letter as key to the nextLetter map, and the rest of the word as value
			nextLetter[word[0]] = append(nextLetter[word[0]], word[1:])
		}
	}
	return numMatching
}

func main() {
	ts := []struct {
		s        string
		words    []string
		expected int
	}{
		{
			s:        "abcde",
			words:    []string{"a", "bb", "acd", "ace"},
			expected: 3,
		},
		{
			s:        "dsahjpjauf",
			words:    []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"},
			expected: 2,
		},
	}
	for _, tc := range ts {
		actual := numMatchingSubseq(tc.s, tc.words)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.s, tc.words, tc.expected, actual)
		}
	}
}
