package numbercomplement

import (
	"fmt"
	"testing"
)

func TestFindComplement(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{in: 5, want: 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findComplement(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
