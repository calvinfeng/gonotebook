# Handlers
When you google around about Go http handlers, you will notice that there is something called 
`http.Handler` and `http.HandlerFunc`. It is natural to ask why do we have two types of handler and 
they both work?!

`http.Handler` is an interface. Any data type that implements the method `ServeHTTP` will qualify 
as a HTTP handler. So, if you somehow can attach a method to a function, then that function is 
indeed an authentic HTTP handler. In Go, you can attach methods to any data type, even a string or 
an integer.
```go
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

Essentially, `HandlerFunc` is a type of `Handler`, just like Fuji apple is a type of apple. You can 
define your own apple, or in this case, your own http handler. For example, I can use a string as my 
HTTP handler! This is weird but it works.
```go
type HandlerString string

// ServeHTTP returns the string itself as a response
func (str HandlerString) ServeHTTP(w ResponseWriter, r *Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(str))
}
```