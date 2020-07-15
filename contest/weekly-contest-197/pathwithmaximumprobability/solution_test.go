package pathwithmaximumprobability

import (
	"fmt"
	"testing"
)

func TestMaxProbability(t *testing.T) {
	type in struct {
		n, start, end int
		edges         [][]int
		succProb      []float64
	}
	tests := []struct {
		in   in
		want float64
	}{
		{in: in{n: 3, edges: [][]int{{0, 1}, {1, 2}, {0, 2}}, succProb: []float64{0.5, 0.5, 0.2}, start: 0, end: 2}, want: 0.25000},
		{in: in{n: 3, edges: [][]int{{0, 1}, {1, 2}, {0, 2}}, succProb: []float64{0.5, 0.5, 0.3}, start: 0, end: 2}, want: 0.30000},
		{in: in{n: 3, edges: [][]int{{0, 1}}, succProb: []float64{0.5}, start: 0, end: 2}, want: 0.00000},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			//t.Parallel()
			got := maxProbability(tt.in.n, tt.in.edges, tt.in.succProb, tt.in.start, tt.in.end)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
