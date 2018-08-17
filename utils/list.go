package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(nums []int) *ListNode {
	dummy := &ListNode{Val: 0}

	link := dummy
	for _, v := range nums {
		n := &ListNode{Val: v}
		link.Next = n
		link = n
	}

	return dummy.Next
}

func WalkList(head *ListNode) {
	if head == nil {
		fmt.Println()
		return
	}

	fmt.Printf("%v;", head.Val)
	WalkList(head.Next)
}
