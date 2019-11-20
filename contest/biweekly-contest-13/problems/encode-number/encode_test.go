package encodenumber

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input int
	want  string
}{
	{id: 1, input: 23, want: "1000"},
	{id: 2, input: 107, want: "101100"},
}

func TestEncode(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := encode(tt.input)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}
