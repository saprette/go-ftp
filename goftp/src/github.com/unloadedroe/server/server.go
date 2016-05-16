package main

import (
	"fmt"
	"github.com/unloadedroe/filesystem"
)

func main() {
	path := filesystem.NewPath("hello.txt")
	fmt.Println(string(filesystem.Read(path).Content))
}
