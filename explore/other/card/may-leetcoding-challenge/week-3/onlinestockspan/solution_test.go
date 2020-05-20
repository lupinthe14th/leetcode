package onlinestockspan

import (
	"fmt"
	"testing"
)

func TestOnlineStockSpan(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{in: 100, want: 1},
		{in: 80, want: 1},
		{in: 60, want: 1},
		{in: 70, want: 2},
		{in: 60, want: 1},
		{in: 75, want: 4},
		{in: 85, want: 6},
	}
	s := Constructor()
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := s.Next(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
