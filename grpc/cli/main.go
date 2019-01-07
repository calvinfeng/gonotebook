package main

import "fmt"

func main() {
	t := NewTodo(1)
	if err := t.Load(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t)
	}
}
