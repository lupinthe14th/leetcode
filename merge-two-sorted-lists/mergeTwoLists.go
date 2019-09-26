package mergetwosortedlists

// ListNode is singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var ln = &ListNode{}
	var out = ln
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			ln.Next = l1
			l1 = l1.Next
		} else {
			ln.Next = l2
			l2 = l2.Next
		}
		ln = ln.Next
	}

	if l1 != nil {
		ln.Next = l1
	} else if l2 != nil {
		ln.Next = l2
	}
	return out.Next
}
