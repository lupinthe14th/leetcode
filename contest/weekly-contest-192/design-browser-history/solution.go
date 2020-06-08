package designbrowserhistory

type BrowserHistory struct {
	stack []string
	cur   int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{stack: []string{homepage}, cur: 0}
}

func (this *BrowserHistory) Visit(url string) {
	this.stack = append(this.stack[:this.cur+1], url)
	this.cur++
}

func (this *BrowserHistory) Back(steps int) string {
	this.cur -= steps
	if this.cur < 0 {
		this.cur = 0
	}
	return this.stack[this.cur]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.cur += steps
	if this.cur > len(this.stack)-1 {
		this.cur = len(this.stack) - 1
	}
	return this.stack[this.cur]
}
