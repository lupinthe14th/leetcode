package mergetwosortedlists

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	t.Parallel()
	type in struct {
		l1, l2 *ListNode
	}
	tests := []struct {
		in   in
		want *ListNode
	}{
		{in: in{
			l1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
			l2: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}}},
		{in: in{l1: nil, l2: nil}, want: nil},
		{in: in{l1: nil, l2: &ListNode{Val: 0}}, want: &ListNode{Val: 0}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := mergeTwoLists(tt.in.l1, tt.in.l2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
