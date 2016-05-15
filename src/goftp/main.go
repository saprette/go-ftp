package main

import (
	"fmt"
	"./filesystem"

	"path"
)

func main() {
	filesystem := new(filesystem.FileSystem)
	path := path.NewPath("/Users/eleon/Github/hello.txt")
	fmt.Println(filesystem.Read(path))
}
