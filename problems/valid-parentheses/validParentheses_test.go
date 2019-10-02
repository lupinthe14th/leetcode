package validparentheses

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input string
	want  bool
}{
	{id: 1, input: "()", want: true},
	{id: 2, input: "()[]{}", want: true},
	{id: 3, input: "(]", want: false},
	{id: 4, input: "([)]", want: false},
	{id: 5, input: "{[]}", want: true},
	{id: 6, input: "", want: true},
	{id: 7, input: "(([]){})", want: true},
}

func TestIsValid(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := isValid(tt.input)
			if got != tt.want {
				t.Errorf("%t, want: %t", got, tt.want)
			}
		})
	}
}
