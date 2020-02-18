package productofthelastknumbers

type ProductOfNumbers struct {
	ProductOfNumber []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		ProductOfNumber: make([]int, 0),
	}
}

func (this *ProductOfNumbers) Add(num int) {
	this.ProductOfNumber = append(this.ProductOfNumber, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	var ans int = 1
	l := len(this.ProductOfNumber)
	for i := 0; i < k; i++ {
		ans *= this.ProductOfNumber[l-1-i]
	}
	return ans
}
