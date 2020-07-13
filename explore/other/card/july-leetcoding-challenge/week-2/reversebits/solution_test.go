package reversebits

import (
	"fmt"
	"testing"
)

func TestReverseBits(t *testing.T) {
	tests := []struct {
		in, want uint32
	}{
		{in: 43261596, want: 964176192},
		{in: 4294967293, want: 3221225471},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := reverseBits(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
