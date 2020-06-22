package main

import (
	"fmt"

	"reflect"
)

type node struct {
	next [26]*node
	stop string
}

func (n *node) insert(s string, origS string) {
	if s == "" {
		n.stop = origS
		return
	}
	c := s[0] - 'a'
	if n.next[c] == nil {
		n.next[c] = &node{next: [26]*node{}}
	}
	n.next[c].insert(s[1:], origS)
}

func (n *node) search(s string) string {
	if s == "" {
		return n.stop
	}
	c := s[0] - 'a'
	if n.next[c] == nil {
		return ""
	}
	return n.next[c].search(s[1:])
}

func findWords(board [][]byte, words []string) []string {
	// Make a trie with all words.
	trie := &node{next: [26]*node{}, stop: ""}
	for _, word := range words {
		trie.insert(word, word)
	}

	// From every starting point in the matrix,
	// follow the trie in all directions to
	// match words.
	found := []string{}
	for y := range board {
		for x := range board[y] {
			dfs(board, x, y, trie, &found)
		}
	}
	return found
}

func dfs(board [][]byte, x, y int, trie *node, found *[]string) {
	if x < 0 || y < 0 || y >= len(board) || x >= len(board[0]) || board[y][x] == ' ' {
		return
	}
	c := board[y][x]
	cMa := c - 'a'
	// Exit the recursion if there's no match in the Trie.
	if trie.next[cMa] == nil {
		return
	}
	if trie.next[cMa].stop != "" {
		*found = append(*found, trie.next[cMa].stop)
		// Don't forget to eliminate the word from the Trie.
		// Otherwise it can be picked up again.
		trie.next[cMa].stop = ""
		// Don't return here! There might be longer words
		// that start with this same characters.
	}
	board[y][x] = ' '
	dfs(board, x-1, y, trie.next[cMa], found)
	dfs(board, x+1, y, trie.next[cMa], found)
	dfs(board, x, y-1, trie.next[cMa], found)
	dfs(board, x, y+1, trie.next[cMa], found)
	board[y][x] = c
}

func main() {
	ts := []struct {
		board    [][]byte
		words    []string
		expected []string
	}{
		{
			board: [][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			},
			words:    []string{"oath", "pea", "eat", "rain"},
			expected: []string{"oath", "eat"},
		},
		{
			board: [][]byte{
				{'a', 'a'},
			},
			words:    []string{"a"},
			expected: []string{"a"},
		},
	}
	for _, tc := range ts {
		actual := findWords(tc.board, tc.words)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.words, tc.expected, actual)
		}
	}
}
