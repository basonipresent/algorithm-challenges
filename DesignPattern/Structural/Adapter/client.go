package main

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client: Inserting lightning connector into computer")
	com.InsertIntoLightningPort()
}
