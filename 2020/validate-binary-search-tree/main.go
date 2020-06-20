package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("(%v,%v,%v)", n.Val, n.Left, n.Right)
}

// Time: O(n)
// Space: O(h)
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// The tricky part is that e.g. the right subtree of a tree's
	// node must have all Vals smaller than that node's Val, so
	// we need to keep a memory of "global" max and min, in
	// addition to the local thresholds.
	return dfs(root, math.MinInt64, math.MaxInt64)
}

func dfs(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && (root.Left.Val >= root.Val || root.Left.Val <= min) {
		return false
	}
	if root.Right != nil && (root.Right.Val <= root.Val || root.Right.Val >= max) {
		return false
	}
	return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
}

func main() {
	ts := []struct {
		input    *TreeNode
		expected bool
	}{
		{
			input:    &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
			expected: true,
		},
		{
			input:    &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}}},
			expected: false,
		},
		{
			input:    &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{10, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}},
			expected: false,
		},
		{
			input:    &TreeNode{-2147483648, nil, &TreeNode{2147483647, nil, nil}},
			expected: true,
		},
	}
	for _, tc := range ts {
		actual := isValidBST(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
