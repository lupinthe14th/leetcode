package romantointeger

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input string
	want  int
}{
	{input: "III", want: 3},
	{input: "IV", want: 4},
	{input: "IX", want: 9},
	{input: "LVIII", want: 58},
	{input: "MCMXCIV", want: 1994},
}

func TestRomanToInt(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			got := romanToInt(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
