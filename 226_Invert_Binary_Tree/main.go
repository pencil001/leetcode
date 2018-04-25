package main

import . "github.com/pencil001/leetcode/utils"

func main() {
	root := NewTree([]interface{}{4, 2, 7, 1, 3, 6, 9})
	t := invertTree(root)
	WalkTree(t)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	node := &TreeNode{Val: root.Val}
	node.Left = invertTree(root.Right)
	node.Right = invertTree(root.Left)
	return node
}
