package greeting

import "fmt"

var name string

// SetGreeter attaches an author name to all greeting messages.
func SetGreeter(n string) {
	name = n
}

// HelloWorld prints hello world to the screen.
func HelloWorld() {
	fmt.Printf("%s says hello world!\n", name)
}

// ByeWorld prints bye world to the screen.
func ByeWorld() {
	fmt.Printf("%s says bye world!\n", name)
}
