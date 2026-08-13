package main

import "fmt"

type Hp struct{}

func (h *Hp) PrintFile() {
	fmt.Println("Print file from HP")
}
