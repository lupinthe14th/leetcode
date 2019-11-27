package removeduplicatesfromsortedlist

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ln := &ListNode{}
	out := ln

	var curr int
	c := 0
	for head != nil {
		if curr != head.Val || c == 0 {
			ln.Next = &ListNode{}
			ln = ln.Next
			ln.Val = head.Val
			curr = head.Val
		}
		head = head.Next
		c++
	}
	return out.Next
}
