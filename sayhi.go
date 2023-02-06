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
	addToCard(Product{"Футболка", "Черный", 200, 10}, 5)
}

// TODO: Сделать функцию addToCard
func addToCard(currentProduct Product, count uint16) {
	if currentProduct.Quantity < count {
		var result uint16 = count - currentProduct.Quantity
		fmt.Printf("Не хватает %d ед. товара \"%s\" \n", result, currentProduct.Name)
	} else {
		var sum uint16 = currentProduct.Price * count
		fmt.Printf("Вы заказали %d ед. товара '%s' на сумму %d \n", count, currentProduct.Name, sum)
	}
}
