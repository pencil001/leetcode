package main

import (
	"fmt"

	. "github.com/pencil001/leetcode/utils"
)

func main() {
	// s := NewTree([]interface{}{3, 4, 5, 1, 2})
	// t := NewTree([]interface{}{4, 1, 2})
	s := NewTree([]interface{}{3, 4, 5, 1, nil, 2})
	t := NewTree([]interface{}{3, 1, 2})
	ok := isSubtree(s, t)
	fmt.Println(ok)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if isSameTree(s, t) {
		return true
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}
