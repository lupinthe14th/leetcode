package middleofthelinkedlist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		in, want []int
	}{
		{in: []int{1, 2, 3, 4, 5}, want: []int{3, 4, 5}},
		{in: []int{1, 2, 3, 4, 5, 6}, want: []int{4, 5, 6}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := middleNode(newListNode(tt.in))
			want := newListNode(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got: %#+v, want: %#+v", got, want)
			}
		})
	}
}

func newListNode(list []int) *ListNode {
	ln := &ListNode{}
	out := ln
	for i, v := range list {
		ln.Val = v
		if i < len(list)-1 {
			ln.Next = &ListNode{}
			ln = ln.Next
		}
	}
	return out
}
