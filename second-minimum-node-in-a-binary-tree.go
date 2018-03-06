package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	var _, m2 = dfs(root, -1, -1)
	return m2
}

func dfs(root *TreeNode, m1, m2 int) (int, int) {
	if root == nil {
		return m1, m2
	}
	if m1 == -1 || root.Val < m1 {
		m1, m2 = root.Val, m1
	} else if root.Val > m1 && (m2 == -1 || root.Val < m2) {
		m2 = root.Val
	}
	m1, m2 = dfs(root.Left, m1, m2)
	return dfs(root.Right, m1, m2)
}

func main() {
	var bt = &TreeNode{2, &TreeNode{2, nil, nil}, &TreeNode{5, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(findSecondMinimumValue(bt))
}
