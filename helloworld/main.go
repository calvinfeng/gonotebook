package main

import (
	"fmt"
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
}
