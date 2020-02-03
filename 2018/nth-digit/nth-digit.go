package main

import (
	"fmt"
)

// Time: O(1)
// Space: O(1)
func findNthDigit(n int) int {
	if n <= 9 {
		return n
	}
	groups := []int{
		0,
		9, // 9 is the last number in group
		99,
		999,
		9999,
		99999,
		999999,
		9999999,
		99999999,
		999999999,
		2147483648,
	}

	// Find the right digit group and n -= accumulated digits
	i := 1 // indexing on groups; note that i is also the group's string length
	for n-(groups[i]-groups[i-1])*i > 0 {
		n -= (groups[i] - groups[i-1]) * i
		i++
	}

	// Find the right number within the digit group i
	// If the modulo == 0 then it's the previous number
	indexInGroup, modIndexInGroup := n/i, n%i
	if modIndexInGroup == 0 { // If there's no modulo, it's the last digit of previous
		indexInGroup--
		modIndexInGroup = i - 1
	} else {
		modIndexInGroup-- // Think of 10: reminder n = 1, 1%2=1 but should be 0
	}
	number := (groups[i-1] + 1) + indexInGroup

	// Return the nth digit in number
	return int([]byte(fmt.Sprintf("%v", number))[modIndexInGroup] - '0')
}

func main() {
	ts := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
		{7, 7},
		{8, 8},
		{9, 9},
		{10, 1},
		{11, 0},
		{12, 1},
		{13, 1},
		{14, 1},
		{15, 2},
		{16, 1},
		{17, 3},
		{18, 1},
		{19, 4},
		{20, 1},
		{21, 5},
		{22, 1},
		{23, 6},
		{24, 1},
		{25, 7},
		{26, 1},
		{27, 8},
		{28, 1},
		{29, 9},
		{30, 2},
		{31, 0},
		{32, 2},
		{33, 1},
		{34, 2},
		{35, 2},
		{36, 2},
	}
	for _, tc := range ts {
		actual := findNthDigit(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
