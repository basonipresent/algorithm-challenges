package main

type IglooBuilder struct {
	windowType  string
	doorType    string
	numberFloor int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (i *IglooBuilder) setWindowType() {
	i.windowType = "Snow Window"
}

func (i *IglooBuilder) setDoorType() {
	i.doorType = "Snow Door"
}

func (i *IglooBuilder) setNumberFloor() {
	i.numberFloor = 1
}

func (i *IglooBuilder) getHouse() House {
	return House{
		windowType:  i.windowType,
		doorType:    i.doorType,
		numberFloor: i.numberFloor,
	}
}
