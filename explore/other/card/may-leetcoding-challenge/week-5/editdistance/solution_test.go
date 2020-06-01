package editdistance

import (
	"fmt"
	"testing"
)

func TestMinDistance(t *testing.T) {
	type in struct {
		w1, w2 string
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{w1: "horse", w2: "ros"}, want: 3},
		{in: in{w1: "intention", w2: "execution"}, want: 5},
		{in: in{w1: "", w2: "a"}, want: 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := minDistance(tt.in.w1, tt.in.w2)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
