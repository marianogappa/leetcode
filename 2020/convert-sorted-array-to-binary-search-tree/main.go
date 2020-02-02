package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) because there's a call to sortedArrayToBST for every int
// Space: O(n) because every number is in the tree once
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// Calculate middle of array
	mid := len(nums)/2 - 1
	if len(nums)%2 == 1 {
		mid = len(nums) / 2
	}

	// Use middle value as node value, and recurse left and right with both halves of array
	return &TreeNode{nums[mid], sortedArrayToBST(nums[:mid]), sortedArrayToBST(nums[mid+1:])}
}

func (n *TreeNode) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("(%v,%v,%v)", n.Val, n.Left, n.Right)
}

func main() {
	ts := []struct {
		input    []int
		expected string
	}{
		{
			input:    []int{},
			expected: "",
		},
		{
			input:    []int{1},
			expected: "(1,,)",
		},
		{
			input:    []int{1, 2},
			expected: "(1,,(2,,))",
		},
		{
			input:    []int{-10, -3, 0, 5, 9},
			expected: "(0,(-10,,(-3,,)),(5,,(9,,)))",
		},
	}
	for _, tc := range ts {
		actual := sortedArrayToBST(tc.input)
		if tc.expected != actual.String() {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
