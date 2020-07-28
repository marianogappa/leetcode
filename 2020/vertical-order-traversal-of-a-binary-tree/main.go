package main

import (
	"fmt"
	"math"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type node struct {
	y, val int
}

// Time: O(nlog(n/k)) n == nodes in the tree; k == width of the tree
// Space: O(n)
//
// 1. Put all nodes into a hashmap using DFS keyed by column,
// and storing row and value for each.
//
// 2. Sort each resulting slice by row and value. Keep track of
// min and max columns.
//
// 3. Output a slice of slices ordered by column.
func verticalTraversal(root *TreeNode) [][]int {
	hmap := map[int][]node{}
	buildHashmapWithDfs(root, 0, 0, hmap)

	var (
		mn = math.MaxInt32
		mx = math.MinInt32
	)
	for x, nodes := range hmap {
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].y < nodes[j].y || (nodes[i].y == nodes[j].y && nodes[i].val < nodes[j].val)
		})
		mn = min(mn, x)
		mx = max(mx, x)
	}

	return hashmapToSlices(hmap, mn, mx)
}

func buildHashmapWithDfs(root *TreeNode, x, y int, hmap map[int][]node) {
	if root == nil {
		return
	}
	hmap[x] = append(hmap[x], node{y: y, val: root.Val})
	buildHashmapWithDfs(root.Left, x-1, y+1, hmap)
	buildHashmapWithDfs(root.Right, x+1, y+1, hmap)
}

func hashmapToSlices(hmap map[int][]node, mn, mx int) [][]int {
	results := [][]int{}
	for i := mn; i <= mx; i++ {
		result := []int{}
		for _, node := range hmap[i] {
			result = append(result, node.val)
		}
		results = append(results, result)
	}
	return results
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(verticalTraversal(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}))
}
