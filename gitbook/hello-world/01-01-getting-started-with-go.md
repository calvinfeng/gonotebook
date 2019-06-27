# Getting Started With Go

## Installation

Go to [Golang.org download page](https://golang.org/dl/) and download `go<version>.darwin-amd64.pkg` for your Mac OS. It should include an installer and provide instruction on how to install Golang step by step.

When you are done, make sure you run `go` in your terminal. If Go has been successfully installed on your machine, you should expect to see the following text.

```text
MacBook: ~Calvin$ go
Go is a tool for managing Go source coode.

Usage:

    go command [arguments]
```

## Project Structure

### Workspace

Go has this concept of a workspace. It's basically a folder where you'd put all your source code for all your Go programs. I usually put my workspace in my `~` directory. I called my workspace `Gopher` or `Workspace`, but feel free to name it whatever you like.

```text
Calvin/
    Applications/
    Desktop/
    Documents/
    etc...
    Gopher/
        bin/
        pkg/
        src/
```

### Go Path

In order for Go to recognize your workspace directory, you must go to your home directory and define your environmental variables in your `.bash_profile`.

For example:

```bash
cd ~
code .bash_profile
```

And then insert the following into your bash profile:

```bash
# Go paths
export GO=/user/local/go
export GOPATH=/Users/Calvin/Gopher
export PATH=$PATH:$GO/bin:$GOPATH/bin
```

Now, whenever you start a new project, put it into `$GOPATH/src` folder. For example: Gopher/ bin/ pkg/ src/ go-academy/ first\_program/ hello\_world.go main.go bye\_world.go

### Go Module

{% hint style="info" %}
Skip this section for now and re-visit when you start reading websocket server section. 
{% endhint %}

If you are using Go 1.11+, you have the option of using go module which allows you to put your source code anywhere on your computer. Although the next 3 sections are written with `go dep` as the dependency manager, you are free to use go module.

To activate go module, insert the following environmental variable in your `.bash_profile`.

```bash
export GO111MODULE=on
```

Navigate to your Desktop and create a new folder `helloworld`.

```bash
cd ~/Desktop
mkdir helloworld
cd helloworld/
```

Initialize a go module

```bash
go mod init github.com/calvinfeng/helloworld
```

It is indicating to Go that my project will be publicly available on `github.com/calvinfeng`, but if you don't intend to share your code, the path is optional. You can initialize it with whatever name you like. Suppose another person who would like to use my packages from my `helloworld` repository, he/she simply needs to perform a import in his/her source code.

```go
package main

import "github.com/calvinfeng/helloworld/dog

func main() {
    dog.SaysHello()
}
```

