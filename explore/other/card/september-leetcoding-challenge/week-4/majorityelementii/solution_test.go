package majorityelementii

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestMajorityElement(t *testing.T) {
	t.Parallel()
	type in struct {
		nums []int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{nums: []int{3, 2, 3}}, want: []int{3}},
		{in: in{nums: []int{1, 1, 1, 3, 3, 2, 2, 2}}, want: []int{1, 2}},
		{in: in{nums: []int{1, 2, 3, 1, 2, 3, 1, 2}}, want: []int{1, 2}},
		{in: in{nums: []int{0, 0, 0}}, want: []int{0}},
	}

	opts := []cmp.Option{
		cmpopts.SortSlices(func(x, y int) bool { return x < y }),
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := majorityElement(tt.in.nums)
			if !cmp.Equal(got, tt.want, opts...) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
