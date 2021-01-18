package mergesortedarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	t.Parallel()
	type in struct {
		nums1, nums2 []int
		m, n         int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3}, want: []int{1, 2, 2, 3, 5, 6}},
		{in: in{nums1: []int{1}, m: 1, nums2: []int{}, n: 0}, want: []int{1}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			merge(tt.in.nums1, tt.in.m, tt.in.nums2, tt.in.n)
			if !reflect.DeepEqual(tt.in.nums1, tt.want) {
				t.Fatalf("in: %v want: %v", tt.in, tt.want)
			}
		})
	}
}
