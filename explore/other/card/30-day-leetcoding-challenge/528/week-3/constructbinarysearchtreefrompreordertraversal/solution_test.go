package constructbinarysearchtreefrompreordertraversal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBstFromPreorder(t *testing.T) {
	tests := []struct {
		in   []int
		want *TreeNode
	}{
		{in: []int{8, 5, 1, 7, 10, 12}, want: &TreeNode{Val: 8, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 10, Right: &TreeNode{Val: 12}}}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := bstFromPreorder(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
