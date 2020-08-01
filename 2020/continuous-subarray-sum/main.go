package main

import "fmt"

// Time: O(n)
// Space: O(n)
//
// This is incredibly tricky to reason about!!!
//
// Compute the incremental sum, and the modulos of k of the sum:
//
// nums = [23,  2,  4,  6,  7]
// sums = [23, 25, 29, 35, 42]
// mods = [ 5,  1,  5,  5,  0]
//
// Special case: if you find a 0 in the mods, unless it's
// mods[0] then definitely you've found a valid subarray.
//
// The tricky part is finding a hidden 0:
//
// Look at sums[2]: it's 29, but by subtracting sums[0]
// it's 6 (2 + 4). In the same way that you can subtract
// any sum by a previous sum, you can do so with the mods.
//
// It follows that if you find two mods with the same value,
// like mods[0] and mods[2], there is a hidden mod zero in
// the subarray (mods[0+1]...mods[2]).
//
// So just find two mods with the same value. The catch is
// that they can't be contiguous, because the subarray
// would be of length 1.
//
// Special case: if k == 0 you can't calculate the mods.
// However, in this case you need to find a subarray whose
// sum is zero. If there's one, then the _sums_ array will
// have a duplicate (again, at a distance of 2 or more).
// An easy way of dealing with this special case is to
// not mod, but to just store the value of sums in mods.
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	sums := calculateSums(nums)

	modToFirstIdx := map[int]int{}
	for i, sum := range sums {
		mod := sum
		if k != 0 {
			mod = sum % k
		}

		if mod == 0 && i > 0 {
			return true
		}

		if idx, ok := modToFirstIdx[mod]; ok {
			if i-idx > 1 {
				return true
			}
		} else {
			modToFirstIdx[mod] = i
		}
	}
	return false
}

func calculateSums(nums []int) []int {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	return sums
}

func main() {
	ts := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{
			nums:     []int{23, 2, 4, 6, 7},
			k:        6,
			expected: true,
		},
		{
			nums:     []int{23, 2, 6, 4, 7},
			k:        6,
			expected: true,
		},
	}
	for _, tc := range ts {
		actual := checkSubarraySum(tc.nums, tc.k)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.nums, tc.k, tc.expected, actual)
		}
	}
}
