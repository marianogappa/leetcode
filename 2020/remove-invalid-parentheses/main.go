package main

import (
	"container/list"
	"fmt"
	"reflect"
)

// This solution is pretty easy to understand
// but it's very inefficient.
//
// Until a valid solution is found, BFS through
// all options of removing characters. Then,
// choose all valid solutions within that level.
//
// Keep a set of solutions to not chase the same
// one and produce duplicates.
func removeInvalidParentheses(s string) []string {
	var (
		isLastLevel bool
		results     = []string{}
		queue       = list.New()
		visited     = map[string]struct{}{}
	)
	queue.PushFront(s)
	for queue.Len() > 0 {
		str := queue.Front().Value.(string)
		queue.Remove(queue.Front())
		if _, ok := visited[str]; ok {
			continue
		}
		visited[str] = struct{}{}

		if isValid(str) {
			isLastLevel = true
			results = append(results, str)
		}

		if isLastLevel {
			continue
		}

		pushAllOptions(str, queue, visited)
	}

	return results
}

func isParens(b byte) bool {
	return b == '(' || b == ')'
}

func isValid(s string) bool {
	count := 0
	for _, b := range s {
		switch b {
		case '(':
			count++
		case ')':
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}

func pushAllOptions(s string, l *list.List, visited map[string]struct{}) {
	for i := range s {
		if !isParens(s[i]) {
			continue
		}
		l.PushBack(fmt.Sprintf("%v%v", s[:i], s[i+1:]))
	}
}

func main() {
	ts := []struct {
		input    string
		expected []string
	}{
		{
			input:    "()())()",
			expected: []string{"()()()", "(())()"},
		},
		{
			input:    "(a)())()",
			expected: []string{"(a)()()", "(a())()"},
		},
		{
			input:    ")(",
			expected: []string{""},
		},
	}
	for _, tc := range ts {
		actual := removeInvalidParentheses(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
