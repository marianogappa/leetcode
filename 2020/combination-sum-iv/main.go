package main

import "fmt"

// Time: O(n*m) where n is target and m is length of nums
// Space: O(n) where n is target
func combinationSum4(nums []int, target int) int {
	return doCombinationSum4(nums, target, map[int]int{})
}

func doCombinationSum4(nums []int, target int, memo map[int]int) int {
	if target == 0 {
		return 1
	}
	if v, ok := memo[target]; ok {
		return v
	}
	count := 0
	for _, num := range nums {
		if target-num < 0 {
			continue
		}
		count += doCombinationSum4(nums, target-num, memo)
	}
	memo[target] = count
	return memo[target]
}

func main() {
	ts := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{1, 2, 3},
			target:   4,
			expected: 7,
		},
	}
	for _, tc := range ts {
		actual := combinationSum4(tc.nums, tc.target)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.nums, tc.target, tc.expected, actual)
		}
	}
}
