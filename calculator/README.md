# Calculator Server

## Setup

Same as before, create a folder named `calculator` in your `go-academy` directory. We don't need to
create additional packages for this project because it is much more straightforward logically speaking.

## HTTP Server

We are going to use Go's built-in package `net/http` to create a calculator server. The server should
have 4 API GET endpoints, each endpoint serves a different arithmetic operation.

For example,

* `api/add` does addition
* `api/sub` does subtraction
* `api/mul` does multiplication
* `api/div` does division

Each endpoint takes in 2 query parameters, left operand and right operand, denoted by `lop` and `rop`.

## Project Calculator

* [Lesson 3 Calculator Server](https://youtu.be/_baFDzyZxPg)

## Homework

Add more endpoints for practice!

## Source

[GitHub](https://github.com/calvinfeng/go-academy/tree/master/calculator)