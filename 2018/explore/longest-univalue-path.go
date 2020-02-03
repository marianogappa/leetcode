package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(max(longestUnivaluePath(root.Left), longestUnivaluePath(root.Right)),
		dfs(root.Left, root.Val)+dfs(root.Right, root.Val))
}

func dfs(n *TreeNode, v int) int {
	if n == nil || n.Val != v {
		return 0
	}
	return 1 + max(dfs(n.Left, v), dfs(n.Right, v))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestUnivaluePath(&TreeNode{5, &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}}, &TreeNode{5, nil, &TreeNode{5, nil, nil}}}) == 2)
	fmt.Println(longestUnivaluePath(&TreeNode{1, &TreeNode{4, &TreeNode{4, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{5, nil, &TreeNode{5, nil, nil}}}) == 2)
}
