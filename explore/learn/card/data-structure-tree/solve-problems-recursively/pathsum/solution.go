package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	var rec func(root *TreeNode, acc int) bool
	rec = func(root *TreeNode, acc int) bool {
		if root == nil {
			return false
		}
		acc += root.Val
		if root.Left == nil && root.Right == nil {
			return acc == sum
		}
		return rec(root.Left, acc) || rec(root.Right, acc)
	}

	return rec(root, 0)
}
