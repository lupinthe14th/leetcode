package sumofleftleaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	out := 0

	var helper func(node *TreeNode, isLeft bool)
	helper = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}
		if node.Left != nil {
			helper(node.Left, true)
		}
		if node.Right != nil {
			helper(node.Right, false)
		}
		if node.Left == nil && node.Right == nil && isLeft {
			out += node.Val
		}
	}
	helper(root, false)
	return out
}
