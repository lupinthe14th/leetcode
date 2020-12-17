package validatebinarysearchtree

import (
	"fmt"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   *TreeNode
		want bool
	}{
		{in: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		}, want: true},
		{in: &TreeNode{
			Val:  5,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{Val: 4,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 6},
			},
		}, want: false},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := isValidBST(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
