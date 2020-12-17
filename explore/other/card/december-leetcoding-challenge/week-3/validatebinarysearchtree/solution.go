package validatebinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	const (
		MIN = -1 << 63
		MAX = 1<<63 - 1
	)
	var helper func(node *TreeNode, l, h int64) bool

	helper = func(node *TreeNode, l, h int64) bool {
		if node == nil {
			return true
		}
		if (l != MIN && int64(node.Val) <= l) || (h != MAX && int64(node.Val) >= h) {
			return false
		}
		return helper(node.Left, l, int64(node.Val)) && helper(node.Right, int64(node.Val), h)
	}
	return helper(root, MIN, MAX)

}
