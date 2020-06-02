package constructbinarytreefrominorderandpostordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Seen: https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/discuss/496756/Golang-A-Golang-recursive-solution-with-picture-explanation
func buildTree(inorder []int, postorder []int) *TreeNode {
	var helper func(ibegin, iend, rpos int) *TreeNode

	helper = func(ibegin, iend, rpos int) *TreeNode {
		if rpos < 0 || ibegin > iend {
			return nil
		}
		i := ibegin
		for i < iend && inorder[i] != postorder[rpos] {
			i++
		}

		root := &TreeNode{}
		root.Val = postorder[rpos]
		root.Left = helper(ibegin, i-1, rpos-iend+i-1)
		root.Right = helper(i+1, iend, rpos-1)
		return root
	}
	lin := len(inorder)
	lpo := len(postorder)

	if lin == 0 || lpo == 0 || lin != lpo {
		return nil
	}

	return helper(0, lin-1, lpo-1)
}
