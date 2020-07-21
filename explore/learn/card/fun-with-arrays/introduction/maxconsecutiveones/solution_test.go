package maxconsecutiveones

import (
	"fmt"
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{1, 1, 0, 1, 1, 1}, want: 3},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := findMaxConsecutiveOnes(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
