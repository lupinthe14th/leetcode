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

	var dfs func(node *TreeNode, level, order int) int
	dfs = func(node *TreeNode, level, order int) int {
		if node == nil {
			return 0
		}
		if len(start) == level {
			start = append(start, order)
			end = append(end, order)
		} else {
			end[level] = order
		}
		cur := end[level] - start[level] + 1
		left := dfs(node.Left, level+1, 2*order)
		right := dfs(node.Right, level+1, 2*order+1)
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
