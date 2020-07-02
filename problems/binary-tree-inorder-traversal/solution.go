package binarytreeinordertraversal

// TreeNode is binary tree node definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	out := []int{}

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if root == nil {
			return
		}
		if node.Left != nil {
			helper(node.Left)
		}
		out = append(out, node.Val)
		if node.Right != nil {
			helper(node.Right)
		}
	}

	helper(root)
	return out
}
