package main

import "fmt"

// Time: O(logn)
// Space: O(1)
func search(nums []int, target int) int {
	switch len(nums) {
	case 0:
		return -1
	case 1:
		if nums[0] != target {
			return -1
		}
		return 0
	case 2:
		if nums[0] == target {
			return 0
		}
		if nums[1] == target {
			return 1
		}
		return -1
	default:
		// _target_ is either on the left half or on the right half
		// of _nums_.
		//
		// Because it's only misordered in (pivot-1, pivot),
		// we know that one of the two halves is perfectly ordered,
		// and thus we can ask if _target_ is inside it, to perform
		// the binary search.
		//
		// The _if_ is complicated but it's figuring out which half
		// is ordered, and then if target is on the left half.
		mid := len(nums) / 2
		if (nums[0] < nums[mid] && nums[0] <= target && nums[mid] >= target) ||
			(nums[mid+1] <= nums[len(nums)-1] && (nums[mid+1] > target || nums[len(nums)-1] < target)) {
			return search(nums[0:mid+1], target)
		}
		// Note that when _target_ is on the right half, the answer
		// to the recursion is about a subarray, so we need to add
		// _mid+1_ to it.
		res := search(nums[mid+1:], target)
		if res == -1 {
			return -1
		}
		return res + mid + 1
	}
}

func main() {
	ts := []struct {
		input    []int
		target   int
		expected int
	}{
		{
			input:    []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			input:    []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			input:    []int{5, 1, 3},
			target:   1,
			expected: 1,
		},
	}
	for _, tc := range ts {
		actual := search(tc.input, tc.target)
		if tc.expected != actual {
			fmt.Printf("For (%v, %v) expected %v but got %v\n", tc.input, tc.target, tc.expected, actual)
		}
	}
}
