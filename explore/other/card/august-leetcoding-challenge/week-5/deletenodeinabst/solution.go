package deletenodeinabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SeeAlso: https://leetcode.com/explore/challenge/card/august-leetcoding-challenge/553/week-5-august-29th-august-31st/3443/discuss/93296/Recursive-Easy-to-Understand-Java-Solution
func deleteNode(root *TreeNode, key int) *TreeNode {

	findMin := func(node *TreeNode) *TreeNode {
		for node.Left != nil {
			node = node.Left
		}
		return node
	}
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		minNode := findMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, root.Val)
	}
	return root
}
