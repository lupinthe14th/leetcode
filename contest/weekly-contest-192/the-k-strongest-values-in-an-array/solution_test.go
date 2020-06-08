package thekstrongestvaluesinanarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetStrongest(t *testing.T) {
	type in struct {
		arr []int
		k   int
	}
	tests := []struct {
		in   in
		want []int
	}{
		{in: in{arr: []int{1, 2, 3, 4, 5}, k: 2}, want: []int{5, 1}},
		{in: in{arr: []int{1, 1, 3, 5, 5}, k: 2}, want: []int{5, 5}},
		{in: in{arr: []int{6, 7, 11, 7, 6, 8}, k: 5}, want: []int{11, 8, 6, 6, 7}},
		{in: in{arr: []int{6, -3, 7, 2, 11}, k: 3}, want: []int{-3, 11, 2}},
		{in: in{arr: []int{-7, 22, 17, 3}, k: 2}, want: []int{22, 17}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := getStrongest(tt.in.arr, tt.in.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
