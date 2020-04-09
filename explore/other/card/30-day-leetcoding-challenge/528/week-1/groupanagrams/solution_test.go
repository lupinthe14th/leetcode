package groupanagrams

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func lessStrings(x, y []string) bool {
	sort.Strings(x)
	sort.Strings(y)
	return reflect.DeepEqual(x, y)
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		in   []string
		want [][]string
	}{
		{in: []string{"eat", "tea", "tan", "ate", "nat", "bat"}, want: [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {

			got := groupAnagrams(tt.in)
			if !cmp.Equal(got, tt.want, cmpopts.SortSlices(lessStrings)) {
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
