package main

import "fmt"

type ISportFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func getSportFactory(brand string) (ISportFactory, error) {
	if brand == "nike" {
		return &Nike{}, nil
	}
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	return nil, fmt.Errorf("brand brand not found")
}
