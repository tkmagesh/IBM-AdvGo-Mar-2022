//go:generate echo "Hi there"
package main

import "fmt"

func main() {
	products := Products{
		{10},
		{20},
		{50},
		{30},
		{40},
	}
	costlyProducts := products.Filter(func(p Product) bool {
		return p.cost > 20
	})

	fmt.Println(costlyProducts)
}

type Product struct {
	cost float32
}

type Products []Product

type ProductPredicate func(Product) bool

func (products Products) Filter(predicate ProductPredicate) Products {
	result := Products{}
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}
