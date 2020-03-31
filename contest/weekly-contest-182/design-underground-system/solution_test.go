package designundergroundsystem

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUndergroundSystem(t *testing.T) {
	ug := Constructor()

	type input struct {
		action, stationName, startStation, endStation string
		id, t                                         int
	}

	type Case struct {
		input input
		want  float64
	}
	var cases = []Case{
		{input: input{action: "in", id: 45, stationName: "Leyton", t: 3}},
		{input: input{action: "in", id: 32, stationName: "Paradise", t: 8}},
		{input: input{action: "in", id: 27, stationName: "Leyton", t: 10}},
		{input: input{action: "out", id: 45, stationName: "Waterloo", t: 15}},
		{input: input{action: "out", id: 27, stationName: "Waterloo", t: 20}},
		{input: input{action: "out", id: 32, stationName: "Cambridge", t: 22}},
		{input: input{action: "get", startStation: "Paradise", endStation: "Cambridge"}, want: 14.0},
		{input: input{action: "get", startStation: "Leyton", endStation: "Waterloo"}, want: 11.0},
		{input: input{action: "in", id: 10, stationName: "Leyton", t: 24}},
		{input: input{action: "get", startStation: "Leyton", endStation: "Waterloo"}, want: 11.0},
		{input: input{action: "out", id: 10, stationName: "Waterloo", t: 38}},
		{input: input{action: "get", startStation: "Leyton", endStation: "Waterloo"}, want: 12.0},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			switch tt.input.action {
			case "in":
				ug.CheckIn(tt.input.id, tt.input.stationName, tt.input.t)
			case "out":
				ug.CheckOut(tt.input.id, tt.input.stationName, tt.input.t)
			case "get":
				got := ug.GetAverageTime(tt.input.startStation, tt.input.endStation)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v, want: %v", got, tt.want)
				}
			}
		})
	}
}
