package subsets

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		in   []int
		want [][]int
	}{
		{in: []int{1, 2, 3}, want: [][]int{{3}, {1}, {2}, {1, 2, 3}, {1, 3}, {2, 3}, {1, 2}, {}}},
		{in: []int{9, 0, 3, 5, 7}, want: [][]int{{}, {9}, {0}, {0, 9}, {3}, {3, 9}, {0, 3}, {0, 3, 9}, {5}, {5, 9}, {0, 5}, {0, 5, 9}, {3, 5}, {3, 5, 9}, {0, 3, 5}, {0, 3, 5, 9}, {7}, {7, 9}, {0, 7}, {0, 7, 9}, {3, 7}, {3, 7, 9}, {0, 3, 7}, {0, 3, 7, 9}, {5, 7}, {5, 7, 9}, {0, 5, 7}, {0, 5, 7, 9}, {3, 5, 7}, {3, 5, 7, 9}, {0, 3, 5, 7}, {0, 3, 5, 7, 9}}},
	}

	trans := cmp.Transformer("Sort", func(in [][]int) [][]int {
		out := append([][]int(nil), in...)
		sort.SliceStable(out, func(i, j int) bool {
			for _, v := range out {
				sort.Ints(v)
			}
			if len(out[i]) == 0 && len(out[j]) == 0 {
				return true
			}
			if len(out[i]) == 0 && len(out[j]) != 0 {
				return true
			}
			if len(out[i]) != 0 && len(out[j]) == 0 {
				return false
			}
			if len(out[i]) != len(out[j]) {
				return len(out[i]) < len(out[j])
			}
			r := false
			for k := 0; k < len(out[i]); k++ {
				if out[i][k] != out[j][k] {
					r = out[i][k] < out[j][k]
					break
				}
			}
			return r
		})
		return out
	})

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := subsets(tt.in)
			if !cmp.Equal(got, tt.want, trans) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
