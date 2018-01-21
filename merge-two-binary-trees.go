package main

import "fmt"

//TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	return &TreeNode{Val: t1.Val + t2.Val, Left: mergeTrees(t1.Left, t2.Left), Right: mergeTrees(t1.Right, t2.Right)}
}

func (t *TreeNode) print() string {
	if t == nil {
		return ""
	}
	if t.Left == nil && t.Right == nil {
		return fmt.Sprintf("%v", t.Val)
	}
	return fmt.Sprintf("%v(%v,%v)", t.Val, t.Left.print(), t.Right.print())
}

func main() {
	var t1 = &TreeNode{1, &TreeNode{3, &TreeNode{5, nil, nil}, nil}, &TreeNode{2, nil, nil}}
	var t2 = &TreeNode{2, &TreeNode{1, nil, &TreeNode{4, nil, nil}}, &TreeNode{3, nil, &TreeNode{7, nil, nil}}}
	fmt.Println(t1.print())
	fmt.Println(t2.print())
	fmt.Println(mergeTrees(t1, t2).print())
}
