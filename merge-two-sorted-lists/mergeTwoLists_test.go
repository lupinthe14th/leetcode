package mergetwosortedlists

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	id      int
	inputL1 *ListNode
	inputL2 *ListNode
	want    *ListNode
}{
	{id: 1,
		inputL1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}},
		inputL2: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
		want:    &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: nil}}}}}}},
	{id: 2,
		inputL1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}},
		inputL2: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
		want:    &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}}}},
	{id: 3, inputL1: nil, inputL2: &ListNode{Val: 0}, want: &ListNode{Val: 0}},
}

func TestMergeTwoLists(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := mergeTwoLists(tt.inputL1, tt.inputL2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%#v, want: %#v", got, tt.want)
			}
		})
	}
}
