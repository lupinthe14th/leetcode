package flippinganimage

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFlipAndInvertImage(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want [][]int
	}{
		{in: [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, want: [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{in: [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}, want: [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := flipAndInvertImage(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
