package cousinsinbinarytree

import (
	"fmt"
	"testing"
)

func TestIsCousins(t *testing.T) {
	type in struct {
		root *TreeNode
		x, y int
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}, x: 2, y: 3}, want: false},
		{in: in{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}, x: 5, y: 4}, want: true},
		{in: in{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}, x: 2, y: 3}, want: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isCousins(tt.in.root, tt.in.x, tt.in.y)
			if got != tt.want {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
