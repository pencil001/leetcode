package main

import (
	. "github.com/pencil001/leetcode/utils"
)

func main() {
	// head := NewList([]int{1, 2, 3, 4, 5})
	// r := reverseBetween2(head, 2, 4)

	head := NewList([]int{3, 5})
	r := reverseBetween2(head, 1, 2)
	WalkList(r)
}

func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Val: 0, Next: head}
	pre := dummy
	for i := 1; i < m; i++ {
		pre = pre.Next
	}

	start := pre.Next
	then := start
	if then != nil {
		then = then.Next
	}
	for i := 0; i < n-m; i++ {
		start.Next = then.Next
		then.Next = pre.Next
		pre.Next = then
		then = start.Next
	}

	return dummy.Next
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	var mprev *ListNode
	mcur := head
	for i := 2; i <= m; i++ {
		tmp := mcur.Next
		mprev = mcur
		mcur = tmp
	}

	var nprev *ListNode
	ncur := head
	for i := 1; i <= n; i++ {
		tmp := ncur.Next
		nprev = ncur
		ncur = tmp
	}

	if mprev != nil {
		mprev.Next = nil
	}
	nprev.Next = nil
	subHead := reverseList(mcur)
	if mprev != nil {
		mprev.Next = nprev
	} else {
		head = subHead
	}
	mcur.Next = ncur
	return head
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
