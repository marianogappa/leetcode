package main

import "fmt"

// Time: O(n)
// Space: O(1)
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// Split by zeroes.
	for i, num := range nums {
		if num == 0 {
			return max(max(0, maxProduct(nums[:i])), maxProduct(nums[i+1:]))
		}
	}
	// Count negatives, keep a total product, and index of first negative.
	var (
		countNegatives int
		firstNegativeI = -1
		prod           = 1
	)
	for i, num := range nums {
		prod *= num
		if num < 0 {
			countNegatives++
			if firstNegativeI == -1 {
				firstNegativeI = i
			}
		}
	}
	// If the number of negatives is even, _prod_ is positive.
	if countNegatives%2 == 0 {
		return prod
	}
	// If there's only one negative, split by it
	if countNegatives == 1 {
		return max(maxProduct(nums[:firstNegativeI]), maxProduct(nums[firstNegativeI+1:]))
	}
	// Otherwise, there are at least 3 negatives. We have to dismiss up to the leftmost
	// or rightmost one (non-inclusive), so let's keep the largest.
	// Remember the product will end up negative, so we want to dismiss the largest side.
	var (
		leftSide  = 1
		leftI     = 0
		rightSide = 1
		rightI    = len(nums) - 1
	)
	for i := 0; i < len(nums); i++ {
		leftSide *= nums[i]
		leftI = i
		if leftSide < 0 {
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		rightSide *= nums[i]
		rightI = i
		if rightSide < 0 {
			break
		}
	}
	if leftSide > rightSide {
		return maxProduct(nums[leftI+1:])
	}
	return maxProduct(nums[:rightI])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ts := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{2, 3, -2, 4},
			expected: 6,
		},
		{
			input:    []int{-2, 0, -1},
			expected: 0,
		},
	}
	for _, tc := range ts {
		actual := maxProduct(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
