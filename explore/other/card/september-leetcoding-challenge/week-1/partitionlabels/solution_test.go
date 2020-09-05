package partitionlabels

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   string
		want []int
	}{
		{in: "ababcbacadefegdehijhklij", want: []int{9, 7, 8}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := partitionLabels(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
