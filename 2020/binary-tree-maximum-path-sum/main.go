package main

import (
	"fmt"
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
		return 0, 0
	}
	nonCumMaxLeft, cumMaxLeft := dfs(root.Left)
	nonCumMaxRight, cumMaxRight := dfs(root.Right)

	// Non-cumulative results are paths that don't add up to
	// upper layers of the tree.
	//
	// Possible options:
	// 1. Value of current node.
	// 2. Value of current node plus sum of cumulatives of children.
	// 3. Cumulatives and non-cumulatives on either side.
	nonCums := []int{root.Val, root.Val + cumMaxLeft + cumMaxRight}
	if root.Left != nil {
		nonCums = append(nonCums, nonCumMaxLeft, cumMaxLeft)
	}
	if root.Right != nil {
		nonCums = append(nonCums, nonCumMaxRight, cumMaxRight)
	}

	// Cumulative results are paths that add up to upper layers of the tree.
	//
	// Possible options:
	// 1. Value of current node.
	// 2. Value of current node plus max of cumulatives of children.
	return maxSlice(nonCums), max(root.Val, root.Val+max(cumMaxLeft, cumMaxRight))
}

func maxSlice(ns []int) int {
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
