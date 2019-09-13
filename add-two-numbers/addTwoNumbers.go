package main

import (
	"log"
)

// ListNode is singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	for n := l1.Next; n != nil; n = n.Next {
		log.Printf("Val: %d", n.Val)
		log.Printf("Next: %#v", n.Next)
	}

	return &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}}
}
