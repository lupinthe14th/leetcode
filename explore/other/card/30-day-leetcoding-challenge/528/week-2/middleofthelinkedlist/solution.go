package middleofthelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	c := 0
	out := head
	for head != nil {
		head = head.Next
		c++
	}
	m := c / 2
	for i := 0; i < m; i++ {
		out = out.Next
	}
	return out
}
