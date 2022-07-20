package main

import "fmt"

// Time: O(n) where n == len(nums), because we iterate nums once.
// Space: O(1)
//
// The strategy is to run a sliding window.
//
// Important intuition: once we find a valid window, the answer cannot be a smaller one, so we should never shrink the
//                      window. That's important, because sliding (i.e. keeping size but moving) means moving start as
//                      well as end. That simplifies the algorithm because end ALWAYS moves.
//
// End always moves, and start moves when k goes below 0 (i.e. becomes an invalid window). This means that k can go well
// in the negatives, but that's ok because we did find a valid window of that size. We will only enlarge the window
// when it's valid.
func longestOnes(nums []int, k int) int {
	var start, end int

	for end = 0; end < len(nums); end++ {
		// Decrease k when end steps into a zero
		if nums[end] == 0 {
			k--
		}
		// Only advance start when window is invalid
		if k < 0 {
			// Increase k when start leaves a zero
			if nums[start] == 0 {
				k++
			}
			start++
		}
	}
	end-- // substracting 1 because in Go, "end" leaves the for-loop as len(nums)
	return end - start + 1
}

func main() {
	ts := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:        2,
			expected: 6,
		},
		{
			nums:     []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:        3,
			expected: 10,
		},
	}
	for _, tc := range ts {
		actual := longestOnes(tc.nums, tc.k)
		if tc.expected != actual {
			fmt.Printf("For (%v, %v) expected %v but got %v\n", tc.nums, tc.k, tc.expected, actual)
		}
	}
}
