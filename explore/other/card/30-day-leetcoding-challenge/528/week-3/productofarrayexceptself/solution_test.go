package productofarrayexceptself

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		in, want []int
	}{
		{in: []int{1, 2, 3, 4}, want: []int{24, 12, 8, 6}},
		{in: []int{0, 0}, want: []int{0, 0}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := productExceptSelf(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
