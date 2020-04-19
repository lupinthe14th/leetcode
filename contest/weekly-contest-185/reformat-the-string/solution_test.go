package reformatthestring

import (
	"fmt"
	"testing"
)

func TestReformat(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{in: "a0b1c2", want: "0a1b2c"},
		{in: "leetcode", want: ""},
		{in: "1229857369", want: ""},
		{in: "covid2019", want: "c2o0v1i9d"},
		{in: "ab123", want: "1a2b3"},
		{in: "j", want: "j"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := reformat(tt.in)
			if got != tt.want {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
