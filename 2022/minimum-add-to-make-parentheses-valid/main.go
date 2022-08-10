package main

import "fmt"

// Time: O(n)
// Space: O(1)
//
// Just calculate a running sum!
//
// The only trick is that we must keep a count of negatives separately
// and add that to the running sum, because otherwise ")(" would be 0, and must be 2.
func minAddToMakeValid(s string) int {
	var (
		runningSum int
		fix        int
	)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			runningSum++
		case ')':
			runningSum--
		}
		if runningSum == -1 {
			fix++
			runningSum++
		}
	}
	return fix + runningSum
}

func main() {
	ts := []struct {
		input    string
		expected int
	}{
		{
			input:    "())",
			expected: 1,
		},
		{
			input:    "(((",
			expected: 3,
		},
		{
			input:    ")(",
			expected: 2,
		},
		{
			input:    ")))(",
			expected: 4,
		},
		{
			input:    "",
			expected: 0,
		},
		{
			input:    "((((",
			expected: 4,
		},
		{
			input:    "))))",
			expected: 4,
		},
	}
	for _, tc := range ts {
		actual := minAddToMakeValid(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
