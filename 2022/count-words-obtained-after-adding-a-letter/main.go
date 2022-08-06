package main

import (
	"fmt"
	"sort"
)

type trieNode struct {
	next [26]*trieNode
	end  bool
}

func (n *trieNode) add(w string) {
	if w == "" {
		n.end = true
		return
	}
	idx := w[0] - 'a'
	if n.next[idx] == nil {
		n.next[idx] = &trieNode{}
	}
	n.next[idx].add(w[1:])
}

func (n *trieNode) find(w string, skippedOnce bool) bool {
	if w == "" {
		return n.end && skippedOnce
	}
	idx := w[0] - 'a'
	if n.next[idx] == nil && skippedOnce {
		return false
	}
	if n.next[idx] == nil && !skippedOnce {
		return n.find(w[1:], true)
	}
	// BIGGEST TRICK HERE! If we reach here, the Trie contains
	// the current letter. But we also have to consider the case
	// where we still haven't skipped a letter! So if we haven't,
	// branch out to both cases: using or skipping the current letter.
	if skippedOnce {
		return n.next[idx].find(w[1:], true)
	}
	return n.next[idx].find(w[1:], false) || n.find(w[1:], true)
}

func sortWord(word string) string {
	bs := []byte(word)
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
	return string(bs)
}

// Time: O(linear to len(startWords)+len(targetWords))
// Space: O(linear to startWords) because they are stored in the Trie
//
// Tricky exercise. Quadratic solution times out, so put all startWords
// in a Trie and search all target words for a linear solution.
//
// The biggest tricks are:
// - Remember that you HAVE to skip a letter. Could happen at the last letter!
// - BIG! You might have to skip a letter even if it's IN the Trie, so
//   basically you have to try to skip once at every letter as long as you
//   haven't.
func wordCount(startWords []string, targetWords []string) int {
	trie := &trieNode{}
	for _, word := range startWords {
		trie.add(sortWord(word))
	}

	count := 0
	for _, word := range targetWords {
		if trie.find(sortWord(word), false) {
			count++
		}
	}
	return count
}

func main() {
	ts := []struct {
		startWords  []string
		targetWords []string
		expected    int
	}{
		{
			startWords:  []string{"ant", "act", "tack"},
			targetWords: []string{"tack", "act", "acti"},
			expected:    2,
		},
		{
			startWords:  []string{"ab", "a"},
			targetWords: []string{"abc", "abcd"},
			expected:    1,
		},
		{
			startWords:  []string{"g", "vf", "ylpuk", "nyf", "gdj", "j", "fyqzg", "sizec"},
			targetWords: []string{"r", "am", "jg", "umhjo", "fov", "lujy", "b", "uz", "y"},
			expected:    2,
		},
	}
	for _, tc := range ts {
		actual := wordCount(tc.startWords, tc.targetWords)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.startWords, tc.targetWords, tc.expected, actual)
		}
	}
}
