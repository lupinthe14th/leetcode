package squaresofasortedarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want []int
	}{
		{in: []int{-4, -1, 0, 3, 10}, want: []int{0, 1, 9, 16, 100}},
		{in: []int{-7, -3, 2, 3, 11}, want: []int{4, 9, 9, 49, 121}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := sortedSquares(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
