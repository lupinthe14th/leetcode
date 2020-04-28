package jumpgame

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		in   []int
		want bool
	}{
		{in: []int{2, 3, 1, 1, 4}, want: true},
		{in: []int{3, 2, 1, 0, 4}, want: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := canJump(tt.in)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
