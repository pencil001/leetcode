package main

import (
	. "github.com/pencil001/leetcode/utils"
)

func main() {
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	tree2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	tree := mergeTrees(tree1, tree2)
	WalkTree(tree)
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	node := &TreeNode{}
	if t1 != nil && t2 != nil {
		node.Val = t1.Val + t2.Val
		node.Left = mergeTrees(t1.Left, t2.Left)
		node.Right = mergeTrees(t1.Right, t2.Right)
	} else if t1 != nil {
		node.Val = t1.Val
		node.Left = mergeTrees(t1.Left, nil)
		node.Right = mergeTrees(t1.Right, nil)
	} else {
		node.Val = t2.Val
		node.Left = mergeTrees(nil, t2.Left)
		node.Right = mergeTrees(nil, t2.Right)
	}
	return node
}
