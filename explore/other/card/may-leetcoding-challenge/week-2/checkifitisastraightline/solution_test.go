package checkifitisastraightline

import (
	"fmt"
	"testing"
)

func TestCheckStraightLine(t *testing.T) {
	tests := []struct {
		in   [][]int
		want bool
	}{
		{in: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}, want: true},
		{in: [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}, want: false},
		{in: [][]int{{17, -14}, {-11, -16}, {-17, 4}, {-7, -11}, {-18, -12}, {2, -3}, {-7, -19}, {-1, -2}, {15, -9}, {-18, 16}, {11, 9}, {-6, -4}, {-18, -19}, {1, 19}, {-11, -7}, {11, 17}, {-18, 4}, {14, -17}, {17, -10}}, want: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := checkStraightLine(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
