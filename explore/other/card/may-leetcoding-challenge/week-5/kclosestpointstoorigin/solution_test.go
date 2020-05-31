package kclosestpointstoorigin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKClosest(t *testing.T) {
	type in struct {
		points [][]int
		K      int
	}
	tests := []struct {
		in   in
		want [][]int
	}{
		{in: in{points: [][]int{{1, 3}, {-2, 2}}, K: 1}, want: [][]int{{-2, 2}}},
		{in: in{points: [][]int{{3, 3}, {5, 1}, {-2, 4}}, K: 2}, want: [][]int{{3, 3}, {-2, 4}}},
		{in: in{points: [][]int{{0, 1}, {1, 0}}, K: 2}, want: [][]int{{0, 1}, {1, 0}}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := kClosest(tt.in.points, tt.in.K)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
