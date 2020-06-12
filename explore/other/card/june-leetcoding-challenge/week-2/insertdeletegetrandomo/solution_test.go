package insertdeletegetrandomo

import (
	"fmt"
	"testing"
)

func TestInsertDeleteGetRandomBigO(t *testing.T) {

	type in struct {
		act string
		val int
	}
	tests := []struct {
		in   in
		want interface{}
	}{
		{in: in{act: "ins", val: 0}, want: true},
		{in: in{act: "ins", val: 1}, want: true},
		{in: in{act: "del", val: 0}, want: true},
		{in: in{act: "ins", val: 2}, want: true},
		{in: in{act: "del", val: 1}, want: true},
		{in: in{act: "get"}, want: 2},
	}
	r := Constructor()
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			var got interface{}
			switch tt.in.act {
			case "ins":
				got = r.Insert(tt.in.val)
			case "del":
				got = r.Remove(tt.in.val)
			case "get":
				got = r.GetRandom()
			}
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
