package main

import "fmt"

type WindowsAdapter struct {
	wMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("WindowsAdapter: Adapter converts Lightning signal to USB")
	w.wMachine.insertIntoUSBPort()
}
