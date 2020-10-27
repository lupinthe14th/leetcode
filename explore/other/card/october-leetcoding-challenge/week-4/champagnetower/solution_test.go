package champagnetower

import (
	"fmt"
	"testing"
)

func TestChampagneTower(t *testing.T) {
	t.Parallel()
	type in struct {
		poured, query_row, query_glass int
	}
	tests := []struct {
		in   in
		want float64
	}{
		{in: in{poured: 1, query_row: 1, query_glass: 1}, want: 0.00000},
		{in: in{poured: 2, query_row: 1, query_glass: 1}, want: 0.50000},
		{in: in{poured: 100000009, query_row: 33, query_glass: 17}, want: 1.00000},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := champagneTower(tt.in.poured, tt.in.query_row, tt.in.query_glass)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
