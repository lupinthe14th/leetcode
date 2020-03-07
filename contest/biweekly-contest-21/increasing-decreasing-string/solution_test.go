package increasingdecreasingstring

import (
	"fmt"
	"testing"
)

type Case struct {
	in   string
	want string
}

var cases = []Case{
	{in: "aaaabbbbcccc", want: "abccbaabccba"},
	{in: "rat", want: "art"},
	{in: "leetcode", want: "cdelotee"},
	{in: "ggggggg", want: "ggggggg"},
	{in: "spo", want: "ops"},
}

func TestSortString(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := sortString(c.in)
			if got != c.want {
				t.Errorf("%s, want: %s", got, c.want)
			}
		})
	}
}
