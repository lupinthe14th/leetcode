package firstbadversion

import (
	"fmt"
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{in: 5, want: 4},
		{in: 50, want: 4},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := firstBadVersion(tt.in)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
