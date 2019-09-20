package longestcommonprefix

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input []string
	want  string
}{
	{id: 1, input: []string{"flower", "flow", "flight"}, want: "fl"},
	{id: 2, input: []string{"dog", "racecar", "car"}, want: ""},
	{id: 3, input: []string{}, want: ""},
	{id: 4, input: []string{"flower"}, want: ""},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := longestCommonPrefix(tt.input)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}
