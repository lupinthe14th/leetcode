package main

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	id   int
	l1   []int
	l2   []int
	want []int
}{
	{id: 1, l1: []int{2, 4, 3}, l2: []int{5, 6, 4}, want: []int{7, 0, 8}},
	{id: 2, l1: []int{5}, l2: []int{5}, want: []int{0, 1}},
	{id: 3, l1: []int{1, 8}, l2: []int{0}, want: []int{1, 8}},
	{id: 4, l1: []int{9, 8}, l2: []int{1}, want: []int{0, 9}},
	{id: 5, l1: []int{1}, l2: []int{9, 9}, want: []int{0, 0, 1}},
	{id: 6, l1: []int{9, 1, 6}, l2: []int{0}, want: []int{9, 1, 6}},
	{id: 7, l1: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, l2: []int{5, 6, 4}, want: []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
}

func intArray2ListNode(ia []int) *ListNode {
	ln := &ListNode{}
	out := ln

	for i, v := range ia {
		ln.Val = v
		if i < len(ia)-1 {
			ln.Next = &ListNode{}
			ln = ln.Next
		}
	}
	return out
}

func listNode2IntArray(l *ListNode) []int {
	var ia []int
	for l != nil {
		ia = append(ia, l.Val)
		l = l.Next
	}
	return ia
}

func TestAddTwoNumbers(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := listNode2IntArray(addTwoNumbers(intArray2ListNode(tt.l1), intArray2ListNode(tt.l2)))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
