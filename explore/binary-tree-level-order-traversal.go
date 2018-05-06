package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		q   = []*TreeNode{root}
		res = make([][]int, 0)
	)
	for len(q) > 0 {
		var (
			i     int
			l     = len(q)
			level = make([]int, 0)
		)
		for i = 0; i < l; i++ {
			level = append(level, q[i].Val)
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		res = append(res, level)
		q = q[i:]
	}
	return res
}

func main() {
	var t = &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(levelOrder(t))
}
