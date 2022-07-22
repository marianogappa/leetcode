package main

import (
	"fmt"
	"strings"
)

// Time: O(n) as all steps are linear time.
// Space: O(n) for recursion stack because we don't know if tree is balanced.
//
// This is not a binary search tree so we cannot make any assumptions about location of values based on tree structure.
//
// First intuition when reading "shortest path" is Dijkstra, which would suggest building a graph from the tree,
// annotating edges with the letters and then running Dijkstra on it. But the time complexity doesn't look optimal.
//
// Second intuition is that it doesn't seem like there would be two paths to any src->dst in a tree (outside of doing
// something like SRC->A->B->A->B->DST). Any path is the shortest path.
//
// Candidate optimal strategy should then be finding the closest common ancestor, and then concatenate the descriptions
// towards (but not including) ancestor for src, and from (but not including) ancestor to dst.
func getDirections(root *TreeNode, startValue int, destValue int) string {
	var (
		pathToStart         = annotatePath(root, startValue, []byte{})
		pathToDest          = annotatePath(root, destValue, []byte{})
		idx                 = idxAfterLastCommonAncestor(pathToStart, pathToDest)
		directionsFromStart = strings.Repeat("U", len(pathToStart[idx:])) // from start -> last common ancestor we go up!
		directionsToDest    = pathToDest[idx:]
	)

	return string(directionsFromStart) + string(directionsToDest)
}

func idxAfterLastCommonAncestor(path1, path2 []byte) int {
	minLen := min(len(path1), len(path2))
	// Either we hit a non-common ancestor here, or one of (src,dst) is the last common ancestor
	for i := 0; i < minLen; i++ {
		if path1[i] != path2[i] {
			return i
		}
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// We're guaranteed to find a path from root to any supplied value
func annotatePath(root *TreeNode, value int, partial []byte) []byte {
	if root.Val == value {
		return partial
	}
	partial = append(partial, '?')
	if root.Left != nil {
		partial[len(partial)-1] = 'L'
		if res := annotatePath(root.Left, value, partial); res != nil {
			return res
		}
	}
	if root.Right != nil {
		partial[len(partial)-1] = 'R'
		if res := annotatePath(root.Right, value, partial); res != nil {
			return res
		}
	}
	return nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	ts := []struct {
		root                  *TreeNode
		startValue, destValue int
		expected              string
	}{
		{
			root:       makeTreeFromArray([]interface{}{5, 1, 2, 3, nil, 6, 4}),
			startValue: 3,
			destValue:  6,
			expected:   "UURL",
		},
		{
			root:       makeTreeFromArray([]interface{}{2, 1}),
			startValue: 2,
			destValue:  1,
			expected:   "L",
		},
		{
			root:       makeTreeFromArray([]interface{}{1, 3, 8, 7, nil, 4, 5, 6, nil, nil, nil, nil, nil, nil, 2}),
			startValue: 2,
			destValue:  1,
			expected:   "UUUU",
		},
	}
	for _, tc := range ts {
		actual := getDirections(tc.root, tc.startValue, tc.destValue)
		if tc.expected != actual {
			fmt.Printf("For (%v, %v) expected [%v] of len %v but got [%v] of len %v\n", tc.startValue, tc.destValue, tc.expected, len(tc.expected), actual, len(actual))
		}
	}
}

func makeTreeFromArray(arr []interface{}) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	root := &TreeNode{Val: arr[0].(int)}
	level := []*TreeNode{root}
	for i := 1; i < len(arr); {
		nextLevel := []*TreeNode{}
		for _, node := range level {
			if i >= len(arr) {
				break
			}
			if arr[i] != nil {
				node.Left = &TreeNode{Val: arr[i].(int)}
				nextLevel = append(nextLevel, node.Left)
			}
			i++
			if i >= len(arr) {
				break
			}
			if arr[i] != nil {
				node.Right = &TreeNode{Val: arr[i].(int)}
				nextLevel = append(nextLevel, node.Right)
			}
			i++
		}
		level = nextLevel
	}
	return root
}
