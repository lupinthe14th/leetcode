package findrightinterval

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindRightInterval(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   [][]int
		want []int
	}{
		{in: [][]int{{1, 2}}, want: []int{-1}},
		{in: [][]int{{3, 4}, {2, 3}, {1, 2}}, want: []int{-1, 0, 1}},
		{in: [][]int{{1, 4}, {2, 3}, {3, 4}}, want: []int{-1, 2, -1}},
		{in: [][]int{{1, 12}, {2, 9}, {3, 10}, {13, 14}, {15, 16}, {16, 17}}, want: []int{3, 3, 3, 4, 5, -1}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := findRightInterval(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
