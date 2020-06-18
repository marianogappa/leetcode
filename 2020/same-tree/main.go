package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) where n == ∑len(node of each tree)
// Space: O(h) where h == height of smallest tree
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) || (p.Val != q.Val) {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	ts := []struct {
		p        *TreeNode
		q        *TreeNode
		expected bool
	}{
		{
			p:        &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
			q:        &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
			expected: true,
		},
		{
			p:        &TreeNode{1, &TreeNode{2, nil, nil}, nil},
			q:        &TreeNode{1, nil, &TreeNode{2, nil, nil}},
			expected: false,
		},
		{
			p:        &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{1, nil, nil}},
			q:        &TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}},
			expected: false,
		},
		{
			p:        &TreeNode{1, nil, nil},
			q:        &TreeNode{1, nil, nil},
			expected: true,
		},
		{
			p:        &TreeNode{2, nil, nil},
			q:        &TreeNode{1, nil, nil},
			expected: false,
		},
		{
			p:        &TreeNode{2, nil, nil},
			q:        nil,
			expected: false,
		},
		{
			p:        nil,
			q:        &TreeNode{2, nil, nil},
			expected: false,
		},
		{
			p:        nil,
			q:        nil,
			expected: true,
		},
	}
	for _, tc := range ts {
		actual := isSameTree(tc.p, tc.q)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.p, tc.q, tc.expected, actual)
		}
	}
}
