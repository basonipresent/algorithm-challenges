package main

import "fmt"

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumberFloor()
	getHouse() House
}

func getBuilder(builderType string) (IBuilder, error) {
	if builderType == "normal" {
		return newNormalBuilder(), nil
	} else if builderType == "igloo" {
		return newIglooBuilder(), nil
	}
	return nil, fmt.Errorf("invalid builder type")
}
