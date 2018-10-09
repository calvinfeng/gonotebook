# Packages in Go
Code is organized by packages in Go. They are similar to Python's modules and JavaScript's npm 
modules.

## What is `main`?
The most common package in Go is `main` because main package is served as an entry point to any Go 
program.

For example:
```go
package main

import "fmt"

func HelloWorld() {
	fmt.Println("Hello World!")
}

HelloWorld()
```

This will not compile and throw an error. You must call your functions inside the `main` function.
```go
package main

import "fmt"

func HelloWorld() {
	fmt.Println("Hello World!")
}

func main() {
	HelloWorld()
}
```

### What about non-main packages?
We can write our own packages. However, we cannot run a non-main package directly. We must import it 
into the `main` package. For example, create a new folder in `helloworld` and call it `foo`. Inside 
the folder, make a file called `foo.go`.
```go
package foo

func SayHello() {
	fmt.Println("Foo says hello")
}
```

Now we can import the package into the main program and run `foo.SayHello`. 
```go
// => hello_world/main.go
package main

import (
	"fmt"
	"go-academy/helloworld/foo"
)

func main() {
	fmt.Println("Hello World, this is Go!")
	SayHello()
	SayBye()
	foo.SayHello()
}
```