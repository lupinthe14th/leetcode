package populatingnextrightpointersineachnode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// SeeAlso:
// - https://leetcode.com/explore/learn/card/data-structure-tree/133/conclusion/994/discuss/37472/A-simple-accepted-solution
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var pre *Node = root
	var cur *Node = nil
	for pre.Left != nil {
		cur = pre
		for cur != nil {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			cur = cur.Next
		}
		pre = pre.Left
	}
	return root
}
