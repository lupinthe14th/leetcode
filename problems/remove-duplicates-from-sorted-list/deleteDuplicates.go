package removeduplicatesfromsortedlist

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	ln := &ListNode{}
	out := ln

	var curr int
	for head != nil {
		if curr != head.Val {
			ln.Next = &ListNode{}
			ln = ln.Next
			ln.Val = head.Val
			curr = head.Val
		}
		head = head.Next
	}
	return out.Next
}
