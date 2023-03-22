package main

import "fmt"

// type Product struct {
// 	Name     string
// 	Color    string
// 	Price    uint16
// 	Quantity uint16
// }

type Person struct {
	Lastname  string
	Firstname string
	Surname   string
	BirhtDate BirhtDate
	Number    int
}

type BirhtDate struct {
	BirthDay   string
	BirthMonth string
	BirthYear  string
}

type Filter struct {
	FilterByName      string
	FilterByBirthDate BirhtDate
}

// TODO: 15.03.2023 Заменить массивы на слайсы СДЕЛАНО 22/03/2023
func main() {
	user1 := Person{
		Lastname:  "Иванов",
		Firstname: "Иван",
		Surname:   "Иванович",
		BirhtDate: BirhtDate{
			BirthDay:   "03",
			BirthMonth: "09",
			BirthYear:  "1945",
		},
		Number: 88005553535,
	}
	user2 := Person{
		Lastname:  "Николаев",
		Firstname: "Николай",
		Surname:   "Николаевич",
		BirhtDate: BirhtDate{
			BirthDay:   "01",
			BirthMonth: "09",
			BirthYear:  "1939",
		},
		Number: 88005553636,
	}
	user3 := Person{
		Lastname:  "Максимов",
		Firstname: "Максим",
		Surname:   "Максимович",
		BirhtDate: BirhtDate{
			BirthDay:   "28",
			BirthMonth: "07",
			BirthYear:  "1914",
		},
		Number: 88005553737,
	}
	user4 := Person{
		Lastname:  "Максимов",
		Firstname: "Николай",
		Surname:   "Максимович",
		BirhtDate: BirhtDate{
			BirthDay:   "18",
			BirthMonth: "11",
			BirthYear:  "1918",
		},
		Number: 88005553838,
	}
	massPerson := []Person{user1, user2, user3, user4}
	// massiveProduct()
	massivePerson(massPerson, Filter{"Николай", BirhtDate{BirthYear: "1918"}})
}

// func massiveProduct() {
// 	var a string
// 	var b string
// 	var prod [4]string
// 	prod[0] = "Name"
// 	prod[1] = "Color"
// 	prod[2] = "Price"
// 	prod[3] = "Quantity"
// 	products := [...]string{"productName", "productColor", "productPrice", "productQuantity"}
// 	a = products[3]
// 	// products[5] = "productCount"
// 	fmt.Println(a)
// 	b = prod[0]
// 	fmt.Println(b)
// }
func (currentPerson Person) getFullName() {
	fmt.Println(currentPerson.Lastname, currentPerson.Firstname, currentPerson.Surname)
}
func massivePerson(currentMassPerson []Person, filter Filter) {
	// TODO: 15.03.2023 Вывести польззователей в мейн и запихнуть массив в аргументы функции СДЕЛАНО 19/03/2023
	// TODO: 15.03.2023 Автоматически высчитывать длинну массива СДЕЛАНО 19/03/2023
	for i := 0; i < len(currentMassPerson); i++ {
		// TODO: 15.03.2023 Передавать фильтры в аргументы СДЕЛАНО 19/03/2023
		if currentMassPerson[i].Firstname == filter.FilterByName && currentMassPerson[i].BirhtDate.BirthYear == filter.FilterByBirthDate.BirthYear {
			// TODO: 15.03.2023 Дополнить стрктуру Person датой рождения, номер телефона и сделать метод который выводит только фио, также сделать структуру фильтров по свойствам структуры Person
			currentMassPerson[i].getFullName()

		}
	}
	// TODO: Добавить проверку на наличие фильтра, и на соответствие фильтру, передавать фильтры вторым аргументом к функции
	// TODO: Вывести последний элемент массива без перебора
}
