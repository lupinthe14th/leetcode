package maximumdifferencebetweennodeandancestor

import (
	"fmt"
	"testing"
)

func TestMaxAncestorDiff(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   *TreeNode
		want int
	}{
		{
			in: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{Val: 6,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 7},
					},
				},
				Right: &TreeNode{
					Val: 10,
					Right: &TreeNode{
						Val:  14,
						Left: &TreeNode{Val: 13},
					},
				},
			},
			want: 7,
		},
		{
			in: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val:  0,
						Left: &TreeNode{Val: 3},
					},
				},
			},
			want: 3,
		},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := maxAncestorDiff(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
