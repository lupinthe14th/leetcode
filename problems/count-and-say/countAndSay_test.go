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
	{id: 5, input: 5, want: "111221"},
	{id: 6, input: 6, want: "312211"},
	{id: 7, input: 7, want: "13112221"},
	{id: 8, input: 8, want: "1113213211"},
	{id: 9, input: 9, want: "31131211131221"},
	{id: 10, input: 10, want: "13211311123113112211"},
}

func TestCountAndSay(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := countAndSay(tt.input)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}
