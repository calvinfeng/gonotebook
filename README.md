# Go Academy
Welcome to Go Academy, the official rip-off of App Academy, created by your beloved a/A alum.

## Table of Contents
1. Prerequisites
2. Go Project Setup
  - Video 01 - Hello World in Go
3. Tic Tac Toe in Go
4. TBA
5. TBA

## Prerequisite(s)
Before you start watching any of the videos listed below. It's important to get yourself familiar with Go's syntax first
with Go tour.

Go to https://tour.golang.org/ and complete the **Basics** and **Methods and Interfaces** sections of the tutorial.

## Getting Started with Go
### Installation
Go to https://golang.org/dl/ and download `go1.9.2.darwin-amd64.pkg` for your Mac OS X. It should include an installer and
provides instruction on how to install Go step by step. When you are done, make sure you run `go` in your terminal. If
Go has been successfully installed on your machine, you should expect to see the following being printed to your screen.

```
MacBook: ~Calvin$ go
Go is a tool for managing Go source coode.

Usage:

        go command [arguments]
```

### Project Structure
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

### Hello World in Go
By now you should have your Go installed, your workspace created, and your `$GOPATH` is pointing to the workspace. Now let's create our first program in Go!

[Hello World in Go](https://youtu.be/5-FFapKA9sM)

In case you don't really know what a compiled language is or how language is being compiled from source code to machine code. Here is a fun introduction to the concept:

[How Do Computers Read Code ](https://www.youtube.com/watch?v=QXjU9qTsYCc)
