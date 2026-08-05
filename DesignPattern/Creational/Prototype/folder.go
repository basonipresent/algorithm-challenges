package main

import "fmt"

type Folder struct {
	children []INode
	name     string
}

func (f *Folder) print(indent string) {
	fmt.Println(indent + f.name)
	for _, i := range f.children {
		i.print(indent + indent)
	}
}

func (f *Folder) clone() INode {
	cloneFolder := &Folder{
		name: f.name + "_clone",
	}
	var tmpChildren []INode
	for _, i := range f.children {
		copy := i.clone()
		tmpChildren = append(tmpChildren, copy)
	}
	cloneFolder.children = tmpChildren
	return cloneFolder
}
