package main

import "fmt"

// Time: O(n)
// Space: O(n)
//
// sumCounts will continue keys for all values
// of the running sum, and with a value of the
// number of times this running sum was reached.
//
// Each time we add a running sum value, we
// check how many times that value minus _k_
// was reached. If n>0, there were n subarrays
// that go up to the current number summing up
// to _k_.
func subarraySum(nums []int, k int) int {
	var total, runningSum int
	var sumCounts = map[int]int{0: 1}
	for _, num := range nums {
		runningSum += num
		if count, ok := sumCounts[runningSum-k]; ok {
			total += count
		}
		sumCounts[runningSum]++
	}
	return total
}

func main() {
	ts := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
	}
	for _, tc := range ts {
		actual := subarraySum(tc.nums, tc.k)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.nums, tc.k, tc.expected, actual)
		}
	}
}
