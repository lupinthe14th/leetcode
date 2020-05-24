package intervallistintersections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntervalIntersection(t *testing.T) {
	type in struct {
		A, B [][]int
	}
	tests := []struct {
		in   in
		want [][]int
	}{
		{in: in{A: [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, B: [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}}, want: [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := intervalIntersection(tt.in.A, tt.in.B)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
