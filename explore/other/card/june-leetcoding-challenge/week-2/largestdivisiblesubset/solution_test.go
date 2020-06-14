package largestdivisiblesubset

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLargestDivisibleSubset(t *testing.T) {
	tests := []struct {
		in, want []int
	}{
		{in: []int{1, 2, 3}, want: []int{1, 2}},
		{in: []int{1, 2, 4, 8}, want: []int{1, 2, 4, 8}},
		{in: []int{}, want: []int{}},
		{in: []int{546, 669}, want: []int{546}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := largestDivisibleSubset(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
