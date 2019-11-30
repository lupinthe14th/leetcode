package hexspeak

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input string
	want  string
}{
	{id: 1, input: "257", want: "IOI"},
	{id: 2, input: "3", want: "ERROR"},
}

func TestToHexspeak(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := toHexspeak(tt.input)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}
