package searchinabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > val && root.Left != nil {
		return searchBST(root.Left, val)
	}
	if root.Val < val && root.Right != nil {
		return searchBST(root.Right, val)
	}
	if root.Val != val {
		return nil
	}
	return root
}
