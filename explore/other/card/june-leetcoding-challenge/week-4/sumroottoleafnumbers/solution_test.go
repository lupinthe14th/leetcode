package sumroottoleafnumbers

import (
	"fmt"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		in   *TreeNode
		want int
	}{
		{in: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		},
			want: 25},
		{in: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   9,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1},
			},
			Right: &TreeNode{Val: 0},
		},
			want: 1026},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := sumNumbers(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
