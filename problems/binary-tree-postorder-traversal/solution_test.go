package binarytreepostordertraversal

import (
	"fmt"
	"reflect"
	"testing"
)

type Case struct {
	input []interface{}
	want  []int
}

var cases = []Case{
	{input: []interface{}{1, nil, 2, 3}, want: []int{3, 2, 1}},
}

func TestPostorderTraversal(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := postorderTraversal(slice2TreeNode(tt.input))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}

func slice2TreeNode(sl []interface{}) *TreeNode {
	var n *TreeNode
	for _, s := range sl {
		n = insert(n, s)
	}
	return n
}

func insert(n *TreeNode, data interface{}) *TreeNode {
	if n == nil {
		return &TreeNode{data.(int), nil, nil}
	}

	if data == nil {
		return n
	}

	if data.(int) < n.Val {
		n.Left = insert(n.Left, data)
	} else {
		n.Right = insert(n.Right, data)
	}
	return n
}
