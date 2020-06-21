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

// Time: O(k)
// Space: O(1)
// In-order traversal up to kth
func kthSmallest(root *TreeNode, k int) int {
	var val int
	dfs(root, &k, &val)
	return val
}

func dfs(root *TreeNode, k *int, val *int) {
	if root.Left != nil {
		dfs(root.Left, k, val)
	}
	*k--
	if *k == 0 {
		*val = root.Val
		return
	}
	if root.Right != nil {
		dfs(root.Right, k, val)
	}
}

func main() {
	ts := []struct {
		root     *TreeNode
		k        int
		expected int
	}{
		{
			root:     &TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}},
			k:        1,
			expected: 1,
		},
		{
			root:     &TreeNode{5, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, nil}},
			k:        3,
			expected: 3,
		},
	}
	for _, tc := range ts {
		actual := kthSmallest(tc.root, tc.k)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.root, tc.k, tc.expected, actual)
		}
	}
}
