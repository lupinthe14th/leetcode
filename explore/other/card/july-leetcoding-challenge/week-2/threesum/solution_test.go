package threesum

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		in   []int
		want [][]int
	}{
		{in: []int{-1, 0, 1, 2, -1, -4}, want: [][]int{{-1, 0, 1}, {-1, -1, 2}}},
	}
	// This Transformer sorts a [][]string.
	trans := cmp.Transformer("Sort", func(in [][]int) [][]int {
		out := append([][]int(nil), in...) // Copy input to avoid mutating it
		sort.Slice(out, func(i, j int) bool {
			for _, v := range out {
				sort.Ints(v)
			}
			if out[i][0] == out[j][0] {
				return out[i][1] < out[j][1]
			}
			return out[i][0] < out[j][0]
		})
		return out
	})
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := threeSum(tt.in)
			if !cmp.Equal(got, tt.want, trans) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
