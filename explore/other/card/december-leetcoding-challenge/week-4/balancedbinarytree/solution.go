package balancedbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var helper func(node *TreeNode) int
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := helper(node.Left)
		r := helper(node.Right)
		if l == -1 || r == -1 || abs(l-r) > 1 {
			return -1
		}
		return max(l, r) + 1
	}

	return helper(root) != -1
}
