package main

import (
	"fmt"

	. "github.com/pencil001/leetcode/utils"
)

func main() {
	// p := NewTree([]interface{}{1, 2, 3})
	// q := NewTree([]interface{}{1, 2, 3})
	p := NewTree([]interface{}{1, 2})
	q := NewTree([]interface{}{1, nil, 2})
	ok := isSameTree(p, q)
	fmt.Println(ok)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
