package lab7

import "fmt"

func CompleteLab7() {
	Vegetable := &Vegetable{"Cucumber", 13000, "Хрустящие"}
	car := &Car{"Lada", 10000000, "Красный"}
	products := []Product{Vegetable, car}
	fmt.Println("Общая стоимость", CalculatePrice(products))
	Vegetable.ApplyDiscount(50)
	car.ApplyDiscount(2)
	fmt.Println("Общая стоимость после скидки", CalculatePrice(products))
	fmt.Println("Цвет машины:", car.Color)
	car.ChangeColor("Малиновая")
	fmt.Println("Да, у нас есть для вас ", car.Color,"Лада")
	fmt.Println("Сорт огурчиков:",Vegetable.Sort)
	Vegetable.GetSort("Малиновка")
	fmt.Println("Ваш сорт:",Vegetable.Sort)
	}