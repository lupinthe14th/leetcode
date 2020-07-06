package plusone

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		in, want []int
	}{
		{in: []int{1, 2, 3}, want: []int{1, 2, 4}},
		{in: []int{4, 3, 2, 1}, want: []int{4, 3, 2, 2}},
		{in: []int{4, 3, 2, 9}, want: []int{4, 3, 3, 0}},
		{in: []int{9}, want: []int{1, 0}},
		{in: []int{9, 9}, want: []int{1, 0, 0}},
		{in: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}, want: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := plusOne(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
