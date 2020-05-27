package possiblebipartition

import (
	"fmt"
	"testing"
)

func TestPossibleBipartition(t *testing.T) {
	type in struct {
		n        int
		dislikes [][]int
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{n: 4, dislikes: [][]int{{1, 2}, {1, 3}, {2, 4}}}, want: true},
		{in: in{n: 3, dislikes: [][]int{{1, 2}, {1, 3}, {2, 3}}}, want: false},
		{in: in{n: 5, dislikes: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}}, want: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := possibleBipartition(tt.in.n, tt.in.dislikes)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
