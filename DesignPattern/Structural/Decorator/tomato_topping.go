package main

type TomatoTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) GetPrice() int {
	pizzaPrice := t.pizza.GetPrice()
	return pizzaPrice + 7
}
