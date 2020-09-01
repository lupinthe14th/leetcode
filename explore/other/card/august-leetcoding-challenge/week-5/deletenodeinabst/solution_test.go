package deletenodeinabst

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	t.Parallel()
	type in struct {
		root *TreeNode
		key  int
	}
	tests := []struct {
		in   in
		want *TreeNode
	}{
		{in: in{root: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   6,
				Right: &TreeNode{Val: 7},
			},
		}, key: 3}, want: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:  4,
				Left: &TreeNode{Val: 2},
			},
			Right: &TreeNode{
				Val:   6,
				Right: &TreeNode{Val: 7},
			},
		}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := deleteNode(tt.in.root, tt.in.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
