package floodfill

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	type in struct {
		image            [][]int
		sr, sc, newColor int
	}
	tests := []struct {
		in   in
		want [][]int
	}{
		{in: in{image: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, sr: 1, sc: 1, newColor: 2}, want: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
		{in: in{image: [][]int{{0, 0, 0}, {0, 0, 0}}, sr: 0, sc: 0, newColor: 2}, want: [][]int{{2, 2, 2}, {2, 2, 2}}},
		{in: in{image: [][]int{{0, 0, 0}, {0, 1, 1}}, sr: 1, sc: 1, newColor: 1}, want: [][]int{{0, 0, 0}, {0, 1, 1}}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := floodFill(tt.in.image, tt.in.sr, tt.in.sc, tt.in.newColor)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
