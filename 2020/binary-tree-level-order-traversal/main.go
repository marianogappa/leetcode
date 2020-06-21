package main

import (
	"fmt"
)

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
// Space: O(n) max width of tree + all nodes anyway (for result)
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	curLevelNodes := []*TreeNode{root}
	for len(curLevelNodes) > 0 {
		levelVals := []int{}
		nextLevelNodes := []*TreeNode{}

		// For the current level's nodes, save all
		// values on _levelVals_ and all next level
		// nodes (i.e. _.Left_ and _.Right_) on
		// _nextLevelNodes_.
		for _, node := range curLevelNodes {
			levelVals = append(levelVals, node.Val)
			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}

		result = append(result, levelVals)
		curLevelNodes = nextLevelNodes
	}
	return result
}

func main() {
	fmt.Println(levelOrder(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}))
	// ts := []struct {
	// 	input    *TreeNode
	// 	expected [][]int
	// }{
	// 	{
	// 		input:    ,
	// 		expected: [][]int{{3}, {9, 20}, {15, 7}},
	// 	},
	// }
	// for _, tc := range ts {
	// 	actual := levelOrder(tc.input)
	// 	if !reflect.DeepEqual(tc.expected, actual) {
	// 		fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
	// 	}
	// }
}
