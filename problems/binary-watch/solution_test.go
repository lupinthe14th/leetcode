package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Case struct {
	input int
	want  []string
}

var cases = []Case{
	{input: 1, want: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}},
}

func TestReadBinaryWatch(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := readBinaryWatch(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}

func TestBitCount(t *testing.T) {
	var cases = []struct {
		input, want int
	}{
		{input: 1, want: 1},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := bitCount(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
