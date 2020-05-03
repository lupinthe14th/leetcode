package jewelsandstones

import (
	"fmt"
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	type in struct {
		J, S string
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{J: "aA", S: "aAAbbbb"}, want: 3},
		{in: in{J: "z", S: "ZZ"}, want: 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := numJewelsInStones(tt.in.J, tt.in.S)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
