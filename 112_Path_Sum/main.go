package main

import (
	"fmt"

	. "github.com/pencil001/leetcode/utils"
)

func main() {
	root := NewTree([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1})
	sum := 22
	ok := hasPathSum2(root, sum)
	fmt.Println(ok)
}

func hasPathSum(root *TreeNode, sum int) bool {
	return hasPathSumEx(root, sum, 0)
}

func hasPathSumEx(root *TreeNode, sum int, acc int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if acc+root.Val == sum {
			return true
		}
	}

	return hasPathSumEx(root.Left, sum, acc+root.Val) || hasPathSumEx(root.Right, sum, acc+root.Val)
}

func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return true
		}
	}

	return hasPathSum2(root.Left, sum-root.Val) || hasPathSum2(root.Right, sum-root.Val)
}
