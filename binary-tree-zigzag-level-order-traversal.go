package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) because each node is evaluated only once
// Space: O(n) because we only keep two levels of the tree at a time in memory
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		res         = [][]int{[]int{root.Val}} // bootstrap root level
		ns          = []*TreeNode{root}
		leftToRight = false
	)
	for len(ns) > 0 { // bfs until a level has no nodes
		var (
			vs  = make([]int, 0)
			nns = make([]*TreeNode, 0)
		)
		for _, n := range ns { // append all child nodes of this level's nodes left to right
			if n.Left != nil {
				nns = append(nns, n.Left)
			}
			if n.Right != nil {
				nns = append(nns, n.Right)
			}
		}
		if leftToRight { // append current level's values left to right
			for _, n := range nns {
				vs = append(vs, n.Val)
			}
		} else { // append current level's values right to left
			for i := len(nns) - 1; i >= 0; i-- {
				vs = append(vs, nns[i].Val)
			}
		}
		if len(vs) > 0 { // append level's values to result
			res = append(res, vs)
		}
		ns = nns
		leftToRight = !leftToRight // for next level, reverse iteration order
	}
	return res
}

func main() {
	var bt = &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(zigzagLevelOrder(bt))
}
