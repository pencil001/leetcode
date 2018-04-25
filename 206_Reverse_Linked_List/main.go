package main

import (
	. "github.com/pencil001/leetcode/utils"
)

func main() {
	head := NewList([]int{1, 2, 3, 4})
	r := reverseList(head)
	WalkList(r)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre *ListNode
	cur := head

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
