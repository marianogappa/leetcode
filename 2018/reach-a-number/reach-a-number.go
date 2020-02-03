package main

import "fmt"

func reachNumber(target int) int {
	if target == 0 {
		return 0
	}
	if target < 0 {
		target = -target
	}
	k := 0
	for ; target > 0; k++ { // Sum (or minus) every number until we reach or exceed target
		target -= k
	}
	k--
	if target%2 == 0 { // If target is even, a +/- change on a number affects result in an
		return k // even number, so there will be a way to find target.
	}
	return k + 1 + k%2
}

func main() {
	ts := []struct {
		input    int
		expected int
	}{
		{
			input:    0,
			expected: 0,
		},
		{
			input:    1,
			expected: 1,
		},
		{
			input:    3,
			expected: 2,
		},
		{
			input:    2,
			expected: 3,
		},
		{
			input:    -1000000000,
			expected: 44723,
		},
	}
	for _, tc := range ts {
		actual := reachNumber(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
