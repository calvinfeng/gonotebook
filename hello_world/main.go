package main

import (
	"fmt"
	"go-academy/foo"
)

func main() {
	fmt.Println("Hello World, this is Go!")

	SayHello()
	SayBye()

	fmt.Printf("Is Foo a Foo? %v \n", foo.IsFoo())
}