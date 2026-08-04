package main

import (
	"fmt"
)

func main() {
	ak47, err := getGun("ak47")
	if err != nil {
		fmt.Println(err)
		return
	}
	printGunInfo(ak47)
	musket, err := getGun("musket")
	if err != nil {
		fmt.Println(err)
		return
	}
	printGunInfo(musket)
}

func printGunInfo(gun IGun) {
	fmt.Printf("Gun name: %s, Power: %d\n", gun.getName(), gun.getPower())
}
