package movezeros

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMoveZeros(t *testing.T) {
	tests := []struct {
		in, want []int
	}{
		{in: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
		{in: []int{0, 0, 1}, want: []int{1, 0, 0}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			moveZeroes(tt.in)
			if diff := cmp.Diff(tt.in, tt.want); diff != "" {
				t.Fatalf("moveZeros() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
