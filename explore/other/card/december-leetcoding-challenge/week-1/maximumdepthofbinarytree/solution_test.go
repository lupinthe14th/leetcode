package maximumdepthofbinarytree

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   *TreeNode
		want int
	}{
		{in: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, want: 3},
		{in: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, want: 2},
		{in: nil, want: 0},
		{in: &TreeNode{Val: 0}, want: 1},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := maxDepth(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
