package rangesumofbst

import (
	"fmt"
	"testing"
)

func TestRangeSumBST(t *testing.T) {
	t.Parallel()
	type in struct {
		root      *TreeNode
		low, high int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 7},
				},
				Right: &TreeNode{
					Val:   15,
					Right: &TreeNode{Val: 18},
				},
			}, low: 7, high: 15}, want: 32},
		{in: in{
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{Val: 3,
						Left: &TreeNode{Val: 1},
					},
					Right: &TreeNode{Val: 7,
						Left: &TreeNode{Val: 6},
					},
				},
				Right: &TreeNode{
					Val:   15,
					Left:  &TreeNode{Val: 13},
					Right: &TreeNode{Val: 18},
				},
			}, low: 6, high: 10}, want: 23},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := rangeSumBST(tt.in.root, tt.in.low, tt.in.high)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
