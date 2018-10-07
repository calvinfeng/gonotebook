package main

import (
	"fmt"
	"go-academy/helloworld/foo"
)

func SayHello() {
	fmt.Println("Hello")
}

func SayBye() {
	fmt.Println("Bye")
}

func main() {
	fmt.Println("Hello World, this is Go!")
	SayHello()
	SayBye()
	foo.SayHello()
}
