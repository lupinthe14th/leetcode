package hammingdistance

import (
	"fmt"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	type in struct {
		x, y int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{x: 1, y: 4}, want: 2},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := hammingDistance(tt.in.x, tt.in.y)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
