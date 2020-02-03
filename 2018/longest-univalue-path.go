package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var rootLen, leftLen, rightLen int
	rootLen = getLen(root.Left, root.Val) + getLen(root.Right, root.Val)
	leftLen = longestUnivaluePath(root.Left)
	rightLen = longestUnivaluePath(root.Right)
	return max(rootLen, max(leftLen, rightLen))
}

func getLen(r *TreeNode, v int) int {
	if r == nil || r.Val != v {
		return 0
	}
	return 1 + max(getLen(r.Left, v), getLen(r.Right, v))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestUnivaluePath(
		&TreeNode{1,
			&TreeNode{4,
				&TreeNode{4,
					nil, nil,
				},
				&TreeNode{4,
					nil, nil,
				},
			},
			&TreeNode{5,
				nil,
				&TreeNode{5,
					nil, nil,
				},
			},
		},
	))
	fmt.Println(longestUnivaluePath(
		&TreeNode{5,
			&TreeNode{4,
				&TreeNode{1,
					nil, nil,
				},
				&TreeNode{1,
					nil, nil,
				},
			},
			&TreeNode{5,
				nil,
				&TreeNode{5,
					nil, nil,
				},
			},
		},
	))
	fmt.Println(longestUnivaluePath(
		&TreeNode{1,
			&TreeNode{1,
				&TreeNode{1,
					nil, nil,
				},
				&TreeNode{1,
					&TreeNode{1, nil, nil}, nil,
				},
			},
			&TreeNode{1,
				&TreeNode{1,
					nil, nil,
				},
				&TreeNode{1,
					nil, nil,
				},
			},
		},
	))
}
