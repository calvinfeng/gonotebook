# Go Module

We will start using Go module from this point on. One of the immediate visible benefits to you is that you no longer need to put your source code in $GOPATH. 

Let's create a new folder on your Desktop and this is going to be our go module.

```bash
cd ~/Desktop
mkdir go-academy
```

Make sure you set environmental variable `GO111MODULE` to be on, either through command line interface or `bash_profile.`

```bash
export GO111MODULE=on
```

Now initialize your first go module.

```bash
go mod init github.com/calvinfeng/go-academy
```

Create a new project within the module and add a `foo` package to it.

```bash
mkdir sillyserver
cd sillyserver
mkdir handler
cd handler
touch silly.go
```

Let's create a simple handler within this package.

```go
package handler

import (
	"net/http"

	"github.com/Pallinder/go-randomdata"
)

func Silly(w http.ResponseWriter, r *http.Request) {
	paragraph := randomdata.Paragraph()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(paragraph))
}
```

Navigate back to the your `sillyserver`

```bash
cd ..
touch main.go
```

Create a server

```go
package main

import (
	"net/http"

	"github.com/calvinfeng/go-academy/sillyserver/handler"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler.Silly),
	}

	srv.ListenAndServe()
}
```

Build and run the server,and witness the magic of go module.

```bash
go install && sillyserver
```

