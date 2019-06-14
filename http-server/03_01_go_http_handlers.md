# HTTP Handler

Inside the HTTP package, there are two handler types, `http.Handler` and `http.HandlerFunc`. It is natural to ask why do we have two types of handler.

* `http.Handler` is an interface.
* `http.HandlerFunc` is a custom function type which takes two arguments, `ResponseWriter` and `*Request`.

Any data type that implements the method `ServeHTTP` will qualify as a HTTP handler. Thus, if we attach a `ServeHTTP` method to a data type, then that data type is automatically a `http.Handler`.

In Go, you can attach methods to any data type, even a string or an integer.

```go
type HandlerString string

// ServeHTTP returns the string itself as a response
func (str HandlerString) ServeHTTP(w ResponseWriter, r *Request) {
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(str))
}
```

Now it is much easier to explain what is a `HandlerFunc`. Essentially it is a function that implements `ServeHTTP`.

```go
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
  f(w, r)
}
```

If this is still confusing, don't worry. The concept will become clear once we work on calculator server.

