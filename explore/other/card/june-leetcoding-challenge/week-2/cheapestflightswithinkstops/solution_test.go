package cheapestflightswithinkstops

import (
	"fmt"
	"testing"
)

func TestFindCheapestPrice(t *testing.T) {
	type in struct {
		n, src, dst, K int
		flights        [][]int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 3, flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, src: 0, dst: 2, K: 1}, want: 200},
		{in: in{n: 3, flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, src: 0, dst: 2, K: 0}, want: 500},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findCheapestPrice(tt.in.n, tt.in.flights, tt.in.src, tt.in.dst, tt.in.K)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
