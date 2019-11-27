package main

// ListNode is singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ln := &ListNode{}
	out := ln

	var carry int
	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}
		sum := carry + v1 + v2
		carry = sum / 10
		ln.Val = sum % 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 != nil || l2 != nil {
			ln.Next = &ListNode{}
			ln = ln.Next
		}
	}

	if carry > 0 {
		ln.Next = &ListNode{}
		ln = ln.Next
		ln.Val = carry

	}

	return out
}
