package howmanynumbersaresmallerthanthecurrentnumber

import (
	"fmt"
	"reflect"
	"testing"
)

type Case struct {
	in   []int
	want []int
}

var cases = []Case{
	{in: []int{8, 1, 2, 2, 3}, want: []int{4, 0, 1, 1, 3}},
	{in: []int{6, 5, 4, 8}, want: []int{2, 1, 0, 3}},
	{in: []int{7, 7, 7, 7}, want: []int{0, 0, 0, 0}},
}

func TestSmallerNumbersThanCurrent(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := smallerNumbersThanCurrent(c.in)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("in: %v, got: %v, want: %v", c.in, got, c.want)
			}
		})
	}
}
