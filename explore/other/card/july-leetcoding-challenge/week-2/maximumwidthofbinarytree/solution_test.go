package maximumwidthofbinarytree

import (
	"fmt"
	"testing"
)

func TestWidthOfVinaryTree(t *testing.T) {
	tests := []struct {
		in   *TreeNode
		want int
	}{
		{in: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 9},
			},
		},
			want: 4},
		{
			in: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 3},
				},
			},
			want: 2},
		{
			in: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val: 2,
				}},
			want: 2},
		{
			in: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 6}}},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 9,
						Right: &TreeNode{
							Val: 7}}}},
			want: 8},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := widthOfBinaryTree(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
