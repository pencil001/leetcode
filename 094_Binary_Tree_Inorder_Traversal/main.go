package main

import (
	"fmt"

	. "github.com/pencil001/leetcode/utils"
)

func main() {
	root := NewTree([]interface{}{1})
	fmt.Println(inorderTraversal2(root))
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	l := inorderTraversal(root.Left)
	v := root.Val
	r := inorderTraversal(root.Right)

	res := append(l, v)
	res = append(res, r...)
	return res
}

func inorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 1000)
	pos := -1
	visit(root, &stack, &pos)

	for pos != -1 {
		curr := stack[pos]
		pos--
		if curr.Left == nil && curr.Right == nil {
			res = append(res, curr.Val)
		} else {
			visit(curr, &stack, &pos)
		}
	}
	return res
}

func visit(root *TreeNode, stack *[]*TreeNode, pos *int) {
	l := root.Left
	r := root.Right
	if r != nil {
		(*pos)++
		(*stack)[*pos] = r
	}
	root.Left, root.Right = nil, nil
	(*pos)++
	(*stack)[*pos] = root
	if l != nil {
		(*pos)++
		(*stack)[*pos] = l
	}
}
