package utils

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(nums []interface{}) *TreeNode {
	var newTreeEx func([]interface{}, int) *TreeNode
	newTreeEx = func(nums []interface{}, rootIdx int) *TreeNode {
		if rootIdx >= len(nums) {
			return nil
		}

		v := nums[rootIdx]
		if v == nil {
			return nil
		}

		n := &TreeNode{Val: v.(int)}
		n.Left = newTreeEx(nums, 2*rootIdx+1)
		n.Right = newTreeEx(nums, 2*rootIdx+2)
		return n
	}

	return newTreeEx(nums, 0)
}

func WalkTree(t *TreeNode) {
	if t == nil {
		return
	}

	WalkTree(t.Left)
	fmt.Printf("%v;", t.Val)
	WalkTree(t.Right)
}
