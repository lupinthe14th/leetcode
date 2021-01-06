package removeduplicatesfromsortedlistii

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		ln := head
		cur := head
		for cur != nil && ln.Val == cur.Val {
			cur = cur.Next
		}
		return deleteDuplicates(cur)
	} else {
		head.Next = deleteDuplicates(head.Next)
	}
	return head
}
