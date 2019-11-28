package intersectionoftwoarrays

import (
	"fmt"
	"reflect"
	"testing"
)

type input struct {
	nums1 []int
	nums2 []int
}

var cases = []struct {
	id    int
	input input
	want  []int
}{
	{id: 1, input: input{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}}, want: []int{2}},
	{id: 2, input: input{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}}, want: []int{4, 9}},
	{id: 3, input: input{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9}}, want: []int{4, 9}},
	{id: 4, input: input{nums1: []int{1, 2}, nums2: []int{1, 1}}, want: []int{1}},
}

func TestIntersection(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := intersection(tt.input.nums1, tt.input.nums2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
