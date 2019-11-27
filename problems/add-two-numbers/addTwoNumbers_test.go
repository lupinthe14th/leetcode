package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	id   int
	l1   *ListNode
	l2   *ListNode
	want *ListNode
}{
	{
		id:   1,
		l1:   &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
		l2:   &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
		want: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}},
	},
	{
		id:   2,
		l1:   &ListNode{Val: 5, Next: nil},
		l2:   &ListNode{Val: 5, Next: nil},
		want: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}},
	},
	{
		id:   3,
		l1:   &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: nil}},
		l2:   &ListNode{Val: 0, Next: nil},
		want: &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: nil}},
	},
	{
		id:   4,
		l1:   &ListNode{Val: 9, Next: &ListNode{Val: 8, Next: nil}},
		l2:   &ListNode{Val: 1, Next: nil},
		want: &ListNode{Val: 0, Next: &ListNode{Val: 9, Next: nil}},
	},
	{
		id:   5,
		l1:   &ListNode{Val: 1, Next: nil},
		l2:   &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}},
		want: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}},
	},
	{
		id:   6,
		l1:   &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: &ListNode{Val: 6, Next: nil}}},
		l2:   &ListNode{Val: 0, Next: nil},
		want: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: &ListNode{Val: 6, Next: nil}}},
	},
	{
		id:   7,
		l1:   &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}},
		l2:   &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
		want: &ListNode{Val: 6, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: nil}}}},
	},
}

func TestAddTwoNumbers(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := addTwoNumbers(tt.l1, tt.l2)
			if !reflect.DeepEqual(got, tt.want) {
				gotj, _ := json.Marshal(got)
				wantj, _ := json.Marshal(tt.want)
				t.Errorf("%s, want: %s", string(gotj), string(wantj))
			}
		})
	}
}
