package main

import "fmt"

// Time:    O(n)
// Space:   O(1)
// Summary: (a) Advance pointers i from the left and j from the right until they're out of order to each other or
//              each one's next.
//          (b) Find minimum and maximum of the subarray (i, j).
//          (c) Go back from i and forth from j while the current number is >min and <max. Return that subarray.
func findUnsortedSubarray(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var (
		i     = 0
		j     = len(nums) - 1
		stopI = false
		stopJ = false
	)
	for i < j {
		if nums[i] > nums[j] || (stopI && stopJ) {
			var min, max = 1<<31 - 1, -1 << 31
			for k := i; k <= j; k++ {
				if nums[k] < min {
					min = nums[k]
				}
				if nums[k] > max {
					max = nums[k]
				}
			}
			for ; i >= 0 && nums[i] > min; i-- {
			}
			for ; j < len(nums) && nums[j] < max; j++ {
			}
			return j - i + 1 - 2
		}
		if nums[i] > nums[i+1] {
			stopI = true
		}
		if nums[j] < nums[j-1] {
			stopJ = true
		}
		if !stopI {
			i++
		}
		if !stopJ {
			j--
		}
	}
	return 0
}

func main() {
	ts := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{2, 6, 4, 8, 10, 9, 15},
			expected: 5,
		},
		{
			input:    []int{},
			expected: 0,
		},
		{
			input:    []int{1},
			expected: 0,
		},
		{
			input:    []int{1, 2},
			expected: 0,
		},
		{
			input:    []int{3, 2},
			expected: 2,
		},
		{
			input:    []int{5, 4, 3, 2, 1},
			expected: 5,
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			input:    []int{2, 6, 4, 8, 10, 16, 15},
			expected: 6,
		},
		{
			input:    []int{1, 3, 2, 2, 2},
			expected: 4,
		},
		{
			input:    []int{1, 2, 5, 3, 4},
			expected: 3,
		},
		{
			input:    []int{1, 3, 2, 3, 3},
			expected: 2,
		},
	}
	for _, tc := range ts {
		actual := findUnsortedSubarray(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
