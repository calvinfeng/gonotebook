package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/add", HandleAdd)
	http.HandleFunc("/api/subtract", HandleSubtract)
	http.HandleFunc("/api/multiply", OperationHandlerCreator(MUL))
	http.HandleFunc("/api/divide", OperationHandlerCreator(DIV))

	fmt.Println("Starting calculator server on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Failed to start server")
	}
}
