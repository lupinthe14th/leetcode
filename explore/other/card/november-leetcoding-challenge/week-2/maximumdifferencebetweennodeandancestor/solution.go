package maximumdifferencebetweennodeandancestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	var traverse func(node *TreeNode, lo, hi int) int
	traverse = func(node *TreeNode, lo, hi int) int {
		if node == nil {
			return hi - lo
		}
		lo = min(lo, node.Val)
		hi = max(hi, node.Val)
		left := traverse(node.Left, lo, hi)
		right := traverse(node.Right, lo, hi)
		return max(left, right)
	}

	return traverse(root, root.Val, root.Val)
}
