package queriesonapermutationwithkey

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProcessQueries(t *testing.T) {
	type in struct {
		q []int
		m int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{q: []int{3, 1, 2, 1}, m: 5}, want: []int{2, 1, 2, 1}},
		{in: in{q: []int{4, 1, 2, 2}, m: 4}, want: []int{3, 1, 2, 0}},
		{in: in{q: []int{7, 5, 5, 8, 3}, m: 8}, want: []int{6, 5, 0, 7, 5}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := processQueries(tt.in.q, tt.in.m)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
			}
		})
	}
}
