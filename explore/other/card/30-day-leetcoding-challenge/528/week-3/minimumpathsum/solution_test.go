package minimumpathsum

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		in   [][]int
		want int
	}{
		{in: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, want: 7},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := minPathSum(tt.in)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
