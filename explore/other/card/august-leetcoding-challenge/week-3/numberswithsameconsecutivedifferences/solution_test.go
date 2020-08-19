package numberswithsameconsecutivedifferences

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNumsSameConsecDiff(t *testing.T) {
	t.Parallel()
	type in struct {
		n, k int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{n: 3, k: 7}, want: []int{181, 292, 707, 818, 929}},
		{in: in{n: 2, k: 1}, want: []int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98}},
		{in: in{n: 1, k: 0}, want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{in: in{n: 2, k: 0}, want: []int{11, 22, 33, 44, 55, 66, 77, 88, 99}},
	}

	opts := []cmp.Option{
		cmpopts.SortSlices(func(x, y int) bool { return x < y }),
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := numsSameConsecDiff(tt.in.n, tt.in.k)
			if !cmp.Equal(got, tt.want, opts...,
			) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
