package binarytreelevelordertraversal

import (
	"fmt"
	"reflect"
	"testing"
)

type Case struct {
	input *TreeNode
	want  [][]int
}

var cases = []Case{
	{input: &TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: nil, Right: nil}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15, Left: nil, Right: nil}, Right: &TreeNode{Val: 7, Left: nil, Right: nil}}}, want: [][]int{{3}, {9, 20}, {15, 7}}},
	{input: nil, want: [][]int{}},
}

func TestLevelOrder(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := levelOrder(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
