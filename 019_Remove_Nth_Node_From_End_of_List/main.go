package main

import (
	. "github.com/pencil001/leetcode/utils"
)

func main() {
	WalkList(removeNthFromEnd2(NewList([]int{1, 2, 3, 4, 5}), 2))
	WalkList(removeNthFromEnd2(NewList([]int{1, 2, 3, 4, 5}), 1))
	WalkList(removeNthFromEnd2(NewList([]int{1, 2}), 2))
	WalkList(removeNthFromEnd2(NewList([]int{1}), 1))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	vp := &ListNode{Next: head}
	var p0 *ListNode
	p1 := vp
	p2 := vp.Next
	tail := vp
	for i := 0; i < n; i++ {
		tail = tail.Next
	}

	for tail != nil {
		tail = tail.Next
		p0 = p1
		p2 = p2.Next
		p1 = p1.Next
	}

	p0.Next = p2
	return vp.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	vp := &ListNode{Next: head}
	front, back := vp, vp
	for i := 0; i <= n; i++ {
		back = back.Next
	}

	for back != nil {
		front = front.Next
		back = back.Next
	}

	front.Next = front.Next.Next
	return vp.Next
}
