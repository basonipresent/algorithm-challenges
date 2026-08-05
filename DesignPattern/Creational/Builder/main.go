package main

import "fmt"

func main() {
	normalBuilder, _ := getBuilder("normal")
	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()
	fmt.Printf("WindowType: %s, DoorType: %s, NumberFloor: %d \n", normalHouse.windowType, normalHouse.doorType, normalHouse.numberFloor)

	iglooBuilder, _ := getBuilder("igloo")
	director.setBuider(iglooBuilder)
	iglooHouse := director.buildHouse()
	fmt.Printf("WindowType: %s, DoorType: %s, NumberFloor: %d \n", iglooHouse.windowType, iglooHouse.doorType, iglooHouse.numberFloor)

	errBuilder, err := getBuilder("err")
	if err != nil {
		fmt.Println(err)
		return
	}
	director.setBuider(errBuilder)
	errHouse := director.buildHouse()
	fmt.Printf("WindowType: %s, DoorType: %s, NumberFloor: %d \n", errHouse.windowType, errHouse.doorType, errHouse.numberFloor)
}
