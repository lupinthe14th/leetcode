package binarytreeinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	out := make([]int, 0)

	var helper func(node *TreeNode)

	helper = func(node *TreeNode) {
		if node == nil {
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
