package constructbinarytreefrompreorderandinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/explore/learn/card/data-structure-tree/133/conclusion/943/discuss/34538/My-Accepted-Java-Solution
func buildTree(preorder []int, inorder []int) *TreeNode {
	lpre := len(preorder)
	lin := len(inorder)
	var helper func(pbegin, ibegin, iend int) *TreeNode

	helper = func(pbegin, ibegin, iend int) *TreeNode {
		if pbegin > lpre-1 || ibegin > iend {
			return nil
		}
		root := &TreeNode{}
		root.Val = preorder[pbegin]
		iIdx := 0
		for i := ibegin; i <= iend; i++ {
			if inorder[i] == root.Val {
				iIdx = i
			}
		}

		root.Left = helper(pbegin+1, ibegin, iIdx-1)
		root.Right = helper(pbegin+iIdx-ibegin+1, iIdx+1, iend)
		return root
	}

	if lpre == 0 || lin == 0 || lpre != lin {
		return nil
	}
	return helper(0, 0, lin-1)
}
