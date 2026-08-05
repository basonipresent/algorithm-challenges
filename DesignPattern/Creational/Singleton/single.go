package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var singleInstance *single

type single struct {
}

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Create single instance")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already exists")
		}
	} else {
		fmt.Println("Single instance already exists")
	}
	return singleInstance
}
