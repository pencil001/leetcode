package p0236

import (
	. "github.com/pencil001/leetcode/utils"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}

	if l == nil {
		return r
	}

	return l
}

//leetcode submit region end(Prohibit modification and deletion)
