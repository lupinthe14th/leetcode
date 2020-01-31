package binarytreepostordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		visit := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visit.Left != nil {
			stack = append(stack, visit.Left)
		}

		if visit.Right != nil {
			stack = append(stack, visit.Right)
		}
		ans = append([]int{visit.Val}, ans...)
	}
	return ans
}
