package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
// Time: O(n)
// Space: O(n)
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var (
		strs  = []string{}
		queue = []*TreeNode{root}
	)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == nil {
			strs = append(strs, "n")
			continue
		} else {
			strs = append(strs, fmt.Sprintf("%v", cur.Val))
		}

		queue = append(queue, cur.Left, cur.Right)
	}

	return strings.Join(strs, ",")
}

// Deserializes your encoded data to tree.
// Time: O(n)
// Space: O(n)
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	var (
		vs    = strings.Split(data, ",")
		n, _  = strconv.Atoi(vs[0])
		root  = &TreeNode{n, nil, nil}
		queue = []*TreeNode{root}
		i     = 1
	)
	for len(queue) > 0 {
		cur := queue[0]

		if vs[i] != "n" {
			n, _ := strconv.Atoi(vs[i])
			cur.Left = &TreeNode{n, nil, nil}
			queue = append(queue, cur.Left)
		}

		if vs[i+1] != "n" {
			n, _ := strconv.Atoi(vs[i+1])
			cur.Right = &TreeNode{n, nil, nil}
			queue = append(queue, cur.Right)
		}

		queue = queue[1:]
		i += 2
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

func main() {

}
