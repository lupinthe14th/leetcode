package populatingnextrightpointersineachnode

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestConnect(t *testing.T) {
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
					Left:  &Node{Val: 6},
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
							Val: 5,
							Next: &Node{
								Val:  6,
								Next: &Node{Val: 7},
							},
						},
					},
					Right: &Node{
						Val: 5,
						Next: &Node{
							Val:  6,
							Next: &Node{Val: 7},
						},
					},
					Next: &Node{
						Val: 3,
						Left: &Node{
							Val:  6,
							Next: &Node{Val: 7},
						},
						Right: &Node{Val: 7},
					},
				},
				Right: &Node{
					Val: 3,
					Left: &Node{Val: 6,
						Next: &Node{Val: 7},
					},
					Right: &Node{Val: 7},
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := connect(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				log.Print(got.Left)
				log.Print(got.Right)
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
