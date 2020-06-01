package main

import "fmt"

// Time: O(n)
// Space: O(1)
//
// 1 + 2 + 3 + 4 = n*(n+1)/2 = (4*(4+1))/2 = 10
//
// Picture that the 0 in _nums_ is substituting
// the missing number.
//
// Thus the difference between the above formula
// and ∑nums is the missing number!
func missingNumber(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return len(nums)*(len(nums)+1)/2 - sum
}

func main() {
	ts := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{3, 0, 1},
			expected: 2,
		},
		{
			input:    []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			expected: 8,
		},
	}
	for _, tc := range ts {
		actual := missingNumber(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
