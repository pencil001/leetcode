package main

import . "github.com/pencil001/leetcode/utils"

func main() {
	root := NewTree([]interface{}{2, 0, 3, -4, 1})
	t := convertBST(root)
	WalkTree(t)
}

func convertBST(root *TreeNode) *TreeNode {
	acc := 0

	var merge func(*TreeNode)
	merge = func(t *TreeNode) {
		if t == nil {
			return
		}

		merge(t.Right)
		t.Val += acc
		acc = t.Val
		merge(t.Left)
	}
	merge(root)
	return root
}
