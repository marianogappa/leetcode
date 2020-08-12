package main

import "fmt"

// Time: O(n*w)
// Space: O(n*w) n == len(words) w == len(longest word)
type status int

const (
	todo status = iota
	doing
	done
)

func alienOrder(words []string) string {
	if len(words) == 0 {
		return ""
	}
	// If there's only one word, any order is valid
	// but it must contain all words.
	if len(words) == 1 {
		// Time: O(n*w)
		bs := map[byte]struct{}{}
		for i := 0; i < len(words[0]); i++ {
			bs[words[0][i]] = struct{}{}
		}
		bss := []byte{}
		for b := range bs {
			bss = append(bss, b)
		}
		return string(bss)
	}

	// Create adjacency list.
	// Time: O(n*w)
	adjList := make(map[byte]map[byte]struct{}, 26)
	for i := 1; i < len(words); i++ {
		// If two words have identical prefixes and the second
		// one is larger, then they are not lexicographically
		// sorted! Return ""!
		if !populateAdjList(words[i-1], words[i], adjList) {
			return ""
		}
	}

	// Need to know all letters involved, because they
	// might not be part of the known ordered but still
	// have to be in the string.
	// Time: O(n*w)
	letters := map[byte]struct{}{}
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			letters[word[i]] = struct{}{}
		}
	}

	// Do topological sort over all letters.
	// Disconnected letters go at the end according
	// to the first example. Largest letter should
	// appear later in the order.
	visited := map[byte]status{}
	order := make([]byte, len(letters))
	orderI := 0
	// There are at most n-1 edges in total, but n*w maximum letters, so:
	// Time: O(n*w)
	for letter := range letters {
		if visited[letter] == done {
			continue
		}
		isCycle := dfs(letter, adjList, visited, order, &orderI)
		if isCycle {
			return ""
		}
	}

	return string(order)
}

func dfs(letter byte, adjList map[byte]map[byte]struct{}, visited map[byte]status, order []byte, orderI *int) bool {
	if visited[letter] == done {
		return false
	}
	if visited[letter] == doing {
		return true
	}
	visited[letter] = doing
	for smallerLetter := range adjList[letter] {
		if dfs(smallerLetter, adjList, visited, order, orderI) {
			return true
		}
	}
	visited[letter] = done
	order[*orderI] = letter
	*orderI++
	return false
}

// Time: O(min(w1,w2)) ~= O(w)
func populateAdjList(w1, w2 string, adjList map[byte]map[byte]struct{}) bool {
	if len(w1) == 0 || len(w2) == 0 {
		return true
	}
	for i := 0; i < len(w1) && i < len(w2); i++ {
		if w1[i] != w2[i] {
			if adjList[w2[i]] == nil {
				adjList[w2[i]] = map[byte]struct{}{}
			}
			adjList[w2[i]][w1[i]] = struct{}{}
			return true
		}
	}
	// If two words have identical prefixes and the second
	// one is larger, then they are not lexicographically
	// sorted!
	return len(w1) <= len(w2)
}

func main() {
	ts := []struct {
		input    []string
		expected string
	}{
		{
			input: []string{"wrt",
				"wrf",
				"er",
				"ett",
				"rftt"},
			expected: "wertf",
		},
		{
			input: []string{"z",
				"x"},
			expected: "zx",
		},
		{
			input: []string{"z",
				"x",
				"z"},
			expected: "",
		},
		{
			input:    []string{"abc", "ab"},
			expected: "",
		},
	}
	for _, tc := range ts {
		actual := alienOrder(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
