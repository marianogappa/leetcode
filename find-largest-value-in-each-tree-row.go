package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		res = []int{root.Val}
		ns  = []*TreeNode{root}
		max int
	)
	for len(ns) > 0 {
		max = -1 << 31
		var nns = make([]*TreeNode, 0)
		for _, n := range ns {
			if n.Left != nil {
				nns = append(nns, n.Left)
				if n.Left.Val > max {
					max = n.Left.Val
				}
			}
			if n.Right != nil {
				nns = append(nns, n.Right)
				if n.Right.Val > max {
					max = n.Right.Val
				}
			}
		}
		if len(nns) > 0 {
			res = append(res, max)
		}
		ns = nns
	}
	return res
}

func main() {
	var bt = &TreeNode{1, &TreeNode{3, &TreeNode{5, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{2, nil, &TreeNode{9, nil, nil}}}
	fmt.Println(largestValues(bt))
	bt = &TreeNode{1, nil, nil}
	fmt.Println(largestValues(bt))
	bt = &TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}}
	fmt.Println(largestValues(bt))
}
