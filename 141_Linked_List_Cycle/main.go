package main

import (
	"fmt"

	. "github.com/pencil001/leetcode/utils"
)

func main() {
	head := &ListNode{Val: 1}
	ok := hasCycle(head)
	fmt.Println(ok)
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
