package binarytreemaximumpathsum

import (
	"fmt"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		in   *TreeNode
		want int
	}{
		{in: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, want: 6},
		{in: &TreeNode{Val: -10, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, want: 42},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maxPathSum(tt.in)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
