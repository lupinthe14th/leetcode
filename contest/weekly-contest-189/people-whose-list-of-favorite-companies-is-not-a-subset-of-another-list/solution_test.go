package peoplewhoselistoffavoritecompaniesisnotasubsetofanotherlist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPeopleIndex(t *testing.T) {
	tests := []struct {
		in   [][]string
		want []int
	}{
		{
			in:   [][]string{{"leetcode", "google", "facebook"}, {"google", "microsoft"}, {"google", "facebook"}, {"google"}, {"amazon"}},
			want: []int{0, 1, 4},
		},
		{
			in: [][]string{
				{"leetcode", "google", "facebook"}, {"leetcode", "amazon"}, {"facebook", "google"},
			},
			want: []int{0, 1},
		},
		{
			in: [][]string{
				{"leetcode"}, {"google"}, {"facebook"}, {"amazon"},
			},
			want: []int{0, 1, 2, 3},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := peopleIndexes(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
