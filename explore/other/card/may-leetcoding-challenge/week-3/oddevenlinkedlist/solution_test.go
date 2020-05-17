package oddevenlinkedlist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	tests := []struct {
		in, want *ListNode
	}{
		{
			in:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}}},
		},
		{
			in:   &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 7}}}}}}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 4}}}}}}},
		},
		{
			in:   &ListNode{},
			want: &ListNode{},
		},
		{
			in:   nil,
			want: nil,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := oddEvenList(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
