package main

import (
	"fmt"
)

type node struct {
	val  byte
	prev *node
}

func (n *node) push(val byte) *node {
	return &node{val, n}
}

func (n *node) pop() *node {
	return n.prev
}

// Time: O(n)
// Space: O(n)
func isValid(s string) bool {
	if s == "" {
		return true
	}

	closerOf := map[byte]byte{'(': ')', '{': '}', '[': ']'}
	q := &node{s[0], nil}
	for i := 1; i < len(s); i++ {
		switch s[i] {
		case '(', '{', '[':
			q = q.push(s[i])
		case ')', '}', ']':
			if q == nil || s[i] != closerOf[q.val] {
				return false
			}
			q = q.pop()
		}
	}
	return q == nil
}

func main() {
	ts := []struct {
		input    string
		expected bool
	}{
		{
			input:    "()",
			expected: true,
		},
		{
			input:    "()[]{}",
			expected: true,
		},
		{
			input:    "(]",
			expected: false,
		},
		{
			input:    "([)]",
			expected: false,
		},
	}
	for _, tc := range ts {
		actual := isValid(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
