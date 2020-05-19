package onlinestockspan

type StockSpanner struct {
	p, w []int
}

func Constructor() StockSpanner {
	return StockSpanner{p: make([]int, 0), w: make([]int, 0)}
}

func (this *StockSpanner) Next(price int) int {
	w := 1
	for len(this.p) != 0 && this.p[len(this.p)-1] <= price {
		this.p = this.p[:len(this.p)-1]
		w += this.w[len(this.w)-1]
		this.w = this.w[:len(this.w)-1]
	}

	this.p = append(this.p, price)
	this.w = append(this.w, w)
	return w
}
