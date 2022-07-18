package main

import (
	"fmt"
	"reflect"
)

// Time: O(n)
// Space: O(n)
func getConcatenation(nums []int) []int {
	ans := make([]int, 2*len(nums))
	for i, num := range nums {
		ans[i] = num
		ans[i+len(nums)] = num
	}
	return ans
}

func main() {
	ts := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 1},
			expected: []int{1, 2, 1, 1, 2, 1},
		},
		{
			input:    []int{1, 3, 2, 1},
			expected: []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
	}
	for _, tc := range ts {
		actual := getConcatenation(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
