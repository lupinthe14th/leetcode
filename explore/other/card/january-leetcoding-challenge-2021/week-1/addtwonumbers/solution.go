package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ln := &ListNode{}
	out := ln
	pre := 0
	for l1 != nil && l2 != nil {
		ln.Next = &ListNode{}
		ln = ln.Next
		ln.Val = (pre + l1.Val + l2.Val) % 10
		ln.Val = (pre + l1.Val + l2.Val) % 10
		pre = (pre + l1.Val + l2.Val) / 10
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		ln.Next = &ListNode{}
		ln = ln.Next
		ln.Val = (pre + l1.Val) % 10
		pre = (pre + l1.Val) / 10
		l1 = l1.Next
	}
	for l2 != nil {
		ln.Next = &ListNode{}
		ln = ln.Next
		ln.Val = (pre + l2.Val) % 10
		pre = (pre + l2.Val) / 10
		l2 = l2.Next
	}
	if pre != 0 {
		ln.Next = &ListNode{}
		ln = ln.Next
		ln.Val = pre
	}
	return out.Next
}
