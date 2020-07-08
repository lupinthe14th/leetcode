package binarytreepostordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	out := make([]int, 0)

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			helper(node.Left)
		}
		if node.Right != nil {
			helper(node.Right)
		}
		out = append(out, node.Val)
	}

	helper(root)
	return out
}
