package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("hello/hello.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("lee", "roy", "akbar"))
}
