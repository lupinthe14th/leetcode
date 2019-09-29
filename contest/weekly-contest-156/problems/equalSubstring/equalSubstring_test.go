package equalsubstring

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id             int
	inputS, inputT string
	inputM         int
	want           int
}{
	{id: 1, inputS: "abcd", inputT: "bcdf", inputM: 3, want: 3},
	{id: 2, inputS: "abcd", inputT: "cdef", inputM: 3, want: 1},
	{id: 3, inputS: "abcd", inputT: "acde", inputM: 0, want: 1},
	{id: 4, inputS: "krrgw", inputT: "zjxss", inputM: 19, want: 2},
}

func TestUniqueOccurrences(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintln(tt.id), func(t *testing.T) {
			got := equalSubstring(tt.inputS, tt.inputT, tt.inputM)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
