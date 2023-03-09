package main

import "fmt"

type Product struct {
	Name     string
	Color    string
	Price    uint16
	Quantity uint16
}

type Person struct {
	Name     string
	Surname  string
	Lastname string
}

func main() {
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
	testPerson := Person{
		Name:     "Антон",
		Surname:  "Алексеевич",
		Lastname: "Максимов",
	}
	testPerson.getFullPersonName()
	addToCard(testProduct3, 10)
	addToCard(testProduct2, 5)
	testProduct2.isShirt()
	testProduct3.isShirt()
	massive()
}

func (currentPerson Person) getFullPersonName() {
	fmt.Printf("%s %s %s\n", currentPerson.Lastname, currentPerson.Name, currentPerson.Surname)
}

func addToCard(currentProduct Product, count uint16) {
	if currentProduct.Quantity < count {
		var result uint16 = count - currentProduct.Quantity
		fmt.Printf("Не хватает %d ед. товара \"%s\" \n", result, currentProduct.Name)
	} else {
		var sum uint16 = currentProduct.Price * count
		fmt.Printf("Вы заказали %d ед. товара '%s' на сумму %d р. \n", count, currentProduct.Name, sum)
	}
}

// TODO: 08.03.2023 ОПРЕДЕЛИТЬ И ДОБАВИТЬ ФУНКЦИЮ ИЛИ МЕТОД ДЛЯ ОБНОВЛЕНИЯ КОЛИЧЕСТВА ТОВАРА НА СКЛАДЕ(УПОР НА МЕТОД)
// test1
func (testProduct Product) isShirt() {
	if testProduct.Name == "Футболка" {
		fmt.Printf("Это \"%s\" цвета \"%s\" \n", testProduct.Name, testProduct.Color)
	} else {
		fmt.Printf("Это не \"Футболка\", это -  \"%s\" \n", testProduct.Name)
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
