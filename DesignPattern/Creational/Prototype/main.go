package main

import "fmt"

func main() {
	file1 := &File{
		name: "file1",
	}
	file2 := &File{
		name: "file2",
	}
	file3 := &File{
		name: "file3",
	}

	folder1 := &Folder{
		name: "folder1",
		children: []INode{
			file1,
		},
	}
	folder2 := &Folder{
		name: "folder2",
		children: []INode{
			folder1, file2, file3,
		},
	}

	fmt.Println("Print Hierarchy of Folder2")
	folder2.print(" ")

	cloneFolder := folder2.clone()
	fmt.Println("Print Hierarchy of Clone Folder")
	cloneFolder.print(" ")
}
