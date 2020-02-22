package applydiscounteverynorders

type Cashier struct {
	cur           int
	n             int
	discount      int
	productsPrice map[int]int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {

	productsPrice := make(map[int]int, len(products))
	for i := 0; i < len(products); i++ {
		productsPrice[products[i]] = prices[i]
	}
	return Cashier{cur: 0, n: n, discount: discount, productsPrice: productsPrice}
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	var pay float64
	for i := 0; i < len(product); i++ {
		pay += float64(this.productsPrice[product[i]] * amount[i])
	}

	this.cur++
	// normal
	if this.cur != this.n {
		return pay
	}
	// discount
	this.cur = 0
	return pay - pay*float64(this.discount)/100
}
