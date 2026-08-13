package main

func main() {
	hp := &Hp{}
	epson := &Epson{}

	mac := &Mac{}
	mac.SetPrinter(hp)
	mac.Print()
	mac.SetPrinter(epson)
	mac.Print()

	windows := &Windows{}
	windows.SetPrinter(hp)
	windows.Print()
	windows.SetPrinter(epson)
	windows.Print()
}
