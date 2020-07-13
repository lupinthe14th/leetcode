package sametree

import (
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	type in struct {
		p, q *TreeNode
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3}},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3}},
		}, want: true},
		{in: in{
			p: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2}},
			q: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2}},
		}, want: false},
		{in: in{
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 1}},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2}},
		}, want: false},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := isSameTree(tt.in.p, tt.in.q)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
