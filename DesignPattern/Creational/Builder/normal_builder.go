package main

type NormalBuilder struct {
	windowType  string
	doorType    string
	numberFloor int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (n *NormalBuilder) setWindowType() {
	n.windowType = "Wooden Window"
}

func (n *NormalBuilder) setDoorType() {
	n.doorType = "Wooden Door"
}

func (n *NormalBuilder) setNumberFloor() {
	n.numberFloor = 2
}

func (n *NormalBuilder) getHouse() House {
	return House{
		windowType:  n.windowType,
		doorType:    n.doorType,
		numberFloor: n.numberFloor,
	}
}
