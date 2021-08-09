package main

import (
	"decorator/beveragedecorator"
	"fmt"
)

func main() {
	fmt.Println("What kind of coffee do you want?\n1. - Espresso 1.99$\n2. - House Blend 0.89$")
	var beverage beveragedecorator.Beverage
	var choice int
	fmt.Scanf("%d\n", &choice)
	switch choice {
	case 1:
		beverage = beveragedecorator.NewEspresso()
		fmt.Println("Your order:", beverage.GetDescription(), beverage.Cost())
		beverage = ChooseCondiment(beverage)
	case 2:
		beverage = beveragedecorator.NewHouseBlend()
		fmt.Println("Your order:", beverage.GetDescription(), beverage.Cost())
		beverage = ChooseCondiment(beverage)
	}
	fmt.Printf("Your choice: %s, cost: %f\n", beverage.GetDescription(), beverage.Cost())
}

func ChooseCondiment(beverage beveragedecorator.Beverage) beveragedecorator.Beverage {
	var condiment int
	fmt.Println("What kind of condiment do you prefer?\n1. - Milk 0.10$\n2. - Mocha 0.20$\n3. - Exit")
	for condiment != 3 {
		fmt.Scanf("%d\n", &condiment)
		switch condiment {
		case 1:
			beverage = beveragedecorator.NewMilk(beverage)
			fmt.Println("Milk added")
			fmt.Println("Your order:", beverage.GetDescription(), beverage.Cost())
			fmt.Println("Want something else?")
		case 2:
			beverage = beveragedecorator.NewMocha(beverage)
			fmt.Println("Your order:", beverage.GetDescription(), beverage.Cost())
			fmt.Println("Want something else?")
		case 3:
			return beverage
		}
	}
	return beverage
}
