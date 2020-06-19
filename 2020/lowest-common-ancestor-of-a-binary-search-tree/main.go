package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) if BST is not balanced
// Space: O(h) max height of tree
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	switch {
	case p.Val < root.Val && q.Val < root.Val:
		return lowestCommonAncestor(root.Left, p, q)
	case p.Val > root.Val && q.Val > root.Val:
		return lowestCommonAncestor(root.Right, p, q)
	default:
		return root
	}
}

func main() {

}
