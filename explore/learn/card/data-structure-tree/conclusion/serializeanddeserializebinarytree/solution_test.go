package serializeanddeserializebinarytree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		in, want *TreeNode
	}{
		{
			in: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			},
		},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			codec := Constructor()
			data := codec.serialize(tt.in)
			got := codec.deserialize(data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
