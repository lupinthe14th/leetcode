package kthsmallestelementinabst

import (
	"fmt"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	type in struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}, k: 1}, want: 1},
		{in: in{root: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 1},
				},
				Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 6}}, k: 3}, want: 3},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := kthSmallest(tt.in.root, tt.in.k)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}

}
