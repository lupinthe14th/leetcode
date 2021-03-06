package friendcircles

import (
	"fmt"
	"testing"
)

func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		in   [][]int
		want int
	}{
		{in: [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, want: 2},
		{in: [][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}}, want: 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findCircleNum(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}

}
