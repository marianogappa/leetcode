package main

import (
	"fmt"
	"reflect"
)

// This solution requires some thought because, although it works,
// the space complexity
func subsets(nums []int) [][]int {
	result := [][]int{{}}
	for i := 1; i <= len(nums); i++ {
		result = append(result, subsetsOfLen(i, nums)...)
	}
	return result
}

func subsetsOfLen(length int, nums []int) [][]int {
	if length == 0 {
		return [][]int{{}}
	}
	results := [][]int{}
	for i := 0; i < len(nums)-length+1; i++ {
		for _, subset := range subsetsOfLen(length-1, nums[i+1:]) {
			results = append(results, append([]int{nums[i]}, subset...))
		}
	}
	return results
}

func main() {
	ts := []struct {
		input    []int
		expected [][]int
	}{
		{
			input:    []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}
	for _, tc := range ts {
		actual := subsets(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
