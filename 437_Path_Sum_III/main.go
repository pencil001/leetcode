package main

import (
	"fmt"

	. "github.com/pencil001/leetcode/utils"
)

func main() {
	// root := NewTree([]interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1})
	// sum := 8
	root := NewTree([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, nil, 5, 1})
	sum := 22
	n := pathSum2(root, sum)
	fmt.Println(n)
}

func pathSum2(root *TreeNode, sum int) int {
	dict := make(map[int]int)
	dict[0] = 1
	return pathSum2Ex(root, sum, 0, &dict)
}

func pathSum2Ex(root *TreeNode, target int, acc int, dict *map[int]int) int {
	if root == nil {
		return 0
	}

	acc += root.Val
	r := (*dict)[acc-target]
	(*dict)[acc]++

	r += pathSum2Ex(root.Left, target, acc, dict) + pathSum2Ex(root.Right, target, acc, dict)
	(*dict)[acc]--
	return r
}

func pathSum(root *TreeNode, sum int) int {
	paths = 0
	pathSumFromRoot(root, sum)
	return paths
}

var paths int

func pathSumFromRoot(root *TreeNode, sum int) {
	if root == nil {
		return
	}

	pathSumEx(root, sum, 0)
	pathSumFromRoot(root.Left, sum)
	pathSumFromRoot(root.Right, sum)
}

func pathSumEx(root *TreeNode, sum int, acc int) {
	if root == nil {
		return
	}

	acc += root.Val
	if acc == sum {
		paths++
	}

	pathSumEx(root.Left, sum, acc)
	pathSumEx(root.Right, sum, acc)
}
