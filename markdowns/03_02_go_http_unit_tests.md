# HTTP Unit Tests

Now you have written a simple HTTP server with multiple endpoints. It is time to write unit tests
for them. You may be wondering, isn't it troublesome to setup a server everytime to test a
particular handler for a given endpoint? Well turns out, you don't actually need a server to test
handler logic.

## `httptest`

Golang has a built-in test package for `http` package related items. Recall that HTTP handler has
the following delcaration.

```go
package http

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

When we write a function like the following

```go
func FooHandler(w http.ResponseWriter, r *http.Request) {
    // Logic...
}
```

This is known as a `Handlerfunc` which is essentially an adapter to allow the use of ordinary
function as HTTP handlers. It does so by attaching a `ServeHTTP` method to `FooHandler`.

```go
http.HandlerFunc(FooHandler) // Casting

// Gain the following method automatically.
func (f FooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    f(w, r)
}
```

In order to test any handler, we only need to provide two inputs, i.e. a response writer and a
request object. If we can somehow mock these two items, we can test any handler function. This is
exactly what `httptest` package provides you.

## Mocks

Let's first mock a response writer using `*httptest.ResponseRecorder`. The recorder records the
response coming from a handler. Remember that handler uses `w.Write` to write bytes into HTTP
responses. We want to record those bytes and see if the handler is behaving correctly. The recorder
gives us a HTTP response which has the recorded bytes via `recorder.Result()`.

```go
import "net/http/httptest"

recorder := httptest.NewRecorder()
resp := recorder.Result()
```

The next item to mock is request. This one is also very straightforward, you simply ask the package
to give you a request object. You can add headers, body, and query parameters to it, just like a
regular request.

```go
req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
```

We are now ready to put everything together.

```go
func TestFooHandler(t *testing.T) {
    rec := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)

    FooHandler(rec, req)

    resp := rec.Result()
    body, _ := ioutil.ReadAll(resp.Body)

    // Perform test logic
    assertEqual(t, http.StatusOK, resp.StatusCode)
    assserEqual(t, "foo", string(body))
}
```