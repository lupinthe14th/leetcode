package validsquare

import (
	"fmt"
	"testing"
)

func TestValidSquare(t *testing.T) {
	t.Parallel()
	type in struct {
		p1, p2, p3, p4 []int
	}
	tests := []struct {
		in   in
		want bool
	}{
		{in: in{p1: []int{0, 0}, p2: []int{1, 1}, p3: []int{1, 0}, p4: []int{0, 1}}, want: true},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := validSquare(tt.in.p1, tt.in.p2, tt.in.p3, tt.in.p4)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
