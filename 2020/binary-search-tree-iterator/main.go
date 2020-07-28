package main

import (
	"container/list"
	"fmt"
)

// The approach is straight up iterative DFS inorder traversal.
//
// On construction, push left nodes up to the first leftmost node
// into a stack.
//
// On Next, pop & return that leftmost node, but also check if a
// right node exists, and if so push it and push the leftmost
// path from it into the stack.
//
// HasNext simply checks the length of the stack.

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack *list.List
}

// Time: O(h) h == height of BST
// Space: O(h) h == height of BST
func Constructor(root *TreeNode) BSTIterator {
	l := list.New()
	if root != nil {
		l.PushFront(root)
	}

	it := BSTIterator{l}
	it.advanceLeft()

	return it
}

// Time: O(h) h == height of BST
// Space: O(h) h == height of BST
/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	curElem := this.stack.Front()
	this.stack.Remove(curElem)

	curNode := curElem.Value.(*TreeNode)
	if curNode.Right != nil {
		this.stack.PushFront(curNode.Right)
		this.advanceLeft()
	}

	return curNode.Val
}

// Time: O(1)
// Space: O(1)
/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.stack == nil {
		return false
	}
	return this.stack.Len() > 0
}

// Time: O(h) h == height of BST
// Space: O(h) h == height of BST
func (this *BSTIterator) advanceLeft() {
	if !this.HasNext() {
		return
	}
	curNode := this.stack.Front().Value.(*TreeNode)
	for curNode.Left != nil {
		this.stack.PushFront(curNode.Left)
		curNode = curNode.Left
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main() {
	tn := &TreeNode{10, &TreeNode{5, &TreeNode{4, nil, nil}, &TreeNode{6, nil, nil}}, &TreeNode{15, nil, nil}}
	nbi := Constructor(tn)
	bi := &nbi
	if bi.HasNext() {
		fmt.Println(bi.Next())
	}
	if bi.HasNext() {
		fmt.Println(bi.Next())
	}
	if bi.HasNext() {
		fmt.Println(bi.Next())
	}
	if bi.HasNext() {
		fmt.Println(bi.Next())
	}
	if bi.HasNext() {
		fmt.Println(bi.Next())
	}
	if bi.HasNext() {
		fmt.Println(bi.Next())
	}
	if bi.HasNext() {
		fmt.Println(bi.Next())
	}
	if bi.HasNext() {
		fmt.Println(bi.Next())
	}
	if bi.HasNext() {
		fmt.Println(bi.Next())
	}
}
