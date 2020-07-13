package sametree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	out := true
	var helper func(p, q *TreeNode)

	helper = func(p, q *TreeNode) {
		if p == nil && q == nil {
			return
		}
		if p == nil || q == nil {
			out = false
			return
		}
		if p.Left != nil && q.Left != nil {
			helper(p.Left, q.Left)
		} else if p.Left != nil || q.Left != nil {
			out = false
			return
		}
		if p.Right != nil && q.Right != nil {
			helper(p.Right, q.Right)
		} else if p.Right != nil || q.Right != nil {
			out = false
			return
		}
		if p.Val != q.Val {
			out = false
			return
		}
	}
	helper(p, q)
	return out
}
