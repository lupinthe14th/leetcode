package sumofleftleaves

import (
	"fmt"
	"testing"
)

func TestSumLeftLeaves(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   *TreeNode
		want int
	}{
		{in: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		}, want: 24},
		{in: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{
				Val: 3,
			},
		}, want: 4},
		{in: (*TreeNode)(nil), want: 0},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := sumOfLeftLeaves(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
