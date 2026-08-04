package main

import "fmt"

func main() {
	factory, _ := getSportFactory("nike")
	shoe := factory.makeShoe()
	shirt := factory.makeShirt()
	printShoe(shoe)
	printShirt(shirt)

	factory, _ = getSportFactory("adidas")
	shoe = factory.makeShoe()
	shirt = factory.makeShirt()
	printShoe(shoe)
	printShirt(shirt)
}

func printShoe(s IShoe) {
	fmt.Println(s.getLogo(), s.getSize())
}

func printShirt(s IShirt) {
	fmt.Println(s.getLogo(), s.getSize())
}
