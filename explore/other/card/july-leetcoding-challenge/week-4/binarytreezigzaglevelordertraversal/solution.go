package binarytreezigzaglevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	out := [][]int{}
	if root == nil {
		return out
	}

	var traversal func(node *TreeNode, l int)

	traversal = func(node *TreeNode, l int) {
		if node == nil {
			return
		}

		if len(out) <= l {
			out = append(out, []int{})
		}

		if l%2 == 0 {
			out[l] = append(out[l], node.Val)
		} else {
			out[l] = append([]int{node.Val}, out[l]...)

		}
		if node.Left != nil {
			traversal(node.Left, l+1)
		}
		if node.Right != nil {
			traversal(node.Right, l+1)
		}
	}
	traversal(root, 0)
	return out
}
