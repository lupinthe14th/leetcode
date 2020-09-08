package sumofroottoleafbinarynumbers

import (
	"fmt"
	"testing"
)

func TestSumRootToLeaf(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   *TreeNode
		want int
	}{
		{in: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: 0},
				Right: &TreeNode{Val: 1},
			},
			Right: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 0},
				Right: &TreeNode{Val: 1},
			},
		}, want: 22},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := sumRootToLeaf(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
