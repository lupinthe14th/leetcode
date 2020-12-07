package populatingnextrightpointersineachnodeii

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var head *Node = nil
	var pre *Node = nil
	var cur *Node = root

	for cur != nil {
		for cur != nil {
			// Left child
			if cur.Left != nil {
				if pre != nil {
					pre.Next = cur.Left
				} else {
					head = cur.Left
				}
				pre = cur.Left
			}
			// Right child
			if cur.Right != nil {
				if pre != nil {
					pre.Next = cur.Right
				} else {
					head = cur.Right
				}
				pre = cur.Right
			}
			cur = cur.Next
		}
		cur = head
		head = nil
		pre = nil
	}
	return root
}
