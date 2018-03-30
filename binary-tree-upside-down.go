package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root != nil && root.Left == nil {
		return &TreeNode{root.Val, root.Right, nil}
	}
	var (
		n = upsideDownBinaryTree(root.Left)
		p = n
	)
	for p != nil && p.Right != nil {
		p = p.Right
	}
	p.Left = root.Right
	p.Right = &TreeNode{root.Val, nil, nil}
	return n
}

func (t *TreeNode) string() string {
	if t == nil {
		return ""
	}
	if t != nil && t.Left == nil && t.Right == nil {
		return fmt.Sprintf("%v", t.Val)
	}
	return fmt.Sprintf("%v (%v, %v)", t.Val, t.Left.string(), t.Right.string())
}

func main() {
	// var bt = &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	// fmt.Println(bt.string())
	// fmt.Println(upsideDownBinaryTree(bt).string())

	// var bt = &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	// fmt.Println(bt.string())
	// fmt.Println(upsideDownBinaryTree(bt).string())

	// var bt = &TreeNode{1, nil, nil}
	// fmt.Println(bt.string())
	// fmt.Println(upsideDownBinaryTree(bt).string())

	// var bt *TreeNode
	// fmt.Println(bt.string())
	// fmt.Println(upsideDownBinaryTree(bt).string())

	// var bt = &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	// fmt.Println(bt.string())
	// fmt.Println(upsideDownBinaryTree(bt).string())

	var bt = &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	fmt.Println(bt.string())
	fmt.Println(upsideDownBinaryTree(bt).string())

}
