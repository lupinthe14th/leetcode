package anglebetweenhandsofaclock

import (
	"fmt"
	"testing"
)

func TestAngleClock(t *testing.T) {
	type in struct {
		hour, minutes int
	}
	tests := []struct {
		in   in
		want float64
	}{
		{in: in{hour: 12, minutes: 30}, want: 165},
		{in: in{hour: 3, minutes: 30}, want: 75},
		{in: in{hour: 3, minutes: 15}, want: 7.5},
		{in: in{hour: 4, minutes: 50}, want: 155},
		{in: in{hour: 12, minutes: 0}, want: 0},
		{in: in{hour: 1, minutes: 57}, want: 76.5},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			//			t.Parallel()
			got := angleClock(tt.in.hour, tt.in.minutes)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
