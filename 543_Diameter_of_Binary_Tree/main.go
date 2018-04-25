package main

import (
	"fmt"

	. "github.com/pencil001/leetcode/utils"
)

func main() {
	root := NewTree([]interface{}{1, 2, 3, 4, 5})
	// root := NewTree([]interface{}{3})
	r := diameterOfBinaryTree(root)
	fmt.Println(r)
}

var ans = 0

func diameterOfBinaryTree(root *TreeNode) int {
	ans = 0
	walk(root)
	return ans
}

func walk(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := walk(root.Left)
	r := walk(root.Right)
	if l+r > ans {
		ans = l + r
	}

	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}
