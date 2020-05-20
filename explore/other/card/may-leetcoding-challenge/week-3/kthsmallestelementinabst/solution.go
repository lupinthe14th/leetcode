package kthsmallestelementinabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		l := len(stack)
		root, stack = stack[l-1], stack[:l-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}
