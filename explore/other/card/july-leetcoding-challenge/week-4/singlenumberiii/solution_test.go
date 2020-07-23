package singlenumberiii

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestSingleNumber(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want []int
	}{
		{in: []int{1, 2, 1, 3, 2, 5}, want: []int{3, 5}},
	}

	opts := []cmp.Option{
		cmpopts.SortSlices(func(x, y int) bool { return x < y }),
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := singleNumber(tt.in)
			if !cmp.Equal(got, tt.want, opts...) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
