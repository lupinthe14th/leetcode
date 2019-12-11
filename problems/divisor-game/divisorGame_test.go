package divisorgame

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input int
	want  bool
}{
	{id: 1, input: 2, want: true},
	{id: 2, input: 3, want: false},
}

func TestDivisorGame(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := divisorGame(tt.input)
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}
}
