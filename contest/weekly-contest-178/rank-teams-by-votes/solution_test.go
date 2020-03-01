package rankteamsbyvotes

import (
	"fmt"
	"testing"
)

type Case struct {
	in   []string
	want string
}

var cases = []Case{
	{in: []string{"ABC", "ACB", "ABC", "ACB", "ACB"}, want: "ACB"},
	{in: []string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"}, want: "ZMNAGUEDSJYLBOPHRQICWFXTVK"},
	{in: []string{"M", "M", "M", "M"}, want: "M"},
	{in: []string{"BCA", "CAB", "CBA", "ABC", "ACB", "BAC"}, want: "ABC"},
}

func TestRankTeams(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := rankTeams(c.in)
			if got != c.want {
				t.Errorf("in: %v, got: %s, want: %s", c.in, got, c.want)
			}
		})
	}
}
