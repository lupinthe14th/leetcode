package binarytreelevelordertraversalii

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLevelOrderBottom(t *testing.T) {
	tests := []struct {
		in   *TreeNode
		want [][]int
	}{
		{in: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		}, want: [][]int{{15, 7}, {9, 20}, {3}}},
		{in: &TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 3,
				Right: &TreeNode{Val: 5}},
		}, want: [][]int{{4, 5}, {2, 3}, {1}}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := levelOrderBottom(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
