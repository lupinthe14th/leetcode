package laststoneweight

import (
	"fmt"
	"testing"
)

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{2, 7, 4, 1, 8, 1}, want: 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := lastStoneWeight(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
