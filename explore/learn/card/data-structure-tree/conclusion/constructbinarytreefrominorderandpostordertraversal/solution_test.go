package constructbinarytreefrominorderandpostordertraversal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	type in struct {
		inorder, postorder []int
	}

	tests := []struct {
		in   in
		want *TreeNode
	}{
		{in: in{inorder: []int{9, 3, 15, 20, 7}, postorder: []int{9, 15, 7, 20, 3}},
			want: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7}}}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := buildTree(tt.in.inorder, tt.in.postorder)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
