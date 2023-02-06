package main

import "fmt"

func main() { // присвоение переменной вызываемой функции,вывод в консоль всех функций
	// var totalOrder, order string = addToCard("Стринги", 1, 100)
	// countSum(totalOrder, order)
	// fmt.Println(test())
	// loop()
	// poop()
	// defer fmt.Println("Алена лох")
	// checkEyesColor("yellow")
	// checkEyesColor("blue")
	// checkEyesColor("green")
	// checkEyesColor("brown")
	biba()
}

// func addToCard(name string, count uint16, price uint16) (string, string) { // возвращаемая функция с именованными аргументами и использованием их в конкатенации с выводом в консоль
// 	var order string = fmt.Sprintf("Вы заказали %s, %d шт по %d руб.", name, count, price)
// 	result := fmt.Sprintf("Результат: %s", order)
// 	return result, order
// }

// func countSum(hyi string, pizda string) { // невозвращаемая функция с именованными аргументами и выводом в консоль
// 	fmt.Println(hyi, pizda)
// }

// func test() string { // возвращаемая функция без аргументов и переменных
// 	return "empty"
// }

// func loop() { // Циклы и ее использование в конкатенации с выводом в консоль
// 	for i := 1; i <= 10; i++ {
// 		// zabud := fmt.Sprintf("%d Алена забудь", i)
// 		// fmt.Println(zabud)

// 		// fmt.Println(fmt.Sprintf("%d Алена забудь", i))

// 		fmt.Printf("%d Алена забудь\n", i)
// 	}

// }

// func poop() {
// 	// defer fmt.Println("Алена лох")
// 	for p := 1; p <= 10; p++ {
// 		if p > 4 && p <= 5 {
// 			fmt.Println("Алена пупа")
// 		} else if p > 2 && p < 4 {
// 			fmt.Println("Алена даже не полупупа")
// 		} else if p < 4 {
// 			fmt.Println("Алена полупупа но еще не лупа")
// 		} else {
// 			fmt.Println("Super лупа")
// 		}
// 	}
// }

//	func checkEyesColor(color string) {
//		defer fmt.Println("Алена лох")
//		switch color {
//		case "blue":
//			fmt.Println("Голубая")
//			// return
//		case "brown":
//			fmt.Println("Какашковая")
//			// return
//		case "green":
//			fmt.Println("Топовая")
//			// return
//		default:
//			fmt.Println("Чмошники")
//			// return
//		}
//	}
func biba() {
	a := "we will never be slavyane"
	p := &a
	fmt.Println(a)
	fmt.Println(*p)
	*p = "we will never be slaves"
	fmt.Println(*p)
	fmt.Println(a)
}

// test1
