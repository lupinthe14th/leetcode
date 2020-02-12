package tweetcountsperfrequency

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTweetSample(t *testing.T) {
	tweetCounts := Constructor()
	tweetCounts.RecordTweet("tweet3", 0)
	tweetCounts.RecordTweet("tweet3", 60)
	tweetCounts.RecordTweet("tweet3", 10)
	tweetCounts.RecordTweet("tweet3", 120)
	tweetCounts.RecordTweet("tweet3", 60*60*10)

	type input struct {
		freq               string
		tweetName          string
		startTime, endTime int
	}

	type Case struct {
		input input
		want  []int
	}
	var cases = []Case{
		{input: input{freq: "minute", tweetName: "tweet3", startTime: 0, endTime: 59}, want: []int{2}},
		{input: input{freq: "minute", tweetName: "tweet3", startTime: 0, endTime: 60}, want: []int{2, 1}},
		{input: input{freq: "hour", tweetName: "tweet3", startTime: 0, endTime: 210}, want: []int{4}},
		{input: input{freq: "day", tweetName: "tweet3", startTime: 0, endTime: 60 * 60 * 21}, want: []int{5}},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := tweetCounts.GetTweetCountsPerFrequency(tt.input.freq, tt.input.tweetName, tt.input.startTime, tt.input.endTime)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v, want: %v", got, tt.want)
			}
		})
	}
}
