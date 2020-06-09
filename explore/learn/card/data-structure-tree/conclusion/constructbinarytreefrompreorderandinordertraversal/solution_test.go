package constructbinarytreefrompreorderandinordertraversal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	type in struct {
		preorder, inorder []int
	}
	tests := []struct {
		in   in
		want *TreeNode
	}{
		{in: in{preorder: []int{3, 9, 20, 15, 7}, inorder: []int{9, 3, 15, 20, 7}}, want: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			}}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := buildTree(tt.in.preorder, tt.in.inorder)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
