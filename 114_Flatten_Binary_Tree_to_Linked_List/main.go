package main

import (
	. "github.com/pencil001/leetcode/utils"
)

func main() {
	root := NewTree([]interface{}{1, 2, 5, 3, 4, nil, 6})
	// root := NewTree([]interface{}{1, 2, 3})
	flatten(root)
	WalkTree(root)
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	tmpRight := root.Right

	if root.Left != nil {
		root.Right = root.Left
		if root.Right != nil {
			next := root.Right
			for next.Right != nil {
				next = next.Right
			}
			next.Right = tmpRight
		}
		root.Left = nil
	}
}
