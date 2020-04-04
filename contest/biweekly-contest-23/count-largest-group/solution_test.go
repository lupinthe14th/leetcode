package countlargestgroup

import (
	"fmt"
	"testing"
)

func TestCountLargestGroup(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{in: 13, want: 4},
		{in: 2, want: 2},
		{in: 15, want: 6},
		{in: 24, want: 5},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := countLargestGroup(tt.in)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
