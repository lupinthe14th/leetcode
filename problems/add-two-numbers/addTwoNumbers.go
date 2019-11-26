package main

// ListNode is singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ln := &ListNode{}
	out := ln

	for l1 != nil && l2 != nil {
		tmp := ln.Val + l1.Val + l2.Val
		ln.Val = tmp % 10
		if l1.Next != nil || l2.Next != nil {
			ln.Next = &ListNode{}
			ln = ln.Next
			if tmp >= 10 {
				ln.Val = 1
			}
		} else if tmp >= 10 {
			ln.Next = &ListNode{}
			ln = ln.Next
			ln.Val = 1
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		tmp := (ln.Val + l1.Val)
		ln.Val = tmp % 10
		if tmp >= 10 {
			ln.Next = &ListNode{}
			ln = ln.Next
			ln.Val = 1
		}
	} else if l2 != nil {
		tmp := (ln.Val + l2.Val)
		ln.Val = tmp % 10
		if tmp >= 10 {
			ln.Next = &ListNode{}
			ln = ln.Next
			ln.Val = 1
		}

	}
	return out
}
