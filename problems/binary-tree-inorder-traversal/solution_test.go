package main

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
	{input: []interface{}{1, 2, 3}, want: []int{2, 1, 3}},
}

func TestInorderTraversal(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := inorderTraversal(slice2Binaryheap(tt.input))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}

// SeeAlso: - https://github.com/golang/tour/blob/master/tree/tree.go
//          - https://github.com/ljun20160606/leetcode/blob/master/algorithms/treenode.go
//          - https://en.wikipedia.org/w/index.php?title=Binary_heap&oldid=927176052
func insert(n **TreeNode, nums []interface{}, idx int) {
	if idx >= len(nums) || nums[idx] == nil {
		return
	}
	*n = new(TreeNode)
	(*n).Val = nums[idx].(int)
	insert(&((*n).Left), nums, 2*idx+1)
	insert(&((*n).Right), nums, 2*idx+2)
}

func slice2Binaryheap(nums []interface{}) *TreeNode {
	n := new(*TreeNode)
	insert(n, nums, 0)
	return *n
}
