package uncrossedlines

import (
	"fmt"
	"testing"
)

func TestMaxUncrossedLines(t *testing.T) {
	type in struct {
		A, B []int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{A: []int{1, 4, 2}, B: []int{1, 2, 4}}, want: 2},
		{in: in{A: []int{2, 5, 1, 2, 5}, B: []int{10, 5, 2, 1, 5, 2}}, want: 3},
		{in: in{A: []int{1, 3, 7, 1, 7, 5}, B: []int{1, 9, 2, 5, 1}}, want: 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maxUncrossedLines(tt.in.A, tt.in.B)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
