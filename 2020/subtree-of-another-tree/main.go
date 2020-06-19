package main

import "fmt"

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

// Time: O(n*m) nodes of the two trees
// Space: O(h) height of longest tree
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return dfs(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func dfs(s *TreeNode, t *TreeNode) bool {
	if t == nil && s == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val && dfs(s.Left, t.Left) && dfs(s.Right, t.Right)
}

func main() {
	ts := []struct {
		s        *TreeNode
		t        *TreeNode
		expected bool
	}{
		{
			s:        &TreeNode{3, &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}, &TreeNode{5, nil, nil}},
			t:        &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}},
			expected: true,
		},
		{
			s:        &TreeNode{3, &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, &TreeNode{0, nil, nil}, nil}}, &TreeNode{5, nil, nil}},
			t:        &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}},
			expected: false,
		},
	}
	for _, tc := range ts {
		actual := isSubtree(tc.s, tc.t)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.s, tc.t, tc.expected, actual)
		}
	}
}
