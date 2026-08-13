package main

import "fmt"

func main() {
	pizza := &VeggieMania{}

	pizzaWithCheese := &CheeseTopping{pizza: pizza}

	pizzaWithCheeseAndTomato := &TomatoTopping{pizza: pizzaWithCheese}

	fmt.Printf("Price of pizza with cheese and tomato: %d\n", pizzaWithCheeseAndTomato.GetPrice())
}
