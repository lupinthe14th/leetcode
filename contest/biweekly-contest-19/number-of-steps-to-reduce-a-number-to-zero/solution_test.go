package numberofstepstoreduceanumbertozero

import (
	"fmt"
	"testing"
)

type Case struct {
	input, want int
}

var cases = []Case{
	{input: 14, want: 6},
}

func TestNumberOfSteps(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := numberOfSteps(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
