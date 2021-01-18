package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ln := &ListNode{}
	out := ln
	carry := 0
	for l1 != nil && l2 != nil {
		ln.Next = &ListNode{}
		ln = ln.Next
		ln.Val = (carry + l1.Val + l2.Val) % 10
		ln.Val = (carry + l1.Val + l2.Val) % 10
		carry = (carry + l1.Val + l2.Val) / 10
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		ln.Next = &ListNode{}
		ln = ln.Next
		ln.Val = (carry + l1.Val) % 10
		carry = (carry + l1.Val) / 10
		l1 = l1.Next
	}
	for l2 != nil {
		ln.Next = &ListNode{}
		ln = ln.Next
		ln.Val = (carry + l2.Val) % 10
		carry = (carry + l2.Val) / 10
		l2 = l2.Next
	}
	if carry != 0 {
		ln.Next = &ListNode{}
		ln = ln.Next
		ln.Val = carry
	}
	return out.Next
}
