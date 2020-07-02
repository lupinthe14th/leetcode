package binarytreelevelordertraversalii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	out := make([][]int, 0)
	if root == nil {
		return out
	}
	var helper func(node *TreeNode, level int)
	helper = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(out) == level {
			out = append(out, []int{node.Val})
		} else {
			out[level] = append(out[level], node.Val)
		}
		if node.Left != nil {
			helper(node.Left, level+1)
		}
		if node.Right != nil {
			helper(node.Right, level+1)
		}
	}

	helper(root, 0)
	return reverse(out)
}

func reverse(a [][]int) [][]int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		j := len(a) - 1 - i
		a[i], a[j] = a[j], a[i]
	}
	return a
}
