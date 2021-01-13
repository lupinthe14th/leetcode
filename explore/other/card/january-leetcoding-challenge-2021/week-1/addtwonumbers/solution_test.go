package addtwonumbers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Parallel()
	type in struct {
		l1, l2 *ListNode
	}
	tests := []struct {
		in   in
		want *ListNode
	}{
		{in: in{
			l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
		}, want: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}},
		{in: in{
			l1: &ListNode{Val: 0},
			l2: &ListNode{Val: 0},
		}, want: &ListNode{Val: 0}},
		{in: in{
			l1: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 9,
									}}}}}}},
			l2: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val:  9,
						Next: &ListNode{Val: 9}}}},
		}, want: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{Val: 9,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val:  0,
									Next: &ListNode{Val: 1},
								}}}}}}}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := addTwoNumbers(tt.in.l1, tt.in.l2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
