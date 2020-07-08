package threesum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		in   []int
		want [][]int
	}{
		{in: []int{-1, 0, 1, 2, -1, -4}, want: [][]int{{-1, 0, 1}, {-1, -1, 2}}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := threeSum(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
