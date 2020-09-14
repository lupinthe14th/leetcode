package insertinterval

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	t.Parallel()
	type in struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		in   in
		want [][]int
	}{
		{in: in{intervals: [][]int{{1, 3}, {6, 9}}, newInterval: []int{2, 5}}, want: [][]int{{1, 5}, {6, 9}}},
		{in: in{intervals: [][]int{{1, 5}}, newInterval: []int{2, 3}}, want: [][]int{{1, 5}}},
		{in: in{intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, newInterval: []int{4, 8}}, want: [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{in: in{intervals: [][]int{}, newInterval: []int{5, 7}}, want: [][]int{{5, 7}}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := insert(tt.in.intervals, tt.in.newInterval)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
