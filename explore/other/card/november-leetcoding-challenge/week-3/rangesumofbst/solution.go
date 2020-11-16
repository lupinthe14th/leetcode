package rangesumofbst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	out := 0
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		if low <= node.Val && node.Val <= high {
			out += node.Val
		}
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return out
}
