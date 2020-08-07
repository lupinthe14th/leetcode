package findallduplicatesinanarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want []int
	}{
		{in: []int{4, 3, 2, 7, 8, 2, 3, 1}, want: []int{2, 3}},
		{in: []int{}, want: []int{}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := findDuplicates(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
