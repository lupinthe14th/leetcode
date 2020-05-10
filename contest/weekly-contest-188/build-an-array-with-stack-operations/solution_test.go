package buildanarraywithstackoperations

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuildArray(t *testing.T) {
	type in struct {
		t []int
		n int
	}
	tests := []struct {
		in   in
		want []string
	}{
		{in: in{t: []int{1, 3}, n: 3}, want: []string{"Push", "Push", "Pop", "Push"}},
		{in: in{t: []int{1, 2, 3}, n: 3}, want: []string{"Push", "Push", "Push"}},
		{in: in{t: []int{1, 2}, n: 4}, want: []string{"Push", "Push"}},
		{in: in{t: []int{2, 3, 4}, n: 4}, want: []string{"Push", "Pop", "Push", "Push", "Push"}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := buildArray(tt.in.t, tt.in.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
