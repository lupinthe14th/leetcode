package checkifastringisavalidsequencefromroottoleavespathinabinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidSequence(root *TreeNode, arr []int) bool {
	return dfs(root, arr, 0)
}

func dfs(root *TreeNode, arr []int, i int) bool {
	if root == nil && i <= len(arr)-1 {
		return false
	}

	if i == len(arr)-1 {
		return root.Val == arr[i] && root.Left == nil && root.Right == nil
	}
	return root.Val == arr[i] && (dfs(root.Left, arr, i+1) || dfs(root.Right, arr, i+1))
}
