# Build a Calculator

## Setup

Create a folder named `calculator` in your `go-academy` directory. We don't need to create additional packages for this project. We can keep everything in the package `main`.

## HTTP Server

We are going to use Go's built-in package `net/http` to create a calculator server. The server should have 4 API GET endpoints, each endpoint serves a different arithmetic operation.

For example,

* `api/add` does addition
* `api/sub` does subtraction
* `api/mul` does multiplication
* `api/div` does division

Each endpoint takes in 2 query parameters, left operand and right operand, denoted by `lop` and `rop`.

## Project Calculator

{% embed url="https://www.youtube.com/watch?v=\_baFDzyZxPg&feature=youtu.be" %}

## Bonus

Add more mathematical operation endpoints for practice!

## Source

[GitHub](https://github.com/calvinfeng/go-academy/tree/master/calculator)

