package main

import (
	"fmt"

	. "github.com/pencil001/leetcode/utils"
)

func main() {
	// head := NewList([]int{1, 2, 3, 2, 1})
	head := NewList([]int{1, 2})
	ok := isPalindrome(head)
	fmt.Println(ok)
}

func isPalindrome(head *ListNode) bool {
	c := []int{}
	p := head
	for p != nil {
		c = append(c, p.Val)
		p = p.Next
	}

	b := 0
	e := len(c) - 1
	for b <= e {
		if c[b] != c[e] {
			return false
		}
		b++
		e--
	}
	return true
}
