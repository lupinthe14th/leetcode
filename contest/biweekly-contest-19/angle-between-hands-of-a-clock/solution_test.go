package anglebetweenhandsofaclock

import (
	"fmt"
	"testing"
)

type Case struct {
	input input
	want  float64
}

type input struct {
	hour, minutes int
}

var cases = []Case{
	{input: input{hour: 12, minutes: 30}, want: 165},
	{input: input{hour: 3, minutes: 30}, want: 75},
	{input: input{hour: 3, minutes: 15}, want: 7.5},
	{input: input{hour: 4, minutes: 50}, want: 155},
	{input: input{hour: 12, minutes: 0}, want: 0},
	{input: input{hour: 1, minutes: 57}, want: 76.5},
}

func TestAngleClock(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := angleClock(tt.input.hour, tt.input.minutes)
			if got != tt.want {
				t.Errorf("%f, want: %f", got, tt.want)
			}
		})
	}
}
