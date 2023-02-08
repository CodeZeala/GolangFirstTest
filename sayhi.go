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
	testProduct2 := Product{
		Name:     "Футболка",
		Color:    "Черный",
		Price:    200,
		Quantity: 10,
	}
	testProduct3 := Product{
		Name:     "Трусы",
		Color:    "Белый",
		Price:    100,
		Quantity: 1,
	}
	testProduct2.method()
	testProduct3.method()
	massive()
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

// test1
func (testProduct Product) method() {
	if testProduct.Quantity == 10 {
		fmt.Println("Это Футболка Черный")
	} else {
		fmt.Println("Такого товара не существует")
	}
}
func massive() {
	var a string
	var b string
	var prod [4]string
	prod[0] = "Name"
	prod[1] = "Color"
	prod[2] = "Price"
	prod[3] = "Quantity"
	products := [...]string{"productName", "productColor", "productPrice", "productQuantity"}
	a = products[3]
	// products[5] = "productCount"
	fmt.Println(a)
	b = prod[0]
	fmt.Println(b)
}
