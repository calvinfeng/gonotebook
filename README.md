# Go Academy
Welcome to Go Academy, the official rip-off of App Academy, created by your beloved a/A alum.

## Table of Contents
1. Prerequisites
2. Getting Started with Go
    * Installation
    * Project Structure
    * Video 01: Hello World in Go
3. Packages in Go
    * What is `main`?
    * Video 02: Tic Tac Toe in Go
    * Bonus - Minimax algorithm
4. HTTP server in Go
    * Project Requirement
    * Handlers
    * Video 03 - Calculator Server
5. WebSocket server in Go
    * Dependency Management
    * Video 04 - Chat server
6. Concurrency
    * Video 05 - MapReduce 
7. Postgres in Go
    * Video 06 - 99Cats in Go

## Prerequisite(s)
Before you start watching any of the videos listed below, it's important to get yourself familiar with Go's syntax by 
going through Go tour.

Go to https://tour.golang.org/ and complete the **Basics** and **Methods and Interfaces** sections of the tutorial. I 
don't expect you to have memorized all the syntax right away. As you start building projects with Go, you will become 
more comfortable with the syntax.

## Getting Started with Go
### Installation
Go to https://golang.org/dl/ and download `go1.9.2.darwin-amd64.pkg` for your Mac OS X. It should include an installer and
provide instruction on how to install Go step by step. When you are done, make sure you run `go` in your terminal. If
Go has been successfully installed on your machine, you should expect to see the following being printed to your screen.

```
MacBook: ~Calvin$ go
Go is a tool for managing Go source coode.

Usage:

        go command [arguments]
```

### Project Structure
#### Workspace
Go has this concept of a workspace. It's basically a folder where you'd put all your source code for all your Go programs.
I usually put my workspace in my home directory. I called my workspace `Gopher` but feel free to name it whatever you like.
```
Calvin
        - Applications
        - Desktop
        - Documents
        - etc...
        - Gopher
                - bin
                - pkg
                - src
```

#### Go Path
In order for Go to recognize your workspace directory, you must go to your home directory and define your environmental
variables in your `.bash_profile`.

For example:
```
cd ~
atom .bash_profile
```

And then insert the following into your bash profile:
```
# Go paths
export GO=/user/local/go
export GOPATH=/Users/Calvin/Gopher
export PATH=$PATH:$GO/bin:$GOPATH/bin
```

Now, whenever you start a new project, put it into `$GOPATH/src` folder.

For example:
```
- Gopher
        - bin
        - pkg
        - src
                - go-academy
                        - first_program
                                - hello_world.go
                                - main.go
                                - bye_world.go
```

### Video 01: Hello World in Go
By now you should have your Go installed, your workspace created, and your `$GOPATH` is pointing to the workspace. 
Let's create our first program in Go!

[Hello World in Go](https://youtu.be/5-FFapKA9sM)

### Compiled Language
In case you don't really know what a compiled language is or how language is being compiled from source code to machine 
code. Here is a fun introduction to the concept:

[How Do Computers Read Code ](https://www.youtube.com/watch?v=QXjU9qTsYCc)

## Packages in Go
### What is `main`?
In my previous video I did not explain `package main` and `func main()` thoroughly. So what is main? First of all, we have
a `main` package, which is the package that is responsible for execution of a program. On the other hand, we have a main
function inside `main`. The `main` package cannot be built without the existence of a `main func`. Unlike dynamic typed,
scripting languages (e.g. Ruby, Python, JavaScript and etc...), Go does not allow you to execute any function outside of
`main`. 

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
Yes you can write your own package. However, you cannot run go install on a non-main package. I mean you can run the command,
but it has no effect. Go will not compile your package into executable. So how do you use your own package? You must import
it into your `main` program. 

For example:
```go
// => foo.go
package foo

func IsFoo() bool {
	return true
}
```

```go
// => main.go
package main

import (
	"fmt"
	"go-academy/foo"
) 

func main() {
	fmt.Printf("Is Foo a Foo? %v", foo.IsFoo())
}
```

### Video 02: Tic Tac Toe in Go
Let's create a project that requires us to write multiple files in the `main` package!

[Tic Tac Toe in Go](https://youtu.be/5-FFapKA9sM)

### Bonus
Create a computer player that implements the `Player` interface and make it undefeatable with *Minimax* algorithm.

## HTTP Server
### Project Requirement
We are going to use Go's built-in package `net/http` to create a calculator server. The server should have 4 API endpoints,
each endpoint serves a different mathematical operation. For example, 

* `api/add` does addition
* `api/sub` does subtraction
* `api/mul` does multiplication
* `api/div` does division

Each endpoint takes in 2 query parameters, left operand and right operand, denoted by `lop` and `rop`.

### Handlers
When you google around about Go http handlers, you will notice that there is something called `http.Handler` and `http.HandlerFunc`. 
It is natural to ask why do we have two types of handler and they both work?! 

`http.Handler` is an interface. Any data type that implements the method `ServeHTTP` will qualify as a HTTP handler. So, 
if you somehow can attach a method to a function, then that function is indeed an authentic HTTP handler. In Go, you can
attach methods to any data type, even a string or an integer. 

```go
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

Essentially, `HandlerFunc` is a type of `Handler`, just like Fuji apple is a type of apple. You can define your own apple,
I mean your own handler type. 

For example
```go
type HandlerString string

// ServeHTTP returns the string itself as a response
func (str HandlerString) ServeHTTP(w ResponseWriter, r *Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(str))
}
```

### Video 03: Calculator Server
### Bonus
Add more endpoints for practice!