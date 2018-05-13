package main

import (
	"fmt"

	. "github.com/pencil001/leetcode/utils"
)

func main() {
	root := NewTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	fmt.Println(levelOrder(root))
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	levelOrderEx(root, 0, &res)
	return res
}

func levelOrderEx(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}

	if len(*res) <= level {
		*res = append(*res, []int{})
	}

	(*res)[level] = append((*res)[level], root.Val)
	levelOrderEx(root.Left, level+1, res)
	levelOrderEx(root.Right, level+1, res)
}
