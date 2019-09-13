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

	log.Printf("1: %d", (l1.Val + l2.Val))
	log.Printf("Next: %d", l1.Next)
	log.Printf("Next: Val %d", l1.Next.Val)

	return &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}}
}
