package allelementsintwobinarysearchtrees

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetAllElements(t *testing.T) {
	t.Parallel()
	type in struct {
		root1, root2 *TreeNode
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{root1: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4}}, root2: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 3}}}, want: []int{0, 1, 1, 2, 3, 4}},
		{in: in{root1: &TreeNode{Val: 0, Left: &TreeNode{Val: -10}, Right: &TreeNode{Val: 10}}, root2: &TreeNode{Val: 5, Left: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 2},
		}, Right: &TreeNode{Val: 7}}}, want: []int{-10, 0, 0, 1, 2, 5, 7, 10}},
		{in: in{root2: &TreeNode{Val: 5, Left: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 2},
		}, Right: &TreeNode{Val: 7}}}, want: []int{0, 1, 2, 5, 7}},
		{in: in{root1: &TreeNode{Val: 0, Left: &TreeNode{Val: -10}, Right: &TreeNode{Val: 10}}}, want: []int{-10, 0, 10}},
		{in: in{root1: &TreeNode{Val: 1, Right: &TreeNode{Val: 8}}, root2: &TreeNode{Val: 8, Left: &TreeNode{Val: 1}}}, want: []int{1, 1, 8, 8}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := getAllElements(tt.in.root1, tt.in.root2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
