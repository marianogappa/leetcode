package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return bfs([]*TreeNode{root}, []int{})
}

func bfs(ns []*TreeNode, rs []int) []int {
	if len(ns) == 0 {
		return rs
	}
	rs = append(rs, ns[len(ns)-1].Val)
	var nns = make([]*TreeNode, 0)
	for _, n := range ns {
		if n.Left != nil {
			nns = append(nns, n.Left)
		}
		if n.Right != nil {
			nns = append(nns, n.Right)
		}
	}
	return bfs(nns, rs)
}

func (t *TreeNode) print() string {
	if t == nil {
		return ""
	}
	return fmt.Sprintf("%v(%v,%v)", t.Val, t.Left.print(), t.Right.print())
}

func main() {
	var t = &TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, &TreeNode{4, nil, nil}}}
	fmt.Println(t.print(), rightSideView(t))
	var t2 = &TreeNode{1, &TreeNode{2, nil, &TreeNode{5, &TreeNode{8, nil, nil}, nil}}, &TreeNode{3, nil, nil}}
	fmt.Println(t2.print(), rightSideView(t2))
	var t3 = &TreeNode{1, nil, nil}
	fmt.Println(t3.print(), rightSideView(t3))
	var t4 *TreeNode
	fmt.Println(t4.print(), rightSideView(t4))
}
