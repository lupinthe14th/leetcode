package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	l1   *ListNode
	l2   *ListNode
	want *ListNode
}{
	{
		l1:   &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
		l2:   &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
		want: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}},
	},
}

func TestAddTwoNumbers(t *testing.T) {
	for _, tt := range cases {
		l1j, _ := json.Marshal(tt.l1)
		l2j, _ := json.Marshal(tt.l2)
		t.Run(fmt.Sprintf("l1: %s, l2: %s\n", string(l1j), string(l2j)), func(t *testing.T) {
			actual := addTwoNumbers(tt.l1, tt.l2)
			if !reflect.DeepEqual(actual, tt.want) {
				actualj, _ := json.Marshal(actual)
				wantj, _ := json.Marshal(tt.want)
				t.Errorf("%s, want: %s", string(actualj), string(wantj))
			}
		})
	}
}
