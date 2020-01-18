package main

// TreeNode is binary tree node definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	visit := root

	for visit != nil || len(stack) != 0 {
		for visit != nil {
			stack = append(stack, visit)
			visit = visit.Left
		}
		visit = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, visit.Val)
		visit = visit.Right

	}
	return ans
}

func main() {}
