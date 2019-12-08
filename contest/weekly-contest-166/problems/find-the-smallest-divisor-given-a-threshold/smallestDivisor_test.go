package findthesmallestdivisorgivenathreshold

import (
	"fmt"
	"testing"
)

type input struct {
	nums      []int
	threshold int
}

var cases = []struct {
	id    int
	input input
	want  int
}{
	{id: 1, input: input{nums: []int{1, 2, 5, 9}, threshold: 6}, want: 5},
	{id: 2, input: input{nums: []int{2, 3, 5, 7, 11}, threshold: 11}, want: 3},
	{id: 3, input: input{nums: []int{19}, threshold: 5}, want: 4},
}

func TestSmallestDivisor(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := smallestDivisor(tt.input.nums, tt.input.threshold)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}

func BenchmarkSmallestDivisor(b *testing.B) {
	for _, tt := range cases {
		for i := 0; i < b.N; i++ {
			smallestDivisor(tt.input.nums, tt.input.threshold)
		}
	}
}
