package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isEqual(root.Left, root.Right)
}

func isEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if (a == nil && b != nil) || (a != nil && b == nil) || (a.Val != b.Val) {
		return false
	}
	return isEqual(a.Left, b.Right) && isEqual(a.Right, b.Left)
}

func main() {
	fmt.Println(isSymmetric(&TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}))
	fmt.Println(isSymmetric(&TreeNode{1, &TreeNode{2, nil, &TreeNode{3, nil, nil}}, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}))
	// var ts = []struct {
	// 	i *TreeNode
	// 	e bool
	// }{
	// 	{

	//      e: false
	//      },
	// }
}
