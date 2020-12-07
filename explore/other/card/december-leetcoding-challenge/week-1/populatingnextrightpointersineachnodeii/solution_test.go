package populatingnextrightpointersineachnodeii

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConnect(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want *Node
	}{
		{
			in: &Node{
				Val: 1,
				Left: &Node{
					Val:   2,
					Left:  &Node{Val: 4},
					Right: &Node{Val: 5},
				},
				Right: &Node{
					Val:   3,
					Right: &Node{Val: 7},
				},
			},
			want: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val: 4,
						Next: &Node{
							Val:  5,
							Next: &Node{Val: 7},
						},
					},
					Right: &Node{
						Val:  5,
						Next: &Node{Val: 7},
					},
					Next: &Node{
						Val:   3,
						Right: &Node{Val: 7},
					},
				},
				Right: &Node{
					Val:   3,
					Right: &Node{Val: 7},
				},
			},
		},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := connect(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
