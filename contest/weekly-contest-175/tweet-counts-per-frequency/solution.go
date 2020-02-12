package tweetcountsperfrequency

type TweetCounts struct {
	Tweet map[string][]int
}

func Constructor() TweetCounts {
	return TweetCounts{
		Tweet: map[string][]int{},
	}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	if _, ok := this.Tweet[tweetName]; !ok {
		this.Tweet[tweetName] = []int{time}
	} else {
		this.Tweet[tweetName] = append(this.Tweet[tweetName], time)

	}
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	interval := GetInterval(freq)

	if _, ok := this.Tweet[tweetName]; !ok {
		return []int{}
	}

	result := make([]int, (endTime-startTime+interval)/interval)
	for _, time := range this.Tweet[tweetName] {
		if time > endTime || time < startTime {
			continue
		}

		thisInterval := (time - startTime) / interval
		result[thisInterval]++
	}
	return result
}

func GetInterval(freq string) int {
	var interval int
	switch freq {
	case "minute":
		interval = 60
	case "hour":
		interval = 60 * 60
	case "day":
		interval = 60 * 60 * 24
	}
	return interval
}
