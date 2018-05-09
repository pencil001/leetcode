package main

import (
	"fmt"

	. "github.com/pencil001/leetcode/utils"
)

func main() {
	// root := NewTree([]interface{}{3, 2, 3, nil, 3, nil, 1})
	root := NewTree([]interface{}{3, 4, 5, 1, 3, nil, 1})
	fmt.Println(rob3(root))
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	val := 0

	if root.Left != nil {
		val += rob(root.Left.Left) + rob(root.Left.Right)
	}

	if root.Right != nil {
		val = rob(root.Right.Left) + rob(root.Right.Right)
	}

	return max(rob(root.Left)+rob(root.Right), root.Val+val)
}

func rob2(root *TreeNode) int {
	hash := make(map[*TreeNode]int)
	return rob2Ex(root, &hash)
}

func rob2Ex(root *TreeNode, hash *map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if v, ok := (*hash)[root]; ok {
		return v
	}

	val := 0

	if root.Left != nil {
		val += rob2Ex(root.Left.Left, hash) + rob2Ex(root.Left.Right, hash)
	}

	if root.Right != nil {
		val = rob2Ex(root.Right.Left, hash) + rob2Ex(root.Right.Right, hash)
	}

	res := max(rob2Ex(root.Left, hash)+rob2Ex(root.Right, hash), root.Val+val)
	(*hash)[root] = res
	return res
}

func rob3(root *TreeNode) int {
	res1, res2 := rob3Ex(root)
	return max(res1, res2)
}

// return value: first is the max profile with root robbed; seconde is the max profile with root unrobbed
func rob3Ex(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	l0, l1 := rob3Ex(root.Left)
	r0, r1 := rob3Ex(root.Right)
	robbed := l1 + r1 + root.Val
	unrobbed := max(l0, l1) + max(r0, r1)
	return robbed, unrobbed
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}
