package main

import (
	"fmt"

	. "github.com/pencil001/leetcode/utils"
)

func main() {
	// root := NewTree([]interface{}{1, 2, 2, 3, 4, 4, 3})
	root := NewTree([]interface{}{1, 2, 2, nil, 3, nil, 3})
	ok := isSymmetric(root)
	fmt.Println(ok)
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricEx(root, root)
}

func isSymmetricEx(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 == nil || r2 == nil {
		return false
	}

	if r1.Val != r2.Val {
		return false
	}

	b1 := isSymmetricEx(r1.Left, r2.Right)
	b2 := isSymmetricEx(r1.Right, r2.Left)
	return b1 && b2
}
