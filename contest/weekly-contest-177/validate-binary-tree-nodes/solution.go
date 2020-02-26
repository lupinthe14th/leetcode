package validatebinarytreenodes

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	temp := n
	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			temp--
		}
		if rightChild[i] != -1 {
			temp--
		}
	}
	return temp == 1
}
