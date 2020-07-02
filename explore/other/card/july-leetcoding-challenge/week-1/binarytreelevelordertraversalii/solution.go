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

		if len(out) < level+1 {
			out = append([][]int{{}}, out...)
		}
		if node.Left != nil {
			helper(node.Left, level+1)
		}
		if node.Right != nil {
			helper(node.Right, level+1)
		}
		out[len(out)-level-1] = append(out[len(out)-level-1], node.Val)
	}

	helper(root, 0)
	return out
}
