package main

import (
	. "github.com/pencil001/leetcode/utils"
)

func main() {
	WalkTree(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	v := preorder[0]
	i := indexOf(inorder, v)
	if i == -1 {
		return nil
	}

	root := &TreeNode{Val: v}
	leftCount := i
	root.Left = buildTree(preorder[1:1+leftCount], inorder[:leftCount])
	root.Right = buildTree(preorder[1+leftCount:], inorder[i+1:])
	return root
}

func indexOf(array []int, k int) int {
	for i, v := range array {
		if v == k {
			return i
		}
	}
	return -1
}
