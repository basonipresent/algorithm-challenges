package main

import "fmt"

type Windows struct {
}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("Windows: USB connector is plugged into windows machine.")
}
