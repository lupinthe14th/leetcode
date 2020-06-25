package countcompletetreenodes

import (
	"fmt"
	"testing"
)

func TestCountNodes(t *testing.T) {
	tests := []struct {
		in   *TreeNode
		want int
	}{
		{in: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 6},
			}}, want: 6},
		{in: &TreeNode{Val: 1}, want: 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := countNodes(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
