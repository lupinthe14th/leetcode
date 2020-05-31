package maximumareaofapieceofcakeafterhorizontalandverticalcuts

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	type in struct {
		h, w                         int
		horizontalCuts, verticalCuts []int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{h: 5, w: 4, horizontalCuts: []int{1, 2, 4}, verticalCuts: []int{1, 3}}, want: 4},
		{in: in{h: 5, w: 4, horizontalCuts: []int{3, 1}, verticalCuts: []int{1}}, want: 6},
		{in: in{h: 5, w: 4, horizontalCuts: []int{3}, verticalCuts: []int{3}}, want: 9},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maxArea(tt.in.h, tt.in.w, tt.in.horizontalCuts, tt.in.verticalCuts)
			if got != tt.want {
				t.Fatalf("in: %v got %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
