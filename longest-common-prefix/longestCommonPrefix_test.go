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
	{id: 4, input: []string{"flower"}, want: "flower"},
	{id: 5, input: []string{"c", "c"}, want: "c"},
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

func BenchmarkLongestCommonPrefix(b *testing.B) {
	for _, tt := range cases {
		for i := 0; i < b.N; i++ {
			longestCommonPrefix(tt.input)
		}
	}
}

func TestLongestCommonPrefixHS(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := longestCommonPrefixHS(tt.input)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}

func BenchmarkLongestCommonPrefixHS(b *testing.B) {
	for _, tt := range cases {
		for i := 0; i < b.N; i++ {
			longestCommonPrefixHS(tt.input)
		}
	}
}
