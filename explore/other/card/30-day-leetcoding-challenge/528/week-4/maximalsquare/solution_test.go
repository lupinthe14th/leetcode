package maximalsquare

import (
	"fmt"
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	tests := []struct {
		in   [][]byte
		want int
	}{
		{in: [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}, want: 4},
		{in: [][]byte{{'1', '1', '0', '1'}, {'1', '1', '0', '1'}, {'1', '1', '1', '1'}}, want: 4},
		{in: [][]byte{}, want: 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maximalSquare(tt.in)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
