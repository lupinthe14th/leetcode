package tweetcountsperfrequency

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTweetSample(t *testing.T) {
	tweetCounts := Constructor()

	type input struct {
		action, freq, tweetName  string
		time, startTime, endTime int
	}

	type Case struct {
		input input
		want  []int
	}
	var cases = []Case{
		{input: input{action: "get", freq: "minute", tweetName: "tweet3", startTime: 0, endTime: 59}, want: []int{}},
		{input: input{action: "record", tweetName: "tweet3", time: 0}},
		{input: input{action: "record", tweetName: "tweet3", time: 60}},
		{input: input{action: "record", tweetName: "tweet3", time: 10}},
		{input: input{action: "get", freq: "minute", tweetName: "tweet3", startTime: 0, endTime: 59}, want: []int{2}},
		{input: input{action: "get", freq: "minute", tweetName: "tweet3", startTime: 0, endTime: 60}, want: []int{2, 1}},
		{input: input{action: "record", tweetName: "tweet3", time: 120}},
		{input: input{action: "record", tweetName: "tweet3", time: 60 * 60 * 10}},
		{input: input{action: "get", freq: "hour", tweetName: "tweet3", startTime: 0, endTime: 210}, want: []int{4}},
		{input: input{action: "get", freq: "day", tweetName: "tweet3", startTime: 0, endTime: 60 * 60 * 21}, want: []int{5}},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			switch tt.input.action {
			case "record":
				tweetCounts.RecordTweet(tt.input.tweetName, tt.input.time)
			case "get":
				got := tweetCounts.GetTweetCountsPerFrequency(tt.input.freq, tt.input.tweetName, tt.input.startTime, tt.input.endTime)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v, want: %v", got, tt.want)
				}
			}
		})
	}
}
