package pathsum

import (
	"fmt"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	type in struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{root: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   11,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 2},
				},
			},
			Right: &TreeNode{
				Val:  8,
				Left: &TreeNode{Val: 13},
				Right: &TreeNode{
					Val:   4,
					Right: &TreeNode{Val: 4},
				},
			},
		}, sum: 22}, want: true},
		{in: in{root: nil, sum: 0}, want: false},
		{in: in{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, sum: 1}, want: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := hasPathSum(tt.in.root, tt.in.sum)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
