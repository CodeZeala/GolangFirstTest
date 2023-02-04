package main

import "fmt"

type Product struct {
	Name     string
	Color    string
	Price    uint16
	Quantity uint16
}

func main() {
	addToCard(Product{"Трусы", "Белый", 100, 1}, 10)
}

// TODO: Сделать функцию addToCard
func addToCard(currentProduct Product, count uint16) {
	fmt.Println(currentProduct)
	fmt.Println(count)
}

// test
