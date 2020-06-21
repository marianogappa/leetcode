package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("(%v,%v,%v)", n.Val, n.Left, n.Right)
}

// Time: O(n)
// Space: O(n)
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// Insight here is that, for a given node, the left
	// subtree nodes are to the left of the _inorder_
	// index of that node, and viceversa for the right
	// subtree nodes.
	//
	// Because tree is sparse, or at least probably not
	// perfectly balanced, we need constant access to the
	// inorder structure from a given value.
	valToIndex := make(map[int]int, len(inorder))
	for i, n := range inorder {
		valToIndex[n] = i
	}
	// The order of creation of nodes is gonna be preorder, so we keep
	// incrementing a _preorderIndex_ across recursive invocations.
	var preorderIndex int

	return doBuildTree(preorder, inorder, valToIndex, 0, len(inorder)-1, &preorderIndex)
}

func doBuildTree(preorder, inorder []int, valToIndex map[int]int, startInOrder, endInOrder int, preorderIndex *int) *TreeNode {
	// Nothing left to do on this subtree
	if startInOrder > endInOrder {
		return nil
	}
	// The current node is constructed with the value pointed at
	// by the "global" _preoorderIndex_.
	node := &TreeNode{preorder[*preorderIndex], nil, nil}
	inorderIndex := valToIndex[preorder[*preorderIndex]]
	*preorderIndex++

	// If the range of the inorder array contained only this value,
	// we're done with this subtree.
	if startInOrder == endInOrder {
		return node
	}

	// Otherwise, build the left and right subtrees on both sides of the _inorderIndex_.
	node.Left = doBuildTree(preorder, inorder, valToIndex, startInOrder, inorderIndex-1, preorderIndex)
	node.Right = doBuildTree(preorder, inorder, valToIndex, inorderIndex+1, endInOrder, preorderIndex)

	return node
}

func main() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	// ts := []struct {
	// 	input    int
	// 	expected int
	// }{
	// 	{
	// 		input:    1,
	// 		expected: 1,
	// 	},
	// }
	// for _, tc := range ts {
	// 	actual := 1
	// 	if tc.expected != actual {
	// 		fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
	// 	}
	// }
}
