package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// My solution is to traverse the binary tree constructing an array,
// heapify the array and pop min k times. It's not ideal because it
// completely ignores the information provided by the binary tree,
// but I couldn't find any better solution.
//
// Time: O(klogn)
// Space: O(n)
func kthSmallest(root *TreeNode, k int) int {
	// 1: traverse tree (dfs) and construct an array
	// This step is not ideal, because the binary tree
	// information is lost.
	// Time: O(n)
	// Space: O(n)
	arr := []int{}
	treeToArray(root, &arr)

	// 2: heapify slice
	// Time: O(n)
	// Space: O(1)
	heapify(arr)

	// 3: pop min k times
	// Time: O(klogn)
	// Space: O(1)
	var kth int
	for i := 0; i < k; i++ {
		kth = popMin(&arr)
	}

	return kth
}

// Time: O(n)
// Space: O(n)
func treeToArray(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	*arr = append(*arr, root.Val)
	treeToArray(root.Left, arr)
	treeToArray(root.Right, arr)
}

// Time: O(n)
// Space: O(1)
func heapify(arr []int) {
	for i := (len(arr) / 2) - 1; i >= 0; i-- {
		siftDown(arr, i)
	}
}

// Time: O(logn)
// Space: O(1)
func siftDown(arr []int, i int) {
	for root := i; root*2+1 < len(arr); { // while root is a parent
		var child = root*2 + 1                               // child = left child of root
		if child+1 < len(arr) && arr[child] > arr[child+1] { // child = max sibling
			child++
		}
		if arr[root] > arr[child] { // if root is unordered to child
			arr[root], arr[child] = arr[child], arr[root] // swap them
			root = child                                  // continue algorithm from child
		} else {
			return // if root and child are ordered, we're done
		}
	}
}

// Time: O(logn)
// Space: O(1)
func popMin(arr *[]int) int {
	max := (*arr)[0]
	(*arr)[0], (*arr)[len(*arr)-1] = (*arr)[len(*arr)-1], (*arr)[0]
	(*arr) = (*arr)[:len(*arr)-1]
	siftDown(*arr, 0)
	return max
}

func main() {
	ts := []struct {
		input    *TreeNode
		k        int
		expected int
	}{
		{
			input:    &TreeNode{1, nil, nil},
			k:        1,
			expected: 1,
		},
		{
			input:    &TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}},
			k:        1,
			expected: 1,
		},
		{
			input:    &TreeNode{5, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, nil}},
			k:        3,
			expected: 3,
		},
	}
	for _, tc := range ts {
		actual := kthSmallest(tc.input, tc.k)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
