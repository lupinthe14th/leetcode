package findthelongestsubstringcontainingvowelsinevencounts

import (
	"fmt"
	"testing"
)

type Case struct {
	in   string
	want int
}

var cases = []Case{
	{in: "eleetminicoworoep", want: 13},
	{in: "leetcodeisgreat", want: 5},
	{in: "bcbcbc", want: 6},
}

func TestFindTheLongestSubstring(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findTheLongestSubstring(c.in)
			if got != c.want {
				t.Errorf("%d, want :%d", got, c.want)
			}
		})
	}
}
