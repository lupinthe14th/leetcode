package constructbinarysearchtreefrompreordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	var node *TreeNode
	for _, num := range preorder {
		node = insert(node, num)
	}
	return node
}

func insert(node *TreeNode, num int) *TreeNode {
	if node == nil {
		return &TreeNode{num, nil, nil}
	}

	if num < node.Val {
		node.Left = insert(node.Left, num)
	} else {
		node.Right = insert(node.Right, num)
	}
	return node
}
