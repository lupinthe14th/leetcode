package groupanagrams

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		in   []string
		want [][]string
	}{
		{in: []string{"eat", "tea", "tan", "ate", "nat", "bat"}, want: [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}},
	}
	trans := cmp.Transformer("Sort", func(in [][]string) [][]string {
		out := in
		sort.Slice(out, func(i, j int) bool {
			for _, v := range out {
				sort.Strings(v)
			}
			return out[i][0] < out[j][0]
		})
		return out
	})
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {

			got := groupAnagrams(tt.in)
			if !cmp.Equal(got, tt.want, trans) {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestSortString(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{in: "eat", want: "aet"},
		{in: "leetcode", want: "cdeeelot"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := sortString(tt.in)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
