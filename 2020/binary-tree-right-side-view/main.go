package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) every node visited once
// Space: O(n) in the worst case of a fully imbalanced tree
func rightSideView(root *TreeNode) []int {
	result := []int{}
	dfsPreorder(root, 0, &result)
	return result
}

func dfsPreorder(node *TreeNode, level int, result *[]int) {
	if node == nil {
		return
	}
	if level == len(*result) {
		*result = append(*result, node.Val)
	}
	dfsPreorder(node.Right, level+1, result)
	dfsPreorder(node.Left, level+1, result)
}

func main() {
	ts := []struct {
		input    *TreeNode
		expected []int
	}{
		{
			input:    &TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, &TreeNode{4, nil, nil}}},
			expected: []int{1, 3, 4},
		},
	}
	for _, tc := range ts {
		actual := rightSideView(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
