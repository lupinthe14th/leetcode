package conntandsay

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input int
	want  string
}{
	{id: 1, input: 1, want: "1"},
	{id: 2, input: 2, want: "11"},
	{id: 3, input: 3, want: "21"},
	{id: 4, input: 4, want: "1211"},
}

func TestCountAndSay(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := countAndSay(tt.input)
			t.Log(got)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}
