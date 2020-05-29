package main

import (
	"fmt"
	"reflect"
)

// Time: O(n)
// Space: O(n) if you consider out used space
func productExceptSelf(nums []int) []int {
	var (
		out        = make([]int, len(nums))
		leftCarry  = 1
		rightCarry = 1
	)
	for i := range out {
		out[i] = 1
	}
	// The carry auxiliaries will do this:
	//
	// nums       = [  1,  2, 3, 4]
	//
	// leftCarry  = [  1,  1, 2, 6]
	// rightCarry = [ 24, 12, 4, 1]
	// out  = l*r = [ 24, 12, 8, 6]
	//
	// Note that the last calculated carries are unused
	// (look at the loop).
	//
	// Also note the calculated carries are used in the
	// next iteration.
	for i := range nums {
		out[i] *= leftCarry
		out[len(nums)-1-i] *= rightCarry
		leftCarry *= nums[i]
		rightCarry *= nums[len(nums)-1-i]
	}
	return out
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
	}
	for _, tc := range ts {
		actual := productExceptSelf(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
