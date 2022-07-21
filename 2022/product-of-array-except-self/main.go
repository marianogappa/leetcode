package main

import (
	"fmt"
	"reflect"
)

// Time: O(n)
// Space: O(1) Note: exercise says "The output array does not count as extra space for space complexity analysis."
//
//
//
// Strategy is to do a "running product" from the (start+1->end),
// and then from the (end-1->start), and multiply those together:
//
// for [1, 2, 3, 4]:
//
// start with the zero value of multiplication: [1, 1, 1, 1]
// run a running product from (start+1 -> end): [1, 1*1, 1*1*2, 1*1*2*3] = [1, 1, 2, 6]
// run a running product from (end-1 -> start): [1*4*3*2, 1*4*3, 1*4, 1] = [24, 12, 4, 1]
// multiply those together: [1*24, 1*12, 2*4, 6*1] = [24, 12, 8, 6]
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = 1

	// Running product from (start+1 -> end)
	runningProduct := 1
	for i := 1; i < len(result); i++ {
		runningProduct *= nums[i-1]
		result[i] = runningProduct
	}

	// Running product from (start+1 -> end)
	runningProduct = 1
	for i := len(nums) - 2; i >= 0; i-- {
		runningProduct *= nums[i+1]
		result[i] *= runningProduct
	}

	return result
}

func main() {
	ts := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			input:    []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
	}
	for _, tc := range ts {
		actual := 1
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
