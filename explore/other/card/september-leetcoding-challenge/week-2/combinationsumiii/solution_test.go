package combinationsumiii

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	t.Parallel()
	type in struct {
		k, n int
	}
	tests := []struct {
		in   in
		want [][]int
	}{
		{in: in{k: 3, n: 7}, want: [][]int{{1, 2, 4}}},
		{in: in{k: 3, n: 9}, want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := combinationSum3(tt.in.k, tt.in.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
