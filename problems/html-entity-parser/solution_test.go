package htmlentityparser

import (
	"fmt"
	"testing"
)

func TestEntityParser(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{in: "&amp; is an HTML entity but &ambassador; is not.", want: "& is an HTML entity but &ambassador; is not."},
		{in: "and I quote: &quot;...&quot;", want: "and I quote: \"...\""},
		{in: "Stay home! Practice on Leetcode :)", want: "Stay home! Practice on Leetcode :)"},
		{in: "x &gt; y &amp;&amp; x &lt; y is always false", want: "x > y && x < y is always false"},
		{in: "leetcode.com&frasl;problemset&frasl;all", want: "leetcode.com/problemset/all"},
		{in: "&amp;gt;", want: "&gt;"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := entityParser(tt.in)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
