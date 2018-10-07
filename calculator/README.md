# Calculator Server
[Source code is here](https://github.com/calvinfeng/go-academy/tree/master/calculator)

## Project Requirement
We are going to use Go's built-in package `net/http` to create a calculator server. The server 
should have 4 API endpoints, each endpoint serves a different mathematical operation. For example,

* `api/add` does addition
* `api/sub` does subtraction
* `api/mul` does multiplication
* `api/div` does division

Each endpoint takes in 2 query parameters, left operand and right operand, denoted by `lop` and 
`rop`.

## Serving & Routing
First let's create a `main.go` file as an entry point to our server. Inside this file, add your main
function and let's begin to use Golang's built-in `http` package.
```golang
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
```

In the snippet above, we are using `http` package's built-in server. We add route to the server by
calling `HandleFunc`. The function takes a path string and a handler function as arguments. What are
the handler functions? They are functions that will handle requests.

For example, this is a handler for the `Add` endpoint.
```golang
func HandleAdd(w http.ResponseWriter, r *http.Request) {
	leftOp, leftErr := strconv.ParseFloat(r.URL.Query().Get("lop"), 64)
	rightOp, rightErr := strconv.ParseFloat(r.URL.Query().Get("rop"), 64)

	result := fmt.Sprintf("%v + %v = %v", leftOp, rightOp, leftOp+rightOp)
	if leftErr == nil && rightErr == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid query parameters"))
	}
}
```

The handler is expecting two query parameters; they represent the left operand and right operand. If
I want to add two numbers, say 3 + 4, then I will hit `http://localhost:8000/api/add/?lop=3&rop=4`.
Once you understand the example above, the rest is pretty straightforward.

## Bonus
Add more endpoints for practice!

## (Optional) Video 03: Calculator Server

* [Calculator Server in Go Part 1](https://youtu.be/QWQjqcDYALU)
* [Calculator Server in Go Part 2](https://youtu.be/8S6YPgo1Tns)



