# Packages in Go
## What is `main`?
In my previous article and video I did not explain `package main` and `func main()` thoroughly. So 
what is main? First of all, we have a `main` package, which is the package that is responsible for 
execution of a program. On the other hand, we have a main function inside `main`. The `main` package
cannot be built without the existence of a `main func`. Unlike dynamic typed, scripting languages 
(e.g. Ruby, Python, JavaScript and etc...), Go does not allow you to execute any function outside 
of `main`.

For example:
```go
// => main.go
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
Yes you can write your own package. However, you cannot run `go install` on a non-main package. 
Technically speaking you can still run the `go install` command in your terminal but it has no 
effect. Go will not compile your package into executable. So how do you use your own package? You 
must import it into your `main` program.

For example, create a directory in `go-academy` and call it `foo` and inside the folder make a file 
called `foo.go`:
```go
// => foo.go
package foo

func IsFoo() bool {
	return true
}
```

That's basically it. You now have a `foo` package. Now let's import it into your `hello_world` 
program.
```go
// => hello_world/main.go
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
```