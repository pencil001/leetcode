package main

import (
	. "github.com/pencil001/leetcode/utils"
)

func main() {
	// l1 := NewList([]int{1, 2, 4})
	// l2 := NewList([]int{1, 3, 4})
	l1 := NewList([]int{})
	l2 := NewList([]int{})
	l := mergeTwoLists(l1, l2)
	WalkList(l)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	current := &ListNode{Val: 0}
	head := current
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			current.Next = p1
			p1 = p1.Next
		} else {
			current.Next = p2
			p2 = p2.Next
		}
		current = current.Next
	}

	if p1 != nil {
		current.Next = p1
	}
	if p2 != nil {
		current.Next = p2
	}
	return head.Next
}
