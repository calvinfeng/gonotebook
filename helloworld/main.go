package main

import (
	"go-academy/helloworld/greeting"
)

func main() {
	greeting.SetGreeter("Calvin")
	greeting.HelloWorld()
	greeting.ByeWorld()
}
