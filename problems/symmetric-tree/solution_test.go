package main

import (
	"fmt"
	"log"
	"testing"
)

type Case struct {
	input []interface{}
	want  bool
}

var cases = []Case{
	{input: []interface{}{1, 2, 2, 3, 4, 4, 3}, want: true},
	{input: []interface{}{1, 2, 2, nil, 3, nil, 3}, want: false},
}

func TestIsSymmetric(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isSymmetric(slice2Binaryheap(tt.input))
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
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
	printTN(*n)
	return *n
}

func printTN(n *TreeNode) {
	p := n
	stack := make([]*TreeNode, 0)
	stack = append(stack, p)
	for len(stack) != 0 {
		visit := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		log.Printf("Val: %v", visit.Val)
		log.Printf("Left: %v", visit.Left)
		log.Printf("Right: %v", visit.Right)
		if visit.Left != nil {
			stack = append(stack, visit.Left)
		}
		if visit.Right != nil {
			stack = append(stack, visit.Right)
		}
	}
}
