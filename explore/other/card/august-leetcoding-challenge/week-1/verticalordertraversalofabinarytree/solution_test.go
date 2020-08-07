package verticalordertraversalofabinarytree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestVerticalTraversal(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   *TreeNode
		want [][]int
	}{
		{in: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		},
			want: [][]int{{9}, {3, 15}, {20}, {7}},
		},
		{in: nil, want: [][]int{}},
		{in: &TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 7},
			},
		},
			want: [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}},
		},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := verticalTraversal(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
