package twocityscheduling

import (
	"fmt"
	"testing"
)

func TestTwoCitySchedCost(t *testing.T) {
	tests := []struct {
		in   [][]int
		want int
	}{
		{in: [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}, want: 110},
		{in: [][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}, want: 1859},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := twoCitySchedCost(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
