package main

import "fmt"

// TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var r = make([][]int, 0)
	doLevelOrder([]*TreeNode{root}, &r)
	return r
}

func doLevelOrder(children []*TreeNode, r *[][]int) {
	// if len(children) == 0 {
	// 	return
	// }
	var (
		vs = make([]int, 0)
		ns = make([]*TreeNode, 0)
	)
	for _, n := range children {
		if n == nil {
			continue
		}
		vs = append(vs, n.Val)
		ns = append(ns, n.Left, n.Right)
	}
	if len(vs) > 0 {
		*r = append(*r, vs)
		doLevelOrder(ns, r)
	}
}

func main() {
	fmt.Println(levelOrder(nil))
	fmt.Println(levelOrder(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}))
	fmt.Println(levelOrder(&TreeNode{1, nil, &TreeNode{3, nil, nil}}))
	fmt.Println(levelOrder(&TreeNode{1, nil, &TreeNode{3, &TreeNode{4, nil, &TreeNode{6, nil, &TreeNode{7, nil, nil}}}, &TreeNode{5, nil, nil}}}))
}
