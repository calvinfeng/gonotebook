# [Introduction](https://calvinfeng.gitbook.io/go-academy/)
## Go Tour
Head to [Go Tour](https://tour.golang.org/) and complete the **Basics** and **Methods and Interfaces** 
sections of the tutorial. I don't expect you to have memorized all the syntax right away. As you start 
building projects with Go, you will become more comfortable with the syntax.

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

[WebSocket Server in Go Part 1](https://youtu.be/KtdZinZIe3A)

[WebSocket Server in Go Part 2](https://youtu.be/Ue7z2BEavBU)


## User Authentication
### Project Requirements
We need to set up our PostgreSQL first, please refer to the section below. I am going to introduce couple new open source
libraries to you for this project:

* `sirupsen/logrus`
* `gorilla/mux`
* `jinzhu/gorm` -  Object Relational Mapping for Go

You should take a look at their Github page and see what they are for before you start working on this project.

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

### Video 05: User Authentication in Go

[User Authentication in Go Introduction](https://youtu.be/t7UaOV0THIQ)


## Additional Resource
If you want to learn more about session storage, security, encryption, and many other topics relating to web applications,
take a look at this eBook: https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/.

## Neural Network
### Project Requirements
If you are familiar with vectorized implementation of neural network, feel free to jump ahead and start watching videos
on Golang part, otherwise we need to go over some the basics and mathematics of neural nets.

#### Jupyter Notebook
I will use Python to teach math because I can write LaTex in Jupyter notebooks. Also, `numpy` is very convenient for matrix
operations. We will see that we have a `numpy` equivalent in Golang called `gonum` (I wonder why not `numgo`?)

So let's get started by installing `pip`. I think `easy_install` is provided by Mac OS X, so we don't need to use Homebrew.
```
sudo easy_install pip
```

Once you have `pip`, now use it to install `virtualenv` (Python virtual environment)
```
pip install virtualenv
```

If permission denied, use `sudo`. This is equivalent to install `npm` globally. Remember that `pip` is a package manager
for Python, just like npm for Node.
```
sudo pip install virtualenv
```

Now go to `neural_net` directory and create a virtual environment
```
cd $GOPATH/src/go-academy/neural_net/
virtualenv environment
```

Activate your environment
```
source environment/bin/activate
```

Install all the required dependencies
```
pip install numpy
pip install matplotlib
pip install jupyter
```

Now you are good to go, let's run Jupyter!!! Make a directory called `notebooks` in `neural_net`
```
mkdir notebooks
cd notebooks
jupyter notebook
```

#### Neural Network Videos
Need more time to work on the videos for this project...
