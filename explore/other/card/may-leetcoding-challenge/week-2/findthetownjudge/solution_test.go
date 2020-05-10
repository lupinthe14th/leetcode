package findthetownjudge

import (
	"fmt"
	"testing"
)

func TestFindJudge(t *testing.T) {
	type in struct {
		N     int
		trust [][]int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{N: 2, trust: [][]int{{1, 2}}}, want: 2},
		{in: in{N: 3, trust: [][]int{{1, 3}, {2, 3}}}, want: 3},
		{in: in{N: 3, trust: [][]int{{1, 3}, {2, 3}, {3, 1}}}, want: -1},
		{in: in{N: 3, trust: [][]int{{1, 2}, {2, 3}}}, want: -1},
		{in: in{N: 4, trust: [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}}, want: 3},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findJudge(tt.in.N, tt.in.trust)
			if got != tt.want {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
