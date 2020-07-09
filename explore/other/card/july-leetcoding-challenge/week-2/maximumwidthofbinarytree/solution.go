package maximumwidthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/maximum-width-of-binary-tree/discuss/106654/JavaC++-Very-simple-dfs-solution
func widthOfBinaryTree(root *TreeNode) int {
	start := []int{}
	end := []int{}

	var dfs func(node *TreeNode, d, o int) int
	dfs = func(node *TreeNode, d, o int) int {
		if node == nil {
			return 0
		}
		if len(start) == d {
			start = append(start, o)
			end = append(end, o)
		} else {
			end[d] = o
		}
		cur := end[d] - start[d] + 1
		left := dfs(node.Left, d+1, 2*o)
		right := dfs(node.Right, d+1, 2*o+1)
		return max(cur, max(left, right))
	}
	return dfs(root, 0, 1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
