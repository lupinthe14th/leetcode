package binarytreezigzaglevelordertraversal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
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
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{in: nil, want: [][]int{}},
		{in: &TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 5},
			},
		},
			want: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := zigzagLevelOrder(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
