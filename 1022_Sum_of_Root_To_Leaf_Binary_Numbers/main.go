package main

import (
	"fmt"
	"strconv"

	. "github.com/pencil001/leetcode/utils"
)

func main() {
	// fmt.Println(sumRootToLeaf(NewTree([]interface{}{1, 0, 1})))
	fmt.Println(sumRootToLeaf2(NewTree([]interface{}{1, 0, 1, 0, 1, 0, 1})))
}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := [][]int{}
	combine(root, []int{}, &res)
	sum := 0
	for _, p := range res {
		s := ""
		for _, v := range p {
			s += strconv.Itoa(v)
		}
		i, _ := strconv.ParseInt(s, 2, 32)
		sum += int(i)
	}
	return sum
}

func combine(root *TreeNode, path []int, res *[][]int) {
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res, append([]int{}, path...))
		return
	}

	l := len(path)
	if root.Left != nil {
		combine(root.Left, path, res)
		path = path[:l]
	}
	if root.Right != nil {
		combine(root.Right, path, res)
		path = path[:l]
	}
}

// ========================================================
// 更好的解决方案

func sumRootToLeaf2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return combine2(0, root)
}

func combine2(upper int, root *TreeNode) int {
	v := (upper << 1) + root.Val
	if root.Left == nil && root.Right == nil {
		return v
	}

	res := 0
	if root.Left != nil {
		res += combine2(v, root.Left)
	}
	if root.Right != nil {
		res += combine2(v, root.Right)
	}
	return res
}
