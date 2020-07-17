package topkfrequentelements

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestTopKFrequent(t *testing.T) {
	t.Parallel()
	type in struct {
		nums []int
		k    int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{nums: []int{1, 1, 1, 2, 2, 3}, k: 2}, want: []int{1, 2}},
		{in: in{nums: []int{1}, k: 1}, want: []int{1}},
		{in: in{nums: []int{1, 2}, k: 2}, want: []int{1, 2}},
		{in: in{nums: []int{1, 1, 1, 2, 2, 333333333}, k: 2}, want: []int{1, 2}},
	}
	opts := []cmp.Option{
		cmpopts.SortSlices(func(x, y int) bool { return x < y }),
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := topKFrequent(tt.in.nums, tt.in.k)
			if !cmp.Equal(got, tt.want, opts...) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
