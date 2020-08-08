package pathsumiii

import (
	"fmt"
	"testing"
)

func TestPathSum(t *testing.T) {
	type in struct {
		root *TreeNode
		sum  int
	}

	tests := []struct {
		in   in
		want int
	}{

		{in: in{root: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: -2},
				},
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 1},
				},
			},
			Right: &TreeNode{Val: -3,
				Right: &TreeNode{Val: 11},
			},
		}, sum: 8}, want: 3},
		{in: in{root: &TreeNode{
			Val: 1,
			Right: &TreeNode{Val: 2,
				Right: &TreeNode{Val: 3,
					Right: &TreeNode{Val: 4,
						Right: &TreeNode{Val: 5}}}},
		}, sum: 3}, want: 2},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := pathSum(tt.in.root, tt.in.sum)
			if got != tt.want {
				t.Fatalf("in %v got: %v  want %v", tt.in, got, tt.want)
			}
		})
	}
}
