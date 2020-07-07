package binarytreepreordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	out := make([]int, 0)

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		out = append(out, node.Val)
		if node.Left != nil {
			helper(node.Left)
		}
		if node.Right != nil {
			helper(node.Right)
		}
	}

	helper(root)
	return out
}
