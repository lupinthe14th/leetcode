package flattenamultileveldoublylinkedlist

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// https://leetcode.com/explore/learn/card/linked-list/213/conclusion/1225/discuss/150321/Easy-Understanding-Java-beat-95.7-with-Explanation
func flatten(root *Node) *Node {
	if root == nil {
		return root
	}

	p := root
	for p != nil {
		// Case 1: if no child, proceed
		if p.Child == nil {
			p = p.Next
			continue
		}
		// Case 2: got child, find the tail of the child and link it to p.Next
		temp := p.Child
		// Find the tail of the child
		for temp.Next != nil {
			temp = temp.Next
		}
		// Connect tail with p.Next, if it is not nil
		temp.Next = p.Next
		if p.Next != nil {
			p.Next.Prev = temp
		}
		// Connect p with p.Cild, and remove p.Child
		p.Next = p.Child
		p.Child.Prev = p
		p.Child = nil
	}
	return root
}
