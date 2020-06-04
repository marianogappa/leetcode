package main

import (
	"fmt"
)

// A Trie is necessary because it's important that going
// through the string whilst matching to the dictionary
// is a linear algorithm. Otherwise we'd have to keep
// indices and loop through words to match.
type trie struct {
	nodes map[byte]*trie
	stop  bool
}

func (t *trie) put(s string) {
	if len(s) == 0 {
		return
	}
	if t.nodes == nil {
		t.nodes = make(map[byte]*trie)
	}
	if _, ok := t.nodes[s[0]]; !ok {
		t.nodes[s[0]] = &trie{make(map[byte]*trie), len(s) == 1}
	}
	t.nodes[s[0]].stop = t.nodes[s[0]].stop || len(s) == 1
	t.nodes[s[0]].put(s[1:])
}

func (t *trie) moveTo(b byte) *trie {
	if t.nodes == nil || t.nodes[b] == nil {
		return nil
	}
	return t.nodes[b]
}

// Time: O(n) because of memo
// Space: O(n+m) length of string (memo and stack) + length of wordDict's letters (stored in trie)
func wordBreak(s string, wordDict []string) bool {
	// First construct a Trie with the dictionary to achieve efficient matching.
	t := &trie{}
	for _, word := range wordDict {
		t.put(word)
	}

	// While solving the matching, keep a memo that answers if a solution can
	// be reached from a given position on the string 's'.
	return doWordBreak(s, 0, t, map[int]bool{})
}

func doWordBreak(s string, i int, t *trie, memo map[int]bool) bool {
	if i >= len(s) {
		return true
	}
	if v, ok := memo[i]; ok {
		return v
	}

	// Starting at the roots of the Trie, go through the Trie as
	// we go through the string, trying to match to dictionary
	// words.
	cursor := t
	for j := i; j < len(s); j++ {
		cursor = cursor.moveTo(s[j])
		// If there was no match on the Trie, then this string
		// cannot be constructed from the dictionary words.
		if cursor == nil {
			memo[i] = false
			return false
		}
		// If there was a match, and the Trie says a word ends
		// on this character, then check recursively if there
		// is a solution from the remaining substring.
		//
		// IMPORTANT! If there's no solution from here, there
		// still might be a solution for a larger dictionary
		// word than the one we just matched, so keep going!
		if cursor.stop && doWordBreak(s, j+1, t, memo) {
			memo[i] = true
			return true
		}
	}
	memo[i] = false
	return false
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
		actual := wordBreak(tc.s, tc.wordDict)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.s, tc.wordDict, tc.expected, actual)
		}
	}
}
