package main

import (
	"fmt"
	"reflect"
)

// Time: O(n)
// Space: O(n)
//
// This exercise is a brainteaser, which sucks because if you haven't seen it before you have to be smart or you won't
// get it :shrugs:
//
// I tried thinking of a tree, but it would be 2^n, and you'd think memoisation helps but it would take too much
// space, and thus time to construct. The only thing I can say is that there's obviously some way to leverage the fact
// that the numbers are even, so that should tell you there's a magical math way to solve this faster.
//
// Solution insight: these are even numbers, so adding them can only yield even numbers.
// Therefore, just add 2 + 4 + 6 ... greedily, and you either reach the target or you pass it, but the difference
// between your even target and the number you have is the number you need to cut from the sum. Note that this is
// guaranteed to be the optimal solution because you started adding from the smallest numbers possible.
//
// You have to return an array, so don't try to be too clever by using a set so you can remove a member in log time.
func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum == 0 || finalSum%2 != 0 {
		return []int64{}
	}

	nums := []int64{}
	for current := int64(2); finalSum > 0; current += 2 {
		nums = append(nums, current)
		finalSum -= current
	}

	if finalSum == 0 {
		return nums
	}

	// As a small time/space optimisation, don't create a new array, but instead shrink it by one, and replace the
	// offending number with the last array element (i.e. the one that would disappear after shrinking).
	for i, num := range nums {
		if num == -finalSum {
			nums[i] = nums[len(nums)-1]
			return nums[:len(nums)-1]
		}
	}

	// Should be unreachable
	return []int64{}
}

func main() {
	ts := []struct {
		input    int64
		expected []int64
	}{
		{
			input:    12,
			expected: []int64{2, 4, 6},
		},
		{
			input:    7,
			expected: []int64{},
		},
		{
			input:    28,
			expected: []int64{10, 4, 6, 8},
		},
	}
	for _, tc := range ts {
		actual := maximumEvenSplit(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
