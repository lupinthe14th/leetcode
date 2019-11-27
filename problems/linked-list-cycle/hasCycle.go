package linkedlistcycle

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// If there is no cycle in the list, the fast pointer will eventually reach the end and we can return false in this case.
	if head == nil || head.Next == nil {
		return false
	}

	//
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
