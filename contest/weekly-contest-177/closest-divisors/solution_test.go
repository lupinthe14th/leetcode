package closestdivisors

import (
	"fmt"
	"reflect"
	"testing"
)

type Case struct {
	in   int
	want []int
}

var cases = []Case{
	{in: 8, want: []int{3, 3}},
	{in: 123, want: []int{5, 25}},
	{in: 999, want: []int{25, 40}},
}

func TestClosestDivisors(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := closestDivisors(c.in)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("%v, want: %v", got, c.want)
			}
		})
	}
}
