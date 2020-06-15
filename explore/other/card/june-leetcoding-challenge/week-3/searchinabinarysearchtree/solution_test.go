package searchinabinarysearchtree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchBST(t *testing.T) {
	type in struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		in   in
		want *TreeNode
	}{
		{
			in: in{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{Val: 7},
				},
				val: 2,
			},
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			in: in{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{Val: 7},
				},
				val: 5,
			},
			want: nil,
		},
		{
			in: in{
				root: &TreeNode{
					Val:  18,
					Left: &TreeNode{Val: 2},
					Right: &TreeNode{Val: 22,
						Right: &TreeNode{Val: 63,
							Right: &TreeNode{Val: 84},
						},
					},
				},
				val: 63,
			},
			want: &TreeNode{
				Val:   63,
				Right: &TreeNode{Val: 84},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := searchBST(tt.in.root, tt.in.val)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
