package removeduplicatesfromsortedarray

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	id         int
	input      []int
	wantList   []int
	wantLength int
}{
	{id: 1, input: []int{1, 1, 2}, wantList: []int{1, 2}, wantLength: 2},
	{id: 2, input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, wantList: []int{0, 1, 2, 3, 4}, wantLength: 5},
}

func TestRemoveDupulicates(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := removeDuplicates(tt.input)
			nums := tt.input
			if !reflect.DeepEqual(nums[:got], tt.wantList) {
				t.Errorf("%v, wantList: %v", nums, tt.wantList)
			}
			if got != tt.wantLength {
				t.Errorf("%d, wantLength: %d", got, tt.wantLength)
			}
		})
	}
}
