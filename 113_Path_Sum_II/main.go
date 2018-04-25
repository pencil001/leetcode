package main

import (
	"fmt"

	. "github.com/pencil001/leetcode/utils"
)

func main() {
	root := NewTree([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, nil, 5, 1})
	sum := 22
	paths := pathSum(root, sum)
	fmt.Println(paths)
}

var ps [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	ps = make([][]int, 0)
	p := []int{}
	pathSumEx(root, sum, &p)
	return ps
}

func pathSumEx(root *TreeNode, sum int, path *[]int) {
	if root == nil {
		return
	}

	*path = append(*path, root.Val)
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			length := len(*path)
			copyPath := make([]int, length)
			copy(copyPath, *path)
			ps = append(ps, copyPath)
			*path = (*path)[:length-1]
			return
		}
	}

	pathSumEx(root.Left, sum-root.Val, path)
	pathSumEx(root.Right, sum-root.Val, path)
	length := len(*path)
	*path = (*path)[:length-1]
}
