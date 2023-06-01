package main

import "fmt"

type Cart struct {
	prices []float32
}

func (p *Cart) add(price float32) {
	p.prices = append(p.prices, price)
}

func (x Cart) show() {
	var sum float32 = 0.0

	for _, price := range x.prices {
		sum += price
	}

	fmt.Println(sum)
}

func main() {
	c := Cart{}
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var price float32
		fmt.Scanln(&price)
		c.add(price)
	}

	c.show()
}
