package diameterofbinarytree

import (
	"fmt"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		in   *TreeNode
		want int
	}{
		{in: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}, want: 3},
		{in: &TreeNode{Val: 1}, want: 0},
		{in: nil, want: 0},
		{in: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, want: 1},
		{in: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, want: 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := diameterOfBinaryTree(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
