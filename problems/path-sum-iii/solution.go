package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// refarence: https://leetcode.com/problems/path-sum-iii/discuss/91960/Golang-recursive-DFS
// ノードの頂点から始めた場合と右と左から始めた場合のカウントを合算したのを返す
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return getSumPaths(root, sum, 0) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

// 再帰関数
// preがあるのはPublicにカウンターを置くと、テストケースを回した場合、前回の結果を引きずってしまうので、それぞれ頂点から始めた場合の初期値として入れる
func getSumPaths(root *TreeNode, sum, pre int) int {
	if root == nil {
		return 0
	}
	cur := pre + root.Val

	var c int

	if cur == sum {
		c = 1
	}
	return c + getSumPaths(root.Left, sum, cur) + getSumPaths(root.Right, sum, cur)
}

func main() {}
