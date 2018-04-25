package main

import (
	"fmt"

	. "github.com/pencil001/leetcode/utils"
)

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	depth := maxDepth(root)
	fmt.Print(depth)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	re := l
	if r > l {
		re = r
	}
	return re + 1
}
