package sortintegersbythenumberof1bits

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
	{in: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, want: []int{0, 1, 2, 4, 8, 3, 5, 6, 7}},
	{in: []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}, want: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}},
	{in: []int{10000, 10000}, want: []int{10000, 10000}},
	{in: []int{2, 3, 5, 7, 11, 13, 17, 19}, want: []int{2, 3, 5, 17, 7, 11, 13, 19}},
	{in: []int{10, 100, 1000, 10000}, want: []int{10, 100, 10000, 1000}},
}

func TestSortByBits(t *testing.T) {
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := sortByBits(c.in)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("%v, want: %v", got, c.want)
			}
		})
	}
}
