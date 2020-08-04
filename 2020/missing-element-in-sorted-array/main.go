package main

import "fmt"

// Time: O(logn)
// Space: O(1)
func missingElement(nums []int, k int) int {
	// If the missing element is to the right of the last number in
	// _nums_, then a simple calculation should suffice.
	if missingNumbers(len(nums)-1, nums) < k {
		return nums[len(nums)-1] + k - missingNumbers(len(nums)-1, nums)
	}

	// Otherwise, binary search is the best option, considering
	// the array is sorted.
	//
	// Note that we're looking for the number *immediately* to the
	// right of where the missing element is.
	left := 0
	right := len(nums) - 1
	for left != right {
		mid := (left + right) / 2
		if missingNumbers(mid, nums) < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// We know the missing element is smaller than _nums[right]_
	// and larger than _nums[right-1]_.
	return nums[right] - (missingNumbers(right, nums) - k + 1)
}

// For [4, 7, 9, 10]
//  => [0  2  3   3]
// How many numbers are missing between the first one
// and the specified index.
func missingNumbers(idx int, nums []int) int {
	return nums[idx] - nums[0] - idx
}

func main() {
	ts := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{4, 7, 9, 10},
			k:        1,
			expected: 5,
		},
		{
			nums:     []int{4, 7, 9, 10},
			k:        3,
			expected: 8,
		},
		{
			nums:     []int{1, 2, 4},
			k:        3,
			expected: 6,
		},
	}
	for _, tc := range ts {
		actual := missingElement(tc.nums, tc.k)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.nums, tc.k, tc.expected, actual)
		}
	}
}
