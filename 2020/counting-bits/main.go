package main

import (
	"fmt"
	"reflect"
)

// Time: O(n)
// Space: O(n) or O(1) if solution space doesn't count
//
// Looking at how the numbers are constructed:
//    0
//    1
//   10
//   11
//  100
//  101
//  110
//  111
// 1000
// 1100
// 1101
// 1110
// 1111
//
// Every time there's a new 1:
// 1) the numbers after that 1 are exactly the
//    same number that came before up to that.
// 2) it's a power of two
//
// Using this clever way to ask for pows of 2:
// i&(i-1) == 0
//
// We can construct the array iteratively.
func countBits(num int) []int {
	counts := make([]int, num+1)
	previousExp := 0

	for i := range counts {
		switch {
		case i == 0:
			counts[i] = 0
		case i == 1:
			counts[i] = 1
			// If it's a power of two
		case i&(i-1) == 0:
			previousExp = i
			counts[i] = 1
		// If it's not a power of two, the story
		// repeats from the previous power of two,
		// with a leading 1.
		default:
			counts[i] = 1 + counts[i-previousExp]
		}
	}
	return counts
}

func main() {
	ts := []struct {
		input    int
		expected []int
	}{
		{
			input:    2,
			expected: []int{0, 1, 1},
		},
		{
			input:    5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tc := range ts {
		actual := countBits(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
