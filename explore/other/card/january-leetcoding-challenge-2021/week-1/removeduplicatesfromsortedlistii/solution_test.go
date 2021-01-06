package removeduplicatesfromsortedlistii

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want *ListNode
	}{
		{
			in: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5},
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5},
				},
			}},
		{
			in: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			}}}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := deleteDuplicates(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
