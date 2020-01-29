package main

import "fmt"

import "reflect"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Trivial recursion for in-order traversal. Careful here to
// use a pointer to the slice; bug otherwise.
// Homework: implement iterative solution rather than recursive.
//
// Time: O(n) because all nodes are traversed once
// Space: O(n) because worst-case n-1 nodes will be in the stack
func inorderTraversal(root *TreeNode) []int {
	var res = []int{}
	doInOrder(root, &res)
	return res
}

func doInOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	doInOrder(root.Left, res)
	*res = append(*res, root.Val)
	doInOrder(root.Right, res)
}

func main() {
	ts := []struct {
		input    *TreeNode
		expected []int
	}{
		{
			input:    &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}},
			expected: []int{1, 3, 2},
		},
	}
	for _, tc := range ts {
		actual := inorderTraversal(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
