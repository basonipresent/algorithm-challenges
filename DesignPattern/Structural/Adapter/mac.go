package main

import "fmt"

type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Mac: Lightning connector is plugged into mac machine.")
}
