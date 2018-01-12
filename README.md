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
4. HTTP Server in Go
    * Project Requirement
    * Handlers
    * Video 03 - Calculator Server
5. WebSocket Server in Go
    * Project Requirement
    * Dependency Management
    * Frontend
    * Node Modules
    * Video 04a - Chat server
    * Video 05b Concurrency
6. User Authentication
    * Project Requirement
    * Video 05 - User Authentication, coming soon
7. Neural Network (Maybe)
    * Project Requirement
    * Video 06 - Neural Network, (still deciding?)

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
Yes you can write your own package. However, you cannot run `go install` on a non-main package. Strictly speaking you can
still run the `go install` command in your terminal but it has no effect. Go will not compile your package into executable.
So how do you use your own package? You must import it into your `main` program.

For example, create a directory in `go-academy` and call it `foo` and inside the folder make a file called `foo.go`:
```go
// => foo.go
package foo

func IsFoo() bool {
	return true
}
```

That's basically it. You now have a `foo` package. Now let's import it into your `hello_world` program.

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

### Video 02: Tic Tac Toe in Go
Let's create a project that requires us to write multiple files in the `main` package!

[Tic Tac Toe in Go Part 1](https://youtu.be/644HhokVkbI)

[Tic Tac Toe in Go Part 2](https://youtu.be/eL6ruTgOQG0)

[Tic Tac Toe in Go Part 3](https://youtu.be/rdSgqye50Qw)

### Bonus
Create a computer player that implements the `Player` interface and make it undefeatable with *Minimax* algorithm.

Here are some other recommendations for projects you can work on

* Connect 4
* Minesweeper
* Sudoku

## HTTP Server in Go
### Project Requirement
We are going to use Go's built-in package `net/http` to create a calculator server. The server should have 4 API endpoints,
each endpoint serves a different mathematical operation. For example,

* `api/add` does addition
* `api/sub` does subtraction
* `api/mul` does multiplication
* `api/div` does division

Each endpoint takes in 2 query parameters, left operand and right operand, denoted by `lop` and `rop`.

### Video 03: Calculator Server
[Calculator Server in Go Part 1](https://youtu.be/QWQjqcDYALU)

[Calculator Server in Go Part 2](https://youtu.be/8S6YPgo1Tns)

### Bonus
Add more endpoints for practice!

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
or in this case, your own http handler.

For example
```go
type HandlerString string

// ServeHTTP returns the string itself as a response
func (str HandlerString) ServeHTTP(w ResponseWriter, r *Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(str))
}
```

## WebSocket Server in Go
### Prerequisite
Now it's time for you to finish Part 3 of [Go Tour](https://tour.golang.org/concurrency/1). The last four sections are
not required for the websocket project. You only need to finish sections from *Goroutines* to *Default Selection.*
Later we will re-visit `sync.mutex` when we dive deeper into concurrency.

### Video 04a: Concurrency

[Concurrency in Go](https://youtu.be/uq9EocsraUQ)

### Project Requirement
We are going to learn about how to perform dependency management in Go and how to use Gorilla library to implement a
websocket connection in Go. We are also going to learn about how to integrate React with Go.

### Dependency Management - `dep`
Dep is an awesome dependency management tool for Go, it's equivalent to `npm` for JavaScript. You can learn more about
Dep on https://github.com/golang/dep.

#### Local Installation
If you wish to install it locally to your project,

```
go get -u github.com/golang/dep/cmd/dep
```

However, I recommend that you do it globally like how you can run `npm` anywhere on your computer.

#### Global Installation
If you are a Mac user, the installation step is very easy for you. First of all you should have Homebrew on your Mac, then
perform the following commands in your terminal:
```
brew install dep
brew upgrade dep
```

And that's it.

### Frontend
You can copy and paste the front end code in my `user_auth/frontend` folder. However, please feel free to write your own
frontend implementation.

We are going to use JavaScript's native `WebSocket` class.
```javascript
this.ws = new WebSocket("ws://localhost:8000/ws")
this.websocket.onopen = this.handleSocketOpen;
this.websocket.onmessage = this.handleSocketMessage;
this.websocket.onerror = this.handleSocketError;
this.websocket.onclose = this.handleSocketClose;
```

Client-side socket connection typically accepts 4 callbacks:

    1. Callback is invoked when socket is opened.
    2. Callback is invoked when socket receives a message.
    3. Callback is invoked when socket encounters an error.
    4. Callback is invoked when socket is closed.

### Node Modules
I am using babel and webpack for compiling the latest ES6/ES7 syntax into browser compatible version. I am also using
`node-sass` for compiling `.scss` into `.css`. I make promise-based requests to server using `axios` instead of jQuery.
I am not a big fan of jQuery anymore.

For the complete list of dependency, please look at the `package.json`.


### Video 04b: WebSocket Server
Please excuse me that I repetitively said *so* in my speech; it's a result of having to think about what to type and what
to say simultaneously. I will make sure that in my next video I will suppress my impulse to say *so*.

[WebSocket Server in Go Part 1](https://youtu.be/X-FeLB35jXs)

[WebSocket Server in Go Part 2](https://youtu.be/fn8nBckuUE8)


## User Authentication
### Project Requirements
TBA

### Postgres in Go
I am going to use PostgreSQL for this project, so let's create one. The superuser on my computer is `cfeng` so I will use
that to create a database named `go_user_auth`

If you don't have a role or wish to create a separate role for this project, then just do the following
```
$ psql postgres
postgres=# create role <name> superuser login;
```

Create a database named `go_user_auth` with owner pointing to whichever role you like. I am using cfeng on my computer.
```
$ psql postgres
postgres=# create database go_user_auth with owner=cfeng;
```

Actually just in case you don't remember the password to your `ROLE`, do the following
```
postgres=# alter user <your_username> with password <whatever you like>
```

I did mine with
```
postgres=# alter user cfeng with password "cfeng";
```

## Additional Resource
If you want to learn more about session storage, security, encryption, and many other topics relating to web applications,
take a look at this eBook: https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/.

## Neural Network
Still under development
