package main

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windows := &Windows{}
	windowsAdapter := &WindowsAdapter{wMachine: windows}
	client.InsertLightningConnectorIntoComputer(windowsAdapter)
}
