package adddigits

import (
	"fmt"
	"testing"
)

func TestAddDigits(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want int
	}{
		{in: 999, want: 9},
		{in: 18, want: 9},
		{in: 38, want: 2},
		{in: 9, want: 9},
		{in: 0, want: 0},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := addDigits(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
