package main

import (
	"fmt"
	"log"
	"testing"
)

type Case struct {
	input input
	want  int
}

type input struct {
	root []interface{}
	sum  int
}

var cases = []Case{
	{input: input{root: []interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1}, sum: 8}, want: 3},
	{input: input{root: []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}, sum: 22}, want: 3},
}

func TestPathSum(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := pathSum(slice2Binaryheap(tt.input.root), tt.input.sum)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}

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
