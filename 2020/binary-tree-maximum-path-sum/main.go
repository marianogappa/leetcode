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
func maxPathSum(root *TreeNode) int {
	return max(dfs(root))
}

func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, math.MinInt32
	}

	maxVerticalPathLeft, maxPathLeft := dfs(root.Left)
	maxVerticalPathRight, maxPathRight := dfs(root.Right)

	// The max vertical path is the max path from a leaf node to current node, or zero
	// if all paths are negative. All possible vertical paths must include the current
	// node, so it's always added at the end.
	maxVerticalPath := maxSlice(0, maxVerticalPathLeft, maxVerticalPathRight) + root.Val

	// It's also possible that the max is not a vertical path but the combination of two
	// paths joined by the current node. It's also possible that the max path has already
	// been found on a deeper iteration to the left or to the right, or that the max path
	// is the maxVertical path calculated above this line.
	maxPath := maxSlice(maxPathLeft, maxPathRight, maxVerticalPath, maxVerticalPathLeft+maxVerticalPathRight+root.Val)

	return maxVerticalPath, maxPath
}

func maxSlice(ns ...int) int {
	m := ns[0]
	for _, n := range ns {
		m = max(m, n)
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ts := []struct {
		input    *TreeNode
		expected int
	}{
		{
			input:    &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
			expected: 6,
		},
		{
			input:    &TreeNode{2, &TreeNode{-1, nil, nil}, &TreeNode{-2, nil, nil}},
			expected: 2,
		},
		{
			input:    &TreeNode{-10, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}},
			expected: 42,
		},
		{
			input:    &TreeNode{-3, nil, nil},
			expected: -3,
		},
		{
			input:    &TreeNode{-1, nil, &TreeNode{9, &TreeNode{-6, nil, nil}, &TreeNode{3, nil, &TreeNode{-2, nil, nil}}}},
			expected: 12,
		},
		{
			input:    &TreeNode{1, &TreeNode{-2, &TreeNode{1, &TreeNode{-1, nil, nil}, nil}, &TreeNode{3, nil, nil}}, &TreeNode{-3, &TreeNode{-2, nil, nil}, nil}},
			expected: 3,
		},
	}
	for _, tc := range ts {
		actual := maxPathSum(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
