package constructbinarysearchtreefrompreordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	i := 0
	var helper func(bound int) *TreeNode
	helper = func(bound int) *TreeNode {
		if i == len(preorder) || preorder[i] > bound {
			return nil
		}
		root := &TreeNode{Val: preorder[i]}
		i++
		root.Left = helper(root.Val)
		root.Right = helper(bound)
		return root
	}
	return helper(1 << 31)
}
