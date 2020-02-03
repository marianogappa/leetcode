package main

import "fmt"

// Time: O(n)
// Space: O(n) (if considering copy)
func checkPossibility(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	wrong := 0
	for i := 1; i < len(nums) && wrong <= 1; i++ {
		if nums[i-1] > nums[i] {
			wrong++
			if i >= 2 && nums[i-2] > nums[i] {
				nums[i] = nums[i-1]
			} else {
				nums[i-1] = nums[i]
			}
		}
	}
	return wrong <= 1
}

func main() {
	var ts = []struct {
		input    []int
		expected bool
	}{
		{
			[]int{4, 2, 3},
			true,
		},
		{
			[]int{4, 2, 1},
			false,
		},
		{
			[]int{},
			true,
		},
		{
			[]int{3, 4, 2, 3},
			false,
		},
		{
			[]int{1},
			true,
		},
		{
			[]int{1, 1},
			true,
		},
		{
			[]int{1, 1, 1},
			true,
		},
		{
			[]int{1, 2},
			true,
		},
		{
			[]int{2, 1},
			true,
		},
		{
			[]int{3, 2, 1},
			false,
		},
		{
			[]int{3, 3, 1},
			true,
		},
		{
			[]int{2, 1, 3, 4, 5},
			true,
		},
		{
			[]int{2, 1, 3, 4, 3},
			false,
		},
	}
	for _, tc := range ts {
		actual := checkPossibility(tc.input)
		if tc.expected != actual {
			fmt.Printf("Failed assertion for input %v\n", tc.input)
		}
	}
}
