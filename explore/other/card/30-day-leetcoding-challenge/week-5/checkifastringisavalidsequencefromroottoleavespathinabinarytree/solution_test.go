package checkifastringisavalidsequencefromroottoleavespathinabinarytree

import (
	"fmt"
	"testing"
)

func TestIsValidSequence(t *testing.T) {
	type in struct {
		root *TreeNode
		arr  []int
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{root: &TreeNode{Val: 0, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 0}}}, Right: &TreeNode{Val: 0, Left: &TreeNode{Val: 0}}}, arr: []int{0, 1, 0, 1}}, want: true},
		{in: in{root: &TreeNode{Val: 0, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 0}}}, Right: &TreeNode{Val: 0, Left: &TreeNode{Val: 0}}}, arr: []int{0, 0, 1}}, want: false},
		{in: in{root: &TreeNode{Val: 0, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 0}}}, Right: &TreeNode{Val: 0, Left: &TreeNode{Val: 0}}}, arr: []int{0, 1, 1}}, want: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isValidSequence(tt.in.root, tt.in.arr)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
