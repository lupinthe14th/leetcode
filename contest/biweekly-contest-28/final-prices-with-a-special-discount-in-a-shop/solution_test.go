package finalpriceswithaspecialdiscountinashop

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFinalPrices(t *testing.T) {
	tests := []struct {
		in, want []int
	}{
		{in: []int{8, 4, 6, 2, 3}, want: []int{4, 2, 4, 2, 3}},
		{in: []int{1, 2, 3, 4, 5}, want: []int{1, 2, 3, 4, 5}},
		{in: []int{10, 1, 1, 6}, want: []int{9, 0, 1, 6}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := finalPrices(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got :%v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
