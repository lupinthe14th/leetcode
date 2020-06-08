package poweroftwo

import (
	"fmt"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		in   int
		want bool
	}{
		{in: 1, want: true},
		{in: 16, want: true},
		{in: 218, want: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isPowerOfTwo(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}

}
