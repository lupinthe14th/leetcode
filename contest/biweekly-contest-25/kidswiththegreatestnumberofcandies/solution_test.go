package kidswiththegreatestnumberofcandies

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	type in struct {
		candies      []int
		extraCandies int
	}
	tests := []struct {
		in   in
		want []bool
	}{
		{in: in{candies: []int{2, 3, 5, 1, 3}, extraCandies: 3}, want: []bool{true, true, true, false, true}},
		{in: in{candies: []int{4, 2, 1, 1, 2}, extraCandies: 1}, want: []bool{true, false, false, false, false}},
		{in: in{candies: []int{12, 1, 12}, extraCandies: 10}, want: []bool{true, false, true}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := kidsWithCandies(tt.in.candies, tt.in.extraCandies)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
