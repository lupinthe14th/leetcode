package permutationsii

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   []int
		want [][]int
	}{
		{in: []int{1, 1, 2}, want: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		{in: []int{1, 2, 3}, want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := permuteUnique(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
