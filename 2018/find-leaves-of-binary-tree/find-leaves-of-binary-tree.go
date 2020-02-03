package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeaves(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res = make([][]int, 0)
	for root.Left != nil || root.Right != nil {
		res = append(res, dfs(root))
	}
	res = append(res, []int{root.Val})
	return res
}

func dfs(r *TreeNode) []int {
	if r == nil || (r.Left == nil && r.Right == nil) {
		return []int{}
	}
	var res = make([]int, 0)
	if r.Left != nil && r.Left.Left == nil && r.Left.Right == nil {
		res = append(res, r.Left.Val)
		r.Left = nil
	}
	res = append(res, dfs(r.Left)...)
	if r.Right != nil && r.Right.Left == nil && r.Right.Right == nil {
		res = append(res, r.Right.Val)
		r.Right = nil
	}
	res = append(res, dfs(r.Right)...)
	return res
}

func main() {
	var bt = &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	fmt.Println(findLeaves(bt))
}
