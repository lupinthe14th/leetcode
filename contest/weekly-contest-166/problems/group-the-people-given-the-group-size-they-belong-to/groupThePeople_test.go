package groupthepeoplegiventhegroupsizetheybelongto

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	id    int
	input []int
	want  [][]int
}{
	{id: 1, input: []int{3, 3, 3, 3, 3, 1, 3}, want: [][]int{{5}, {0, 1, 2}, {3, 4, 6}}},
	{id: 2, input: []int{2, 1, 3, 3, 3, 2}, want: [][]int{{1}, {0, 5}, {2, 3, 4}}},
}

func TestGroupThePeople(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := groupThePeople(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
