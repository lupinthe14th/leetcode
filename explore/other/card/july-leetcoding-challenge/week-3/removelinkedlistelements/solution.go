package removelinkedlistelements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	node := &ListNode{}
	out := node

	for head != nil {
		if head.Val != val {
			node.Next = &ListNode{}
			node = node.Next
			node.Val = head.Val
		}
		head = head.Next
	}
	return out.Next
}
