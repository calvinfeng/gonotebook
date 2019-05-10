package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/add", CreateOperationHandler(Addition))
	http.HandleFunc("/api/sub", CreateOperationHandler(Subtraction))
	http.HandleFunc("/api/mul", CreateOperationHandler(Multiplication))
	http.HandleFunc("/api/div", CreateOperationHandler(Division))
	fmt.Println("starting calculator server on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("failed to serve", err)
	}
}
