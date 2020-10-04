package combinationsum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	t.Parallel()
	type in struct {
		candidates []int
		target     int
	}
	tests := []struct {
		in   in
		want [][]int
	}{
		{in: in{candidates: []int{2, 3, 6, 7}, target: 7}, want: [][]int{{2, 2, 3}, {7}}},
		{in: in{candidates: []int{2, 3, 5}, target: 8}, want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{in: in{candidates: []int{2}, target: 1}, want: [][]int{}},
		{in: in{candidates: []int{1}, target: 1}, want: [][]int{{1}}},
		{in: in{candidates: []int{1}, target: 2}, want: [][]int{{1, 1}}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := combinationSum(tt.in.candidates, tt.in.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
