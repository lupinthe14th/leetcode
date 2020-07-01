package xoroperationinanarray

import (
	"fmt"
	"testing"
)

func TestXorOperation(t *testing.T) {
	type in struct {
		n, start int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 5, start: 0}, want: 8},
		{in: in{n: 4, start: 3}, want: 8},
		{in: in{n: 1, start: 7}, want: 7},
		{in: in{n: 10, start: 5}, want: 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := xorOperation(tt.in.n, tt.in.start)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
