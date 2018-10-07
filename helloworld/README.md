# Hello World in Go
By now you should have your Go installed, your workspace created, and your `$GOPATH` is pointing to 
the workspace. It's time for hello world.

## Functions
Create a new folder named `helloworld`. Inside this new folder, create a `main.go` file. Package 
`fmt` is Go's builtin package for common I/O usage, e.g. printing a string to screen, scanning 
terminal input. Let's create two functions inside `main.go` that will use `fmt`. 
```golang
package main

import "fmt"

func SayHello() {
	fmt.Println("Hello")
}

func SayBye() {
	fmt.Println("Bye")
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

In your terminal, run your program with
```
go run main.go
```

## (Optional) Video 01: Hello World in Go
[Hello World in Go](https://youtu.be/5-FFapKA9sM)