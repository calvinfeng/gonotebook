# Hello World in Go
By now you should have your Go installed, your workspace created, and your `$GOPATH` is pointing to 
the workspace. Let's create our first program in Go! If you prefer to have everything explained step
by step, please skip ahead and watch the video. Otherwise, text will be the quickest way to go 
through this tutorial.

## Functions
First of all, create a new folder in your Go workspace and name it whatever you like. Inside the new
folder, create a file called `main.go`. Package `fmt` is Go's builtin package for string formatting 
and string printing. Let's create two functions inside `main.go` that will use `fmt`. 
```golang
package main

import "fmt"

func SayHello() {
	fmt.Println("Hello")
}

func SayBye() {
	fmt.Println("Bye!!!")
}
```

The functions above are very self-explanatory. The next step is to invoke them, which must occur in
the main function, following C's tradition.
```golang
func main() {
    SayHello()
    SayBye()
}
```

Now run `go run main.go`. You are done! This is Hello World in Go.

## (Optional) Video 01: Hello World in Go
[Hello World in Go](https://youtu.be/5-FFapKA9sM)