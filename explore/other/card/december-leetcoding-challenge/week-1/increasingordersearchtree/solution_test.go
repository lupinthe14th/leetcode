package increasingordersearchtree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIncreasingBST(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want *TreeNode
	}{
		{
			in: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 1},
					},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val: 6,
					Right: &TreeNode{
						Val:   8,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 9},
					},
				},
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 5,
								Right: &TreeNode{
									Val: 6,
									Right: &TreeNode{
										Val: 7,
										Right: &TreeNode{
											Val: 8,
											Right: &TreeNode{
												Val: 9},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			in: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 7},
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 7},
				},
			},
		},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := increasingBST(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
