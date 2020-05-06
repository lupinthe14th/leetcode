package majorityelement

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{3, 2, 3}, want: 3},
		{in: []int{2, 2, 1, 1, 1, 2, 2}, want: 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := majorityElement(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
