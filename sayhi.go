package main

import "fmt"

type Product struct {
	Name     string
	Color    string
	Price    uint16
	Quantity uint16
}

func main() {
	// addToCard(Product{"Трусы", "Белый", 100, 1}, 10)
	// addToCard(Product{"Футболка", "Черный", 200, 10}, 5)
	calculator(calc{10, 2, "+"})
	calculator(calc{10, 0, "+"})
	calculator(calc{10, 2, "-"})
	calculator(calc{10, 0, "-"})
	calculator(calc{10, 2, "*"})
	calculator(calc{10, 0, "*"})
	calculator(calc{10, 2, "/"})
	calculator(calc{10, 0, "/"})
}

// TODO: Сделать функцию addToCard
// func addToCard(currentProduct Product, count uint16) {
// 	if currentProduct.Quantity < count {
// 		var result uint16 = count - currentProduct.Quantity
// 		fmt.Printf("Не хватает %d ед. товара \"%s\" \n", result, currentProduct.Name)
// 	} else {
// 		var sum uint16 = currentProduct.Price * count
// 		fmt.Printf("Вы заказали %d ед. товара '%s' на сумму %d \n", count, currentProduct.Name, sum)
// 	}
// }

// test1
type calc struct {
	X       uint64
	Y       uint64
	Operand string
}

func calculator(newCalc calc) {
	switch newCalc.Operand {
	// case newCalc.Y == 0:
	// 	fmt.Println("Ошибка вывода")
	case "+":
		fmt.Println(newCalc.X + newCalc.Y)
	case "-":
		fmt.Println(newCalc.X - newCalc.Y)
	case "*":
		{
			fmt.Println(newCalc.X * newCalc.Y)
		}
	case "/":
		if newCalc.Y == 0 {
			fmt.Println("Ошибка вывода")
			return
		}
		fmt.Println(newCalc.X / newCalc.Y)
	}
}
