package coinchange2

import (
	"fmt"
	"testing"
)

func TestChange(t *testing.T) {
	type in struct {
		amount int
		coins  []int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{amount: 5, coins: []int{1, 2, 5}}, want: 4},
		{in: in{amount: 3, coins: []int{2}}, want: 0},
		{in: in{amount: 10, coins: []int{10}}, want: 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := change(tt.in.amount, tt.in.coins)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
