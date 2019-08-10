package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return head
	}

	next := removeElements(head.Next, val)
	if head.Val == val {
		return next
	}

	head.Next = next

	return head
}

func main() {

}
